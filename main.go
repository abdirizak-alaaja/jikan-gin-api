package main

import (
	"fmt"
	"net/url"

	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
)

func main() {
	fmt.Print("Starter jikan jin api...")

	// Setup query
	query := url.Values{}
	query.Set("q", "solo leveling")
	query.Set("type", "tv")

	// Search anime
	search, err := anime.GetAnimeSearch(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(search.Data[0])
}
