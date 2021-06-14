package index

import (
	"sf/MicroSearch/pkg/db"
	"testing"
)

func TestIndex_Build(t *testing.T) {
	tracks := []db.Track{
		{ID: 1, Artist: "Madonna", Title: "A power Of Goodbye", Popularity: 5},
		{ID: 2, Artist: "Madonna", Title: "American Dream", Popularity: 5},
	}
	index := New(tracks)
	if len(index) != 5 {
		t.Error("неверный размер индекса")
	}
	if len(index["madonna"]) != 2 {
		t.Error("неверный размер значения в индексе")
	}
}
