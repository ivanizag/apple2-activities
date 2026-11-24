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
under the screen as they are pressed, each fading after a while. The pictures
the activity takes are marked in a list of chapters, with their times, for the
description of the video.

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
