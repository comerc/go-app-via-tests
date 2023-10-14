package poker_test

import (
	"io"
	"testing"

	poker "github.com/comerc/go-app-via-tests"
)

func TestTapeWrite(t *testing.T) {
	file, clean := mustCreateTempFile(t, "12345")
	defer clean()
	tape := &poker.Tape{file}
	tape.Write([]byte("abc"))
	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
