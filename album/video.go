package album

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/ivanizag/apple2-activities/operator"
)

/*
A Video is the whole of what a machine does while an activity runs, every
frame, with its sound, as an MP4 file to watch or put on YouTube. It starts
with a card with the title of the activity, and the keys pressed are shown
under the screen as they are pressed, each fading after a while; the game
port, once used, on the left of the screen, a joystick or the paddles and the
buttons, and the mouse on the right. The pictures
the activity takes are marked in a list of chapters, with their times, for the
description of the video. The last frame is held for a while at the end.

The frames are made at half the size of 1080p, the screen in its frame and the
keys under it, and ffmpeg doubles them, sharp, and encodes them; the sound is
put in at the end, after the title.
*/
type Video struct {
	op      *operator.Operator
	album   *Album
	path    string
	title   string
	sound   *Sound
	encoder *exec.Cmd
	frames  io.WriteCloser
	canvas  *image.RGBA
	frame   int
	keys    []shownKey
	chapter []chapterMark
	err     error

	// The game port and the mouse are shown once used: the paddles plugged
	// and the buttons used so far, and the last frame each button was down
	plugged    [4]bool
	buttonUsed [3]bool
	buttonDown [3]int
	mouseUsed  bool
	mouseDown  int
}

// The size of the frames made, and what ffmpeg makes them, 1080p
const (
	videoWidth  = 960
	videoHeight = 540
	videoScale  = 2
)

// The times of the video, in frames
const (
	// titleFrames is how long the title is shown, before the machine starts
	titleFrames = 4 * operator.FramesPerSecond
	// keyFrames is how long a key is shown after it is pressed
	keyFrames = 3 * operator.FramesPerSecond / 2
	// keysShown is the most keys shown at once, the last pressed
	keysShown = 16
	// endFrames is how long the last frame is held at the end, which is
	// often what the activity was for
	endFrames = 3 * operator.FramesPerSecond
)

type shownKey struct {
	name  string
	frame int
}

type chapterMark struct {
	name  string
	frame int
}

// videos is the video of each operator filmed, for the album to mark its
// pictures in
var videos = struct {
	sync.Mutex
	of map[*operator.Operator]*Video
}{of: map[*operator.Operator]*Video{}}

/*
StartVideo starts filming the machine of an operator, on a monitor, into an
MP4 file, with a title. Like Listen, it has to be called before the machine
runs. Close finishes the file.
*/
func StartVideo(op *operator.Operator, monitor Monitor, path string, title string) (*Video, error) {
	v := &Video{
		op:     op,
		album:  New("", monitor),
		path:   path,
		title:  title,
		canvas: image.NewRGBA(image.Rect(0, 0, videoWidth, videoHeight)),
	}
	v.encoder = exec.Command("ffmpeg", "-y", "-loglevel", "error",
		"-f", "rawvideo", "-pix_fmt", "rgba",
		"-s", fmt.Sprintf("%dx%d", videoWidth, videoHeight),
		"-r", fmt.Sprint(operator.FramesPerSecond), "-i", "-",
		"-vf", fmt.Sprintf("scale=%d:%d:flags=neighbor", videoWidth*videoScale, videoHeight*videoScale),
		"-c:v", "libx264", "-preset", "veryfast", "-crf", "18", "-pix_fmt", "yuv420p",
		v.silentPath())
	v.encoder.Stderr = os.Stderr
	frames, err := v.encoder.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("could not connect to ffmpeg: %w", err)
	}
	v.frames = frames
	if err := v.encoder.Start(); err != nil {
		return nil, fmt.Errorf("could not start ffmpeg: %w", err)
	}

	v.drawTitle()
	for range titleFrames {
		v.write()
	}
	v.sound = Listen(op)
	op.Watch(v)
	videos.Lock()
	videos.of[op] = v
	videos.Unlock()
	return v, nil
}

// silentPath is the file of the frames, before the sound is put in
func (v *Video) silentPath() string {
	return strings.TrimSuffix(v.path, ".mp4") + ".frames.mp4"
}

// Frame films a frame the machine has run
func (v *Video) Frame() {
	v.frame++
	draw.Draw(v.canvas, v.canvas.Bounds(), image.Black, image.Point{}, draw.Src)

	// The screen in its frame, at the top, made smaller if it does not fit
	screen := Frame(v.album.Screen(v.op))
	b := screen.Bounds()
	room := image.Rect(0, 0, videoWidth, videoHeight-keyStrip)
	scale := min(1, float64(room.Dx())/float64(b.Dx()), float64(room.Dy())/float64(b.Dy()))
	w, h := int(float64(b.Dx())*scale), int(float64(b.Dy())*scale)
	at := image.Rect((videoWidth-w)/2, (room.Dy()-h)/2, (videoWidth+w)/2, (room.Dy()-h)/2+h)
	if scale == 1 {
		draw.Draw(v.canvas, at, screen, b.Min, draw.Over)
	} else {
		xdraw.NearestNeighbor.Scale(v.canvas, at, screen, b, draw.Over, nil)
	}

	v.drawKeys()
	v.drawGamePort(at)
	v.drawMouse(at)
	v.write()
}

// Key shows a key pressed
func (v *Video) Key(name string) {
	v.keys = append(v.keys, shownKey{name, v.frame})
}

// mark marks a picture of the activity in the chapters
func (v *Video) mark(name string, frame int) {
	v.chapter = append(v.chapter, chapterMark{name, frame})
}

// write sends the canvas to ffmpeg as the next frame
func (v *Video) write() {
	if v.err != nil {
		return
	}
	if _, err := v.frames.Write(v.canvas.Pix); err != nil {
		v.err = fmt.Errorf("could not send a frame to ffmpeg: %w", err)
	}
}

// keyStrip is the height of the strip under the screen with the keys
const keyStrip = 90

// The colours of the keys, as the keycaps of the Apple II, and of the title
var (
	keyCap    = color.RGBA{0x40, 0x3a, 0x34, 0xff}
	keyEdge   = color.RGBA{0x90, 0x88, 0x7c, 0xff}
	keyLetter = color.RGBA{0xf0, 0xe8, 0xd8, 0xff}
	titleGrey = color.RGBA{0xa0, 0xa0, 0xa0, 0xff}
)

// drawKeys draws the last keys pressed, the newest on the right, the older
// fading
func (v *Video) drawKeys() {
	var shown []shownKey
	for _, k := range v.keys {
		if v.frame-k.frame < keyFrames {
			shown = append(shown, k)
		}
	}
	shown = shown[max(0, len(shown)-keysShown):]
	v.keys = shown

	const textScale, pad, gap = 2, 6, 6
	width := 0
	for _, k := range shown {
		width += textWidth(k.name)*textScale + 2*pad + gap
	}
	x := (videoWidth - width + gap) / 2
	y := videoHeight - keyStrip + 22
	height := 13*textScale + 2*pad
	for _, k := range shown {
		fade := 1 - float64(v.frame-k.frame)/keyFrames
		w := textWidth(k.name)*textScale + 2*pad
		box := image.Rect(x, y, x+w, y+height)
		fill(v.canvas, box, keyEdge, fade)
		fill(v.canvas, box.Inset(2), keyCap, fade)
		drawText(v.canvas, k.name, x+pad, y+pad, textScale, keyLetter, fade)
		x += w + gap
	}
}

// drawTitle draws the card of the title, as large as it fits
func (v *Video) drawTitle() {
	draw.Draw(v.canvas, v.canvas.Bounds(), image.Black, image.Point{}, draw.Src)
	lines := wrap(v.title, 30)
	longest := 0
	for _, line := range lines {
		longest = max(longest, textWidth(line))
	}
	scale := max(2, min(6, (videoWidth-80)/max(longest, 1)))
	lineHeight := 16 * scale
	y := (videoHeight-lineHeight*len(lines))/2 - 20
	for _, line := range lines {
		drawText(v.canvas, line, (videoWidth-textWidth(line)*scale)/2, y, scale, keyLetter, 1)
		y += lineHeight
	}
	under := "on an Apple II, emulated by izapple2"
	drawText(v.canvas, under, (videoWidth-textWidth(under)*2)/2, y+30, 2, titleGrey, 1)
}

// wrap splits a text in lines of at most so many characters, at its spaces
func wrap(text string, width int) []string {
	var lines []string
	line := ""
	for _, word := range strings.Fields(text) {
		if line != "" && len(line)+1+len(word) > width {
			lines = append(lines, line)
			line = ""
		}
		if line != "" {
			line += " "
		}
		line += word
	}
	return append(lines, line)
}

// textWidth is the width of a text in the font of the video, at its size
func textWidth(text string) int {
	return font.MeasureString(basicfont.Face7x13, text).Round()
}

// drawText draws a text with its top left at a point, its pixels made larger,
// faded to an opacity
func drawText(dst *image.RGBA, text string, x int, y int, scale int, c color.RGBA, opacity float64) {
	small := image.NewRGBA(image.Rect(0, 0, textWidth(text), 13))
	writer := font.Drawer{Dst: small, Src: image.NewUniform(c), Face: basicfont.Face7x13}
	writer.Dot = fixed.P(0, 10)
	writer.DrawString(text)
	big := image.NewRGBA(image.Rect(0, 0, small.Bounds().Dx()*scale, 13*scale))
	xdraw.NearestNeighbor.Scale(big, big.Bounds(), small, small.Bounds(), draw.Src, nil)
	mask := image.NewUniform(color.Alpha{A: uint8(opacity * 255)})
	draw.DrawMask(dst, big.Bounds().Add(image.Pt(x, y)), big, image.Point{}, mask, image.Point{}, draw.Over)
}

// fill paints a rectangle in a colour, faded to an opacity
func fill(dst *image.RGBA, r image.Rectangle, c color.RGBA, opacity float64) {
	mask := image.NewUniform(color.Alpha{A: uint8(opacity * 255)})
	draw.DrawMask(dst, r, image.NewUniform(c), image.Point{}, mask, image.Point{}, draw.Over)
}

// litFrames is how long a button is shown lit after it is let go, as a press
// of the operator lasts only a few frames
const litFrames = 12

// The colours of the game port and the mouse
var (
	panelLine = color.RGBA{0x60, 0x5a, 0x52, 0xff}
	panelDot  = color.RGBA{0xf0, 0x60, 0x40, 0xff}
)

// panelWidth is the width of the drawings of the game port and the mouse
const panelWidth = 110

/*
drawGamePort draws the game port on the left of the screen, once a paddle is
plugged or a button pressed: a joystick as a square with where it points, a
paddle alone as a slider, and the buttons, lit while down.
*/
func (v *Video) drawGamePort(screen image.Rectangle) {
	port := v.op.GamePort()
	for i := range port.Plugged {
		v.plugged[i] = v.plugged[i] || port.Plugged[i]
	}
	for i, down := range port.Buttons {
		if down {
			v.buttonUsed[i] = true
			v.buttonDown[i] = v.frame
		}
	}
	used := false
	for _, u := range append(v.plugged[:], v.buttonUsed[:]...) {
		used = used || u
	}
	if !used {
		return
	}

	x := max(8, (screen.Min.X-panelWidth)/2)
	y := screen.Min.Y + 40
	for pair := 0; pair < 4; pair += 2 {
		switch {
		case v.plugged[pair] && v.plugged[pair+1]:
			label := "Joystick"
			if pair > 0 {
				label = "Joystick 2"
			}
			drawText(v.canvas, label, x, y, 1, keyLetter, 1)
			y += 18
			box := image.Rect(x, y, x+panelWidth, y+panelWidth)
			fill(v.canvas, box, panelLine, 1)
			fill(v.canvas, box.Inset(2), color.RGBA{A: 0xff}, 1)
			mid := box.Min.Add(image.Pt(panelWidth/2, panelWidth/2))
			fill(v.canvas, image.Rect(box.Min.X+2, mid.Y, box.Max.X-2, mid.Y+1), panelLine, 1)
			fill(v.canvas, image.Rect(mid.X, box.Min.Y+2, mid.X+1, box.Max.Y-2), panelLine, 1)
			dot := image.Pt(box.Min.X+6+int(port.Paddles[pair])*(panelWidth-12)/255,
				box.Min.Y+6+int(port.Paddles[pair+1])*(panelWidth-12)/255)
			disc(v.canvas, dot, 6, panelDot)
			y += panelWidth + 16
		default:
			for i := pair; i < pair+2; i++ {
				if !v.plugged[i] {
					continue
				}
				drawText(v.canvas, fmt.Sprintf("Paddle %d", i), x, y, 1, keyLetter, 1)
				y += 18
				bar := image.Rect(x, y+6, x+panelWidth, y+10)
				fill(v.canvas, bar, panelLine, 1)
				knob := image.Pt(x+6+int(port.Paddles[i])*(panelWidth-12)/255, y+8)
				disc(v.canvas, knob, 7, panelDot)
				y += 30
			}
		}
	}

	for i := range port.Buttons {
		// A button is shown once used, or with its paddle: the buttons 0 and
		// 1 are those of the paddles 0 and 1, or of the joystick
		if !v.buttonUsed[i] && !(i < 2 && v.plugged[i]) {
			continue
		}
		lit := v.buttonUsed[i] && v.frame-v.buttonDown[i] < litFrames
		v.drawButton(image.Pt(x, y), fmt.Sprintf("Button %d", i), lit)
		y += 30
	}
}

// drawButton draws a round button with its name, lit or not
func (v *Video) drawButton(at image.Point, name string, lit bool) {
	c := keyCap
	if lit {
		c = panelDot
	}
	disc(v.canvas, at.Add(image.Pt(10, 10)), 10, panelLine)
	disc(v.canvas, at.Add(image.Pt(10, 10)), 8, c)
	drawText(v.canvas, name, at.X+28, at.Y+4, 1, keyLetter, 1)
}

/*
drawMouse draws the mouse on the right of the screen, once moved or pressed:
where it is, in a pad of the shape of the screen, and its button, lit while
down. The pointer itself is the program's, on the screen.
*/
func (v *Video) drawMouse(screen image.Rectangle) {
	pressed, used := v.op.MouseButton()
	v.mouseUsed = v.mouseUsed || used
	if !v.mouseUsed {
		return
	}
	if pressed {
		v.mouseDown = v.frame
	}
	x := min(videoWidth-panelWidth-8, screen.Max.X+(videoWidth-screen.Max.X-panelWidth)/2)
	y := screen.Min.Y + 40
	drawText(v.canvas, "Mouse", x, y, 1, keyLetter, 1)
	y += 18
	height := panelWidth * 3 / 4
	pad := image.Rect(x, y, x+panelWidth, y+height)
	fill(v.canvas, pad, panelLine, 1)
	fill(v.canvas, pad.Inset(2), color.RGBA{A: 0xff}, 1)
	mx, my := v.op.MousePosition()
	disc(v.canvas, image.Pt(x+4+int(mx*float64(panelWidth-8)), y+4+int(my*float64(height-8))), 5, panelDot)
	y += height + 16
	v.drawButton(image.Pt(x, y), "Button", v.frame-v.mouseDown < litFrames && v.mouseDown > 0)
}

// disc paints a filled circle
func disc(dst *image.RGBA, centre image.Point, radius int, c color.RGBA) {
	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			if dx*dx+dy*dy <= radius*radius {
				dst.SetRGBA(centre.X+dx, centre.Y+dy, c)
			}
		}
	}
}

/*
Close finishes the video: the frames are encoded, the sound is put in after
the title, and the chapters are written next to it, a line each with its
time, as the description of a YouTube video takes them.
*/
func (v *Video) Close() error {
	videos.Lock()
	delete(videos.of, v.op)
	videos.Unlock()
	v.sound.Stop()
	for range endFrames {
		v.write()
	}
	v.frames.Close()
	if err := v.encoder.Wait(); err != nil && v.err == nil {
		v.err = fmt.Errorf("ffmpeg could not encode the frames: %w", err)
	}
	if v.err != nil {
		return v.err
	}

	wav := strings.TrimSuffix(v.path, ".mp4") + ".wav"
	f, err := os.Create(wav)
	if err != nil {
		return err
	}
	if err := v.sound.Encode(f); err != nil {
		f.Close()
		return err
	}
	f.Close()
	delay := titleFrames * 1000 / operator.FramesPerSecond
	mux := exec.Command("ffmpeg", "-y", "-loglevel", "error",
		"-i", v.silentPath(), "-i", wav,
		"-filter_complex", fmt.Sprintf("[1:a]adelay=%d:all=1[a]", delay),
		"-map", "0:v", "-map", "[a]", "-c:v", "copy", "-c:a", "aac", "-b:a", "192k",
		v.path)
	mux.Stderr = os.Stderr
	if err := mux.Run(); err != nil {
		return fmt.Errorf("ffmpeg could not put in the sound: %w", err)
	}
	os.Remove(wav)
	os.Remove(v.silentPath())

	chapters := "0:00 " + v.title + "\n"
	for _, c := range v.chapter {
		seconds := (titleFrames + c.frame) / operator.FramesPerSecond
		chapters += fmt.Sprintf("%d:%02d %s\n", seconds/60, seconds%60, c.name)
	}
	return os.WriteFile(strings.TrimSuffix(v.path, ".mp4")+".chapters.txt", []byte(chapters), 0o644)
}

// markPicture marks a picture an album takes of an operator in its video,
// if it is filmed, at a frame or now
func markPicture(op *operator.Operator, name string, frame int) {
	videos.Lock()
	v := videos.of[op]
	videos.Unlock()
	if v == nil {
		return
	}
	if frame < 0 {
		frame = v.frame
	}
	v.mark(name, frame)
}

// videoFrame is the frame of the video of an operator now, or -1
func videoFrame(op *operator.Operator) int {
	videos.Lock()
	defer videos.Unlock()
	if v := videos.of[op]; v != nil {
		return v.frame
	}
	return -1
}
