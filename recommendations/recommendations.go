package recommendations

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// RecentRecommendations struct
type RecentRecommendations struct {
	Data []struct {
		MalId   string        `json:"mal_id"`
		Entry   []common.EntryTitle3 `json:"entry"`
		Content string        `json:"content"`
		User    struct {
			Url      string `json:"url"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"data"`
	Pagination common.Pagination `json:"pagination"`
}

// GetRecentAnimeRecommendations returns recent anime recommendations
func GetRecentAnimeRecommendations() (*RecentRecommendations, error) {
	res := &RecentRecommendations{}
	err := helper.UrlToStruct("/recommendations/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRecentMangaRecommendations returns recent manga recommendations
func GetRecentMangaRecommendations() (*RecentRecommendations, error) {
	res := &RecentRecommendations{}
	err := helper.UrlToStruct("/recommendations/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
