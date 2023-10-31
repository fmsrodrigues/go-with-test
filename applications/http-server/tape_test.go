package poker

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	File, clean := createTempFile(t, "12345")
	defer clean()

	tape := &Tape{File}

	tape.Write([]byte("abc"))

	File.Seek(0, 0)
	newFileContents, _ := io.ReadAll(File)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
