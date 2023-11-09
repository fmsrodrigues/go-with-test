package poker_test

import (
	poker "applications/http-server"
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

		store, err := poker.NewFileSystemPlayerStore(Database)
		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		poker.AssertLeague(t, got, want)

		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("Get player score", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := poker.NewFileSystemPlayerStore(Database)

		got := store.GetPlayerScore("Chris")
		want := 33
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("Store wins for existing players", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := poker.NewFileSystemPlayerStore(Database)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("Store wins for new players", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, _ := poker.NewFileSystemPlayerStore(Database)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("Works with an empty file", func(t *testing.T) {
		Database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(Database)

		poker.AssertNoError(t, err)
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
