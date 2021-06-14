package db

import (
	"testing"
)

func TestLoadTracks(t *testing.T) {
	tracks, err := LoadTracks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("получено %d записей", len(tracks))
}
