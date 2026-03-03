package genres

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// Genres struct
type Genres struct {
	Data []common.MalItemCount `json:"data"`
}

type GenreFilter string

const (
	GenresFilterGenres         GenreFilter = "genres"
	GenresFilterExplicitGenres GenreFilter = "explicit_genres"
	GenresFilterThemes         GenreFilter = "themes"
	GenresFilterDemographics   GenreFilter = "demographics"
)

// GetAnimeGenres returns anime genres
func GetAnimeGenres(page, limit int, filter GenreFilter) (*Genres, error) {
	res := &Genres{}
	req := fmt.Sprintf("/genres/anime?page=%d&limit=%d", page, limit)
	if filter != "" {
		req += fmt.Sprintf("&filter=%s", filter)
	}
	err := helper.UrlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMangaGenres returns manga genres
func GetMangaGenres(page, limit int, filter GenreFilter) (*Genres, error) {
	res := &Genres{}
	req := fmt.Sprintf("/genres/manga?page=%d&limit=%d", page, limit)
	if filter != "" {
		req += fmt.Sprintf("&filter=%s", filter)
	}
	err := helper.UrlToStruct(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
