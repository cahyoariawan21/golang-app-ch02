package main

type Music struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	AlbumTitle string  `json:"album_title"`
	Singer     *Singer `json:"singer"`
}

type Singer struct {
	Name string `json:"name"`
}

var music []Music

func main() {
	// Your code start here
}
