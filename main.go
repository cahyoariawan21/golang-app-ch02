package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Music struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	AlbumTitle string  `json:"album_title"`
	Singer     *Singer `json:"singer"`
}

type Singer struct {
	Name string `json:"name"`
}

// slice of music
var musics []Music

func getAllMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)
}

func getMusicByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// looping to check the data on Musics
	for _, item := range musics {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var music Music
	_ = json.NewDecoder(r.Body).Decode(&music)
	music.ID = strconv.Itoa(rand.Intn(1000000))
	musics = append(musics, music)
	json.NewEncoder(w).Encode(musics)
}

func main() {
	// Your code start here
	r := mux.NewRouter()

	// insert data into slice music
	musics = append(musics, Music{ID: "1", Title: "Amazing", AlbumTitle: "Who Cares?", Singer: &Singer{Name: "Rex Orange County"}})
	musics = append(musics, Music{ID: "2", Title: "Shy Away", AlbumTitle: "Scaled and Icy", Singer: &Singer{Name: "Twenty One Pilots"}})
	// Router
	r.HandleFunc("/musics", getAllMusic).Methods("GET")
	r.HandleFunc("/musics/{id}", getMusicByID).Methods("GET")
	r.HandleFunc("/musics", createMusic).Methods("POST")

	// print log while http server is running on port: 8080
	fmt.Printf("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
