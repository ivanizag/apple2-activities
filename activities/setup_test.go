package activities

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"unicode"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

/*
What the activities run on: the disks fetch-disks.sh downloads into disks/,
which a reader gets the same of, copied for each machine so that what it
writes does not change them, and the folder the pictures of each guide go to.
*/

const (
	// guideImages is where the pictures of the guides go, a folder each
	guideImages = "../guides/images"

	// disks is where fetch-disks.sh leaves the disks
	disks = "../disks"
)

// newAlbum is the album of the pictures of a guide
func newAlbum(guide string, monitor album.Monitor) *album.Album {
	return album.New(filepath.Join(guideImages, guide), monitor)
}

// disk is a copy of a disk of disks/, in the test's own directory
func disk(t testing.TB, name string) string {
	t.Helper()

	source, err := os.Open(filepath.Join(disks, name))
	if err != nil {
		t.Fatalf("the disk %v is missing, run ./fetch-disks.sh: %v", name, err)
	}
	defer source.Close()

	copied := filepath.Join(t.TempDir(), name)
	target, err := os.Create(copied)
	if err != nil {
		t.Fatal(err)
	}
	defer target.Close()

	if _, err := io.Copy(target, source); err != nil {
		t.Fatal(err)
	}
	return copied
}

/*
start builds the machine of a command line of izapple2, the one its guide
gives, and sits an operator at it. The disks the command names as disks/<name>,
in the parameters of a card or alone as the recording of -tape, are copies of
the ones fetch-disks.sh downloads; files gives the paths of the others, as the
diskette the reader makes of blank.dsk.
*/
func start(t testing.TB, command string, files map[string]string) *operator.Operator {
	t.Helper()
	model, overrides, positional, err := parseCommand(command)
	if err != nil {
		t.Fatal(err)
	}

	// Each file named, alone or in the parameters of a card, is the copy
	path := func(name string) string {
		if copied, ok := files[name]; ok {
			return copied
		}
		if fetched, ok := strings.CutPrefix(name, "disks/"); ok {
			return disk(t, fetched)
		}
		return name
	}
	for key, value := range overrides {
		params := splitParams(value)
		for i, param := range params {
			name, file, ok := strings.Cut(param, "=")
			if !ok {
				// A file given alone, as the recording of -tape
				params[i] = path(param)
				continue
			}
			if unquoted, ok := strings.CutPrefix(file, `"`); ok {
				params[i] = name + `="` + path(strings.TrimSuffix(unquoted, `"`)) + `"`
			} else {
				params[i] = name + "=" + path(file)
			}
		}
		overrides[key] = strings.Join(params, ",")
	}
	for i, name := range positional {
		positional[i] = path(name)
	}

	o, err := operator.Start(model, overrides, positional...)
	if err != nil {
		t.Fatal(err)
	}
	return o
}

// splitParams splits the parameters of a card at its commas, as izapple2
// does, but not at the ones in double quotes, in the name of a disk
func splitParams(value string) []string {
	var params []string
	quoted := false
	start := 0
	for i, c := range value {
		switch {
		case c == '"':
			quoted = !quoted
		case c == ',' && !quoted:
			params = append(params, value[start:i])
			start = i + 1
		}
	}
	return append(params, value[start:])
}

// izapple2Switches are the options of izapple2 that take no value
var izapple2Switches = []string{"profile", "showConfig", "forceCaps", "rgb", "romx"}

/*
parseCommand reads a command line of izapple2 as the shell and izapple2 would:
the model, the rest of the options, and the files given after them. Words in
quotes, single or double, are one word, without them.
*/
func parseCommand(command string) (model string, options map[string]string, files []string, err error) {
	words, err := splitWords(command)
	if err != nil {
		return "", nil, nil, err
	}
	if len(words) == 0 || words[0] != "izapple2" {
		return "", nil, nil, fmt.Errorf("not a command line of izapple2: %v", command)
	}

	model = "2enh"
	options = map[string]string{}
	words = words[1:]
	for len(words) > 0 && strings.HasPrefix(words[0], "-") {
		name := strings.TrimLeft(words[0], "-")
		words = words[1:]
		if slices.Contains(izapple2Switches, name) {
			options[name] = "true"
			continue
		}
		if len(words) == 0 {
			return "", nil, nil, fmt.Errorf("-%v has no value in %v", name, command)
		}
		if name == "model" {
			model = words[0]
		} else {
			options[name] = words[0]
		}
		words = words[1:]
	}
	return model, options, words, nil
}

// splitWords splits a command line in words, by spaces and line continuations
func splitWords(command string) ([]string, error) {
	command = strings.ReplaceAll(command, "\\\n", " ")
	var words []string
	var word strings.Builder
	inWord := false
	var quote rune
	for _, c := range command {
		switch {
		case quote == 0 && (c == '"' || c == '\''):
			quote = c
			inWord = true
		case c == quote:
			quote = 0
		case quote == 0 && unicode.IsSpace(c):
			if inWord {
				words = append(words, word.String())
				word.Reset()
				inWord = false
			}
		default:
			word.WriteRune(c)
			inWord = true
		}
	}
	if quote != 0 {
		return nil, fmt.Errorf("unclosed quote in %v", command)
	}
	if inWord {
		words = append(words, word.String())
	}
	return words, nil
}

// must stops the activity when the operator could not do what it was asked
func must(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// typeListing types a program of the guides, a line at a time
func typeListing(o *operator.Operator, path string) error {
	text, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(strings.TrimRight(string(text), "\n"), "\n") {
		if err := o.TypeLines(line); err != nil {
			return fmt.Errorf("typing %q: %w", line, err)
		}
	}
	return nil
}

// typeText types a program of the guides whole, as one text, its lines ended
// with Return, into an editor
func typeText(o *operator.Operator, path string) error {
	text, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return o.Type(string(text))
}
