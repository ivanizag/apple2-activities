package album

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"sync"

	"github.com/ivanizag/izapple2"

	"github.com/ivanizag/apple2-activities/operator"
)

/*
A Sound is what the machine plays while it is listened to, kept as the level
changes of each of its sound generators, the speaker and the sound cards, with
the cycle of each. It is made into a WAV file afterwards, the same on every
run, with the same box filter and the same DC-blocking filter as the mixer of
the izapple2 frontends: the level of each generator averaged over each sample,
and the constant levels decaying to silence as on a real speaker.
*/
type Sound struct {
	a        *izapple2.Apple2
	sinks    []*levels
	start    uint64
	end      uint64
	clockMhz float64
}

// levels is the AudioSink of one sound generator, every change it made
type levels struct {
	changes []levelChange
}

type levelChange struct {
	cycle uint64
	level float32
}

// PushLevel keeps a level change of the generator
func (l *levels) PushLevel(cycle uint64, level float32) {
	l.changes = append(l.changes, levelChange{cycle, level})
}

// SoundRate is the sample rate of the WAV files, enough for the speaker and
// the Mockingboard and half the size of CD quality
const SoundRate = 22050

/*
Listen starts keeping what the machine an operator is at plays, from now until
Stop. It has to be called before the machine runs: a sound card does not
work out its sound while nobody listens, and when it starts being listened to
late it has all that time to catch up with first. Clip takes the part of it
to keep.
*/
func Listen(op *operator.Operator) *Sound {
	a := op.Apple2()
	listening.Lock()
	defer listening.Unlock()
	l := listening.of[a]
	if l == nil {
		// The first to listen connects the generators; the others share them
		l = &listener{}
		for _, source := range a.GetAudioSources() {
			sink := &levels{}
			source.SetAudioSink(sink)
			l.sources = append(l.sources, source)
			l.sinks = append(l.sinks, sink)
		}
		listening.of[a] = l
	}
	l.count++
	return &Sound{a: a, sinks: l.sinks, start: a.GetCycles(), clockMhz: a.GetClockMhz()}
}

/*
A listener is the connection of the sound generators of a machine, shared by
all who listen to it at the same time, as the video of a whole page and the
clip of one of its steps: a generator has only one sink.
*/
type listener struct {
	sources []izapple2.AudioSource
	sinks   []*levels
	count   int
}

// listening is the listener of each machine listened to
var listening = struct {
	sync.Mutex
	of map[*izapple2.Apple2]*listener
}{of: map[*izapple2.Apple2]*listener{}}

// Stop stops listening; the generators are let go when nobody listens
func (s *Sound) Stop() {
	if s.end != 0 {
		return
	}
	s.end = s.a.GetCycles()
	listening.Lock()
	defer listening.Unlock()
	l := listening.of[s.a]
	if l == nil {
		return
	}
	l.count--
	if l.count == 0 {
		for _, source := range l.sources {
			source.SetAudioSink(nil)
		}
		delete(listening.of, s.a)
	}
}

// Now is the cycle of the machine now, to mark where a clip starts or ends
func (s *Sound) Now() uint64 {
	return s.a.GetCycles()
}

// Clip is the part of the sound from one cycle to another, as marked with
// Now, to save on its own while the machine goes on
func (s *Sound) Clip(from uint64, to uint64) *Sound {
	return &Sound{sinks: s.sinks, start: from, end: to, clockMhz: s.clockMhz}
}

// Seconds is how long the sound lasts, in seconds of the machine
func (s *Sound) Seconds() float64 {
	return float64(s.end-s.start) / (s.clockMhz * 1e6)
}

// samplesChunk is how many samples are rendered at a time, so that a sound
// of hours, as a video of a whole page, is not held in memory at once
const samplesChunk = 1 << 16

// count is how many samples the sound lasts
func (s *Sound) count() int {
	return int(float64(s.end-s.start) / (s.clockMhz * 1e6 / SoundRate))
}

// samples renders the sound, mono, from -1 to 1, a chunk at a time, each
// given to emit
func (s *Sound) samples(emit func([]float64) error) error {
	cyclesPerSample := s.clockMhz * 1e6 / SoundRate
	count := s.count()

	// Where each generator is, from one chunk to the next
	levels := make([]float64, len(s.sinks))
	nexts := make([]int, len(s.sinks))
	// DC-blocking filter, about 30 Hz
	const pole = 1 - 30*2*math.Pi/SoundRate
	prevIn, prevOut := 0.0, 0.0

	out := make([]float64, samplesChunk)
	for first := 0; first < count; first += samplesChunk {
		chunk := out[:min(samplesChunk, count-first)]
		clear(chunk)
		for k, sink := range s.sinks {
			level, next := levels[k], nexts[k]
			for j := range chunk {
				from := float64(s.start) + float64(first+j)*cyclesPerSample
				to := from + cyclesPerSample
				// The average of the level over the window of the sample
				sum := 0.0
				at := from
				for next < len(sink.changes) && float64(sink.changes[next].cycle) < to {
					change := float64(sink.changes[next].cycle)
					if change > at {
						sum += level * (change - at)
						at = change
					}
					level = float64(sink.changes[next].level)
					next++
				}
				sum += level * (to - at)
				chunk[j] += sum / cyclesPerSample
			}
			levels[k], nexts[k] = level, next
		}
		for j, sample := range chunk {
			filtered := sample - prevIn + pole*prevOut
			prevIn, prevOut = sample, filtered
			chunk[j] = filtered
		}
		if err := emit(chunk); err != nil {
			return err
		}
	}
	return nil
}

// Encode writes the sound as a WAV file, 16 bits mono
func (s *Sound) Encode(w io.Writer) error {
	if s.end == 0 {
		return errors.New("the sound was not stopped")
	}
	size := 2 * s.count()
	header := []any{
		[4]byte{'R', 'I', 'F', 'F'}, uint32(36 + size), [4]byte{'W', 'A', 'V', 'E'},
		[4]byte{'f', 'm', 't', ' '}, uint32(16), uint16(1), uint16(1),
		uint32(SoundRate), uint32(SoundRate * 2), uint16(2), uint16(16),
		[4]byte{'d', 'a', 't', 'a'}, uint32(size),
	}
	for _, field := range header {
		if err := binary.Write(w, binary.LittleEndian, field); err != nil {
			return err
		}
	}
	data := make([]byte, 2*samplesChunk)
	return s.samples(func(chunk []float64) error {
		for i, sample := range chunk {
			v := int16(max(min(sample, 1), -1) * 32767)
			binary.LittleEndian.PutUint16(data[2*i:], uint16(v))
		}
		_, err := w.Write(data[:2*len(chunk)])
		return err
	})
}

// SaveSound writes a sound of the album as a WAV file
func (a *Album) SaveSound(s *Sound, name string) error {
	f, err := a.create(name + ".wav")
	if err != nil {
		return err
	}
	if err := s.Encode(f); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
