package reviews

import (
	"time"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// RecentAnimeReviews struct
type RecentAnimeReviews struct {
	Data []struct {
		User            common.UserItem    `json:"user"`
		Anime           common.EntryTitle3 `json:"anime"`
		MalId           int         `json:"mal_id"`
		Url             string      `json:"url"`
		Type            string      `json:"type"`
		Votes           int         `json:"votes"`
		Date            time.Time   `json:"date"`
		Review          string      `json:"review"`
		EpisodesWatched int         `json:"episodes_watched"`
		Scores          common.ScoresAnime `json:"scores"`
	} `json:"data"`
	Pagination common.Pagination `json:"pagination"`
}

// GetRecentAnimeReviews returns recent anime reviews
func GetRecentAnimeReviews() (*RecentAnimeReviews, error) {
	res := &RecentAnimeReviews{}
	err := helper.UrlToStruct("/reviews/anime", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RecentMangaReviews struct
type RecentMangaReviews struct {
	Data []struct {
		User         common.UserItem    `json:"user"`
		Manga        common.EntryTitle3 `json:"manga"`
		MalId        int         `json:"mal_id"`
		Url          string      `json:"url"`
		Type         string      `json:"type"`
		Votes        int         `json:"votes"`
		Date         time.Time   `json:"date"`
		ChaptersRead int         `json:"chapters_read"`
		Review       string      `json:"review"`
		Scores       common.ScoresManga `json:"scores"`
	} `json:"data"`
	Pagination common.Pagination `json:"pagination"`
}

// GetRecentMangaReviews returns recent manga reviews
func GetRecentMangaReviews() (*RecentMangaReviews, error) {
	res := &RecentMangaReviews{}
	err := helper.UrlToStruct("/reviews/manga", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
