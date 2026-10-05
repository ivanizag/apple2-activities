package album

import (
	"encoding/binary"
	"errors"
	"io"
	"math"

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
	sources  []izapple2.AudioSource
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
	s := &Sound{a: a, start: a.GetCycles(), clockMhz: a.GetClockMhz()}
	for _, source := range a.GetAudioSources() {
		sink := &levels{}
		source.SetAudioSink(sink)
		s.sources = append(s.sources, source)
		s.sinks = append(s.sinks, sink)
	}
	return s
}

// Stop stops listening
func (s *Sound) Stop() {
	s.end = s.a.GetCycles()
	for _, source := range s.sources {
		source.SetAudioSink(nil)
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

// samples renders the sound, mono, from -1 to 1
func (s *Sound) samples() []float64 {
	cyclesPerSample := s.clockMhz * 1e6 / SoundRate
	count := int(float64(s.end-s.start) / cyclesPerSample)
	out := make([]float64, count)

	for _, sink := range s.sinks {
		level := 0.0
		next := 0
		for i := range out {
			from := float64(s.start) + float64(i)*cyclesPerSample
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
			out[i] += sum / cyclesPerSample
		}
	}

	// DC-blocking filter, about 30 Hz
	const pole = 1 - 30*2*math.Pi/SoundRate
	prevIn, prevOut := 0.0, 0.0
	for i, sample := range out {
		filtered := sample - prevIn + pole*prevOut
		prevIn, prevOut = sample, filtered
		out[i] = filtered
	}
	return out
}

// Encode writes the sound as a WAV file, 16 bits mono
func (s *Sound) Encode(w io.Writer) error {
	if s.end == 0 {
		return errors.New("the sound was not stopped")
	}
	samples := s.samples()
	data := make([]byte, 2*len(samples))
	for i, sample := range samples {
		v := int16(max(min(sample, 1), -1) * 32767)
		binary.LittleEndian.PutUint16(data[2*i:], uint16(v))
	}

	header := []any{
		[4]byte{'R', 'I', 'F', 'F'}, uint32(36 + len(data)), [4]byte{'W', 'A', 'V', 'E'},
		[4]byte{'f', 'm', 't', ' '}, uint32(16), uint16(1), uint16(1),
		uint32(SoundRate), uint32(SoundRate * 2), uint16(2), uint16(16),
		[4]byte{'d', 'a', 't', 'a'}, uint32(len(data)),
	}
	for _, field := range header {
		if err := binary.Write(w, binary.LittleEndian, field); err != nil {
			return err
		}
	}
	_, err := w.Write(data)
	return err
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
