package main_test

import (
	"io"
	"os"
	"testing"

	http_server "applications/http-server"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := http_server.FileSystemPlayerStore{Database}

		got := store.GetLeague()
		want := []http_server.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := http_server.FileSystemPlayerStore{Database}

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("Store wins for existing players", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store := http_server.FileSystemPlayerStore{Database}
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
