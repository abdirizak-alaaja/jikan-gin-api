package seasons

import (
	"fmt"
	"net/url"

	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// Season struct
type Season struct {
	Data       []anime.AnimeBase `json:"data"`
	Pagination common.Pagination  `json:"pagination"`
}

// GetSeason returns season
func GetSeason(year int, season string) (*Season, error) {
	res := &Season{}
	err := helper.UrlToStruct(fmt.Sprintf("/seasons/%d/%s", year, url.QueryEscape(season)), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSeasonNow returns season now
func GetSeasonNow() (*Season, error) {
	res := &Season{}
	err := helper.UrlToStruct("/seasons/now", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SeasonsList struct {
	Data []struct {
		Year    int      `json:"year"`
		Seasons []string `json:"seasons"`
	} `json:"data"`
}

// GetSeasonsList returns seasons list
func GetSeasonsList() (*SeasonsList, error) {
	res := &SeasonsList{}
	err := helper.UrlToStruct("/seasons", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSeasonUpcoming returns season upcoming
func GetSeasonUpcoming() (*Season, error) {
	res := &Season{}
	err := helper.UrlToStruct("/seasons/upcoming", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
