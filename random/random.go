package random

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
	"github.com/abdirizak-alaaja/jikan-gin-api/characters"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/manga"
	"github.com/abdirizak-alaaja/jikan-gin-api/people"
	"github.com/abdirizak-alaaja/jikan-gin-api/users"
)

// RandomAnime struct
type RandomAnime struct {
	Data anime.AnimeBase `json:"data"`
}

// GetRandomAnime returns random anime
func GetRandomAnime() (*RandomAnime, error) {
	res := &RandomAnime{}
	err := helper.UrlToStruct("/random/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomManga struct
type RandomManga struct {
	Data manga.MangaBase `json:"data"`
}

// GetRandomManga returns random manga
func GetRandomManga() (*RandomManga, error) {
	res := &RandomManga{}
	err := helper.UrlToStruct("/random/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomCharacters struct
type RandomCharacters struct {
	Data characters.CharactersBase `json:"data"`
}

// GetRandomCharacters returns random characters
func GetRandomCharacters() (*RandomCharacters, error) {
	res := &RandomCharacters{}
	err := helper.UrlToStruct("/random/characters", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomPeople struct
type RandomPeople struct {
	Data people.PeopleBase `json:"data"`
}

// GetRandomPeople returns random people
func GetRandomPeople() (*RandomPeople, error) {
	res := &RandomPeople{}
	err := helper.UrlToStruct("/random/people", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RandomUsers struct
type RandomUsers struct {
	Data users.UsersBase `json:"data"`
}

// GetRandomUsers returns random users
func GetRandomUsers() (*RandomUsers, error) {
	res := &RandomUsers{}
	err := helper.UrlToStruct("/random/users", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
