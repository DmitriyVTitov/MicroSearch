// Обратный индекс с каталогом треков.
package index

import (
	"sf/MicroSearch/pkg/db"
	"strings"
)

// Обратный индекс поиска.
type Index map[string][]int

// Конструктор.
func New(tracks []db.Track) Index {
	ind := Index(make(map[string][]int))
	ind.Build(tracks)
	return ind
}

// Построение индекса из массива треков.
func (ind Index) Build(tracks []db.Track) {
	for _, track := range tracks {
		s := strings.ReplaceAll(track.Artist, "']", "")
		s = strings.ReplaceAll(s, "['", "")
		tokens := strings.Split(s, " ")
		tokens = append(tokens, strings.Split(track.Title, " ")...)
		for _, token := range tokens {
			if len(token) < 4 {
				continue
			}
			ind[strings.ToLower(token)] = append(ind[strings.ToLower(token)], track.ID)
		}
	}
}

// Поиск по индексу.
func (ind Index) Search(token string) []int {
	return ind[token]
}
