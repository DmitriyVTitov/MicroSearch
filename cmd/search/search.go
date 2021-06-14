// Микросервис MicroSearch.
// Осуществляет поиск музыкальных треков по
// ключевым словам.
//
// Пример URL: http://localhost/search/sting
package main

import (
	"log"
	"net/http"
	"sf/MicroSearch/pkg/api"
	"sf/MicroSearch/pkg/db"
	"sf/MicroSearch/pkg/index"
)

func main() {
	log.Println("Загрузка треков из БД.")
	tracks, err := db.LoadTracks()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Треки успешно загружены.")
	ind := index.New(tracks)
	api := api.New(ind, tracks)
	http.ListenAndServe(":80", api.Router())
}
