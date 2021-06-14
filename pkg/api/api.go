// Программный интерфейс взаимодействия.
package api

import (
	"encoding/json"
	"net/http"

	"sf/MicroSearch/pkg/db"
	"sf/MicroSearch/pkg/index"

	"github.com/gorilla/mux"
)

// Программный интерфейс взаимодействия.
type API struct {
	r      *mux.Router
	ind    index.Index
	tracks []db.Track
}

// Конструктор API.
func New(ind index.Index, tracks []db.Track) *API {
	api := API{
		r:      mux.NewRouter(),
		ind:    ind,
		tracks: tracks,
	}
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.r.HandleFunc("/search/{token}", api.searchHandler)
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) searchHandler(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	ids := api.ind.Search(token)
	var tracks []db.Track
	for _, track := range api.tracks {
		for _, id := range ids {
			if track.ID == id {
				tracks = append(tracks, track)
			}
		}
	}
	b, _ := json.Marshal(tracks)
	w.Write(b)
}
