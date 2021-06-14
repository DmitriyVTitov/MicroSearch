// База данных с каталогом треков.
package db

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Музыкальный трек.
type Track struct {
	ID         int
	Artist     string
	Title      string
	Popularity int
}

// Возвращает все треки из БД.
func LoadTracks() ([]Track, error) {
	var tracks []Track

	f, err := os.Open("./tracks.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	id := 0
	for {
		recs, err := reader.Read()
		if err != nil {
			break
		}
		var track Track
		if len(recs) < 6 {
			continue
		}
		track.ID = id
		id++
		track.Title = recs[1]
		p, _ := strconv.Atoi(recs[2])
		track.Popularity = p
		track.Artist = recs[5]
		tracks = append(tracks, track)
	}
	return tracks, nil
}
