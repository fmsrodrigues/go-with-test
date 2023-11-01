package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League sorted", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(Database)
		AssertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(Database)

		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScoreEquals(t, got, want)
	})

	t.Run("Store wins for existing players", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(Database)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScoreEquals(t, got, want)
	})

	t.Run("Store wins for new players", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemPlayerStore(Database)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(t, got, want)
	})

	t.Run("Works with an empty file", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(Database)

		AssertNoError(t, err)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
