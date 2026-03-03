package top

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
	"github.com/abdirizak-alaaja/jikan-gin-api/characters"
	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/manga"
	"github.com/abdirizak-alaaja/jikan-gin-api/people"
)

// TopAnime struct
type TopAnime struct {
	Data       []anime.AnimeBase `json:"data"`
	Pagination common.Pagination  `json:"pagination"`
}

type TopAnimeType string

const (
	TopAnimeTypeTv      TopAnimeType = "tv"
	TopAnimeTypeMovie   TopAnimeType = "movie"
	TopAnimeTypeOva     TopAnimeType = "ova"
	TopAnimeTypeSpecial TopAnimeType = "special"
	TopAnimeTypeOna     TopAnimeType = "ona"
	TopAnimeTypeMusic   TopAnimeType = "music"
)

type TopAnimeFilter string

const (
	TopAnimeFilterAiring       TopAnimeFilter = "airing"
	TopAnimeFilterUpcoming     TopAnimeFilter = "upcoming"
	TopAnimeFilterByPopularity TopAnimeFilter = "bypopularity"
	TopAnimeFilterFavorite     TopAnimeFilter = "favorite"
)

// GetTopAnime returns top anime
func GetTopAnime(subType TopAnimeType, filter TopAnimeFilter, page int) (*TopAnime, error) {
	res := &TopAnime{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(page))
	if subType != "" {
		query.Set("type", string(subType))
	}
	if filter != "" {
		query.Set("filter", string(filter))
	}
	err := helper.UrlToStruct(fmt.Sprintf("/top/anime?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopManga struct
type TopManga struct {
	Data       []manga.MangaBase `json:"data"`
	Pagination common.Pagination  `json:"pagination"`
}

type TopMangaType string

const (
	TopMangaTypeManga      TopMangaType = "manga"
	TopMangaTypeNodel      TopMangaType = "novel"
	TopMangaTypeLightnovel TopMangaType = "lightnovel"
	TopMangaTypeOneshot    TopMangaType = "oneshot"
	TopMangaTypeDoujin     TopMangaType = "doujin"
	TopMangaTypeManhwa     TopMangaType = "manhwa"
	TopMangaTypeManhua     TopMangaType = "manhua"
)

type TopMangaFilter string

const (
	TopMangaFilterPublishing   TopMangaFilter = "publishing"
	TopMangaFilterUpcoming     TopMangaFilter = "upcoming"
	TopMangaFilterByPopularity TopMangaFilter = "bypopularity"
	TopMangaFilterFavorite     TopMangaFilter = "favorite"
)

// GetTopManga returns top manga
func GetTopManga(subType TopMangaType, filter TopMangaFilter, page int) (*TopManga, error) {
	res := &TopManga{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(page))
	if subType != "" {
		query.Set("type", string(subType))
	}
	if filter != "" {
		query.Set("filter", string(filter))
	}
	err := helper.UrlToStruct(fmt.Sprintf("/top/manga?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopPeople struct
type TopPeople struct {
	Data       []people.PeopleBase `json:"data"`
	Pagination common.Pagination   `json:"pagination"`
}

// GetTopPeople returns top people
func GetTopPeople(page int) (*TopPeople, error) {
	res := &TopPeople{}
	err := helper.UrlToStruct(fmt.Sprintf("/top/people?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopCharacters struct
type TopCharacters struct {
	Data       []characters.CharactersBase `json:"data"`
	Pagination common.Pagination       `json:"pagination"`
}

// GetTopCharacters returns top characters
func GetTopCharacters(page int) (*TopCharacters, error) {
	res := &TopCharacters{}
	err := helper.UrlToStruct(fmt.Sprintf("/top/characters?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TopReviews struct
type TopReviews struct {
	Data []struct {
		MalId           int         `json:"mal_id"`
		Url             string      `json:"url"`
		Type            string      `json:"type"`
		Votes           int         `json:"votes"`
		Date            time.Time   `json:"date"`
		Review          string      `json:"review"`
		EpisodesWatched int         `json:"episodes_watched"`
		Scores          common.ScoresLong  `json:"scores"`
		Entry           common.EntryTitle3 `json:"entry"`
		User            common.UserItem    `json:"user"`
		ChaptersRead    int         `json:"chapters_read"`
	} `json:"data"`
	Pagination common.Pagination `json:"pagination"`
}

// GetTopReviews returns top reviews
func GetTopReviews(page int) (*TopReviews, error) {
	res := &TopReviews{}
	err := helper.UrlToStruct(fmt.Sprintf("/top/reviews?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
