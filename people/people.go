package people

import (
	"fmt"
	"net/url"
	"time"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// PeopleBase struct
type PeopleBase struct {
	MalId      int    `json:"mal_id"`
	Url        string `json:"url"`
	WebsiteUrl string `json:"website_url"`
	Images     struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Name           string    `json:"name"`
	GivenName      string    `json:"given_name"`
	FamilyName     string    `json:"family_name"`
	AlternateNames []string  `json:"alternate_names"`
	Birthday       time.Time `json:"birthday"`
	Favorites      int       `json:"favorites"`
	About          string    `json:"about"`
}

// PersonById struct
type PersonById struct {
	Data PeopleBase `json:"data"`
}

// GetPersonById returns person by id
func GetPersonById(id int) (*PersonById, error) {
	res := &PersonById{}
	err := helper.UrlToStruct(fmt.Sprintf("/people/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonAnime struct
type PersonAnime struct {
	Data []struct {
		Position string      `json:"position"`
		Anime    common.EntryTitle3 `json:"anime"`
	} `json:"data"`
}

// GetPersonAnime returns person anime
func GetPersonAnime(id int) (*PersonAnime, error) {
	res := &PersonAnime{}
	err := helper.UrlToStruct(fmt.Sprintf("/people/%d/anime", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonVoices struct
type PersonVoices struct {
	Data []struct {
		Role      string      `json:"role"`
		Anime     common.EntryTitle3 `json:"anime"`
		Character common.EntryName2  `json:"character"`
	} `json:"data"`
}

// GetPersonVoices returns person voices
func GetPersonVoices(id int) (*PersonVoices, error) {
	res := &PersonVoices{}
	err := helper.UrlToStruct(fmt.Sprintf("/people/%d/voices", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonManga struct
type PersonManga struct {
	Data []struct {
		Position string      `json:"position"`
		Manga    common.EntryTitle3 `json:"manga"`
	} `json:"data"`
}

// GetPersonManga returns person manga
func GetPersonManga(id int) (*PersonManga, error) {
	res := &PersonManga{}
	err := helper.UrlToStruct(fmt.Sprintf("/people/%d/manga", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PersonPictures struct
type PersonPictures struct {
	Data []struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"data"`
}

// GetPersonPictures returns person pictures
func GetPersonPictures(id int) (*PersonPictures, error) {
	res := &PersonPictures{}
	err := helper.UrlToStruct(fmt.Sprintf("/people/%d/pictures", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PeopleSearch struct
type PeopleSearch struct {
	Data       []PeopleBase `json:"data"`
	Pagination common.Pagination   `json:"pagination"`
}

// GetPeopleSearch returns people search
func GetPeopleSearch(query url.Values) (*PeopleSearch, error) {
	res := &PeopleSearch{}
	err := helper.UrlToStruct(fmt.Sprintf("/people?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
