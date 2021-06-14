package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sf/MicroSearch/pkg/db"
	"sf/MicroSearch/pkg/index"
	"testing"
)

func TestAPI_searchHandler(t *testing.T) {

	tracks := []db.Track{
		{ID: 1, Artist: "Madonna", Title: "A power Of Goodbye", Popularity: 5},
		{ID: 1, Artist: "Madonna", Title: "American Dream", Popularity: 5},
	}
	index := index.New(tracks)
	api := New(index, tracks)
	req := httptest.NewRequest(http.MethodGet, "/search/Madonna", nil)
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	t.Logf("code: %d", rr.Code)
	b, _ := io.ReadAll(rr.Body)
	t.Logf("%v", string(b))
}
