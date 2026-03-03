package clubs

import (
	"fmt"
	"net/url"
	"time"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// ClubsBase struct
type ClubsBase struct {
	MalId  int    `json:"mal_id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Images struct {
		Jpg struct {
			ImageUrl string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Members  int       `json:"members"`
	Category string    `json:"category"`
	Created  time.Time `json:"created"`
	Access   string    `json:"access"`
}

// ClubsById struct
type ClubsById struct {
	Data ClubsBase `json:"data"`
}

// GetClubsById returns club by id
func GetClubsById(id int) (*ClubsById, error) {
	res := &ClubsById{}
	err := helper.UrlToStruct(fmt.Sprintf("/clubs/%d", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ClubMembers struct {
	Data       []common.UserItem `json:"data"`
	Pagination common.Pagination      `json:"pagination"`
}

// GetClubMembers returns club members
func GetClubMembers(id, page int) (*ClubMembers, error) {
	res := &ClubMembers{}
	err := helper.UrlToStruct(fmt.Sprintf("/clubs/%d/members?page=%d", id, page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubStaff struct
type ClubStaff struct {
	Data []struct {
		Url      string `json:"url"`
		Username string `json:"username"`
	} `json:"data"`
}

// GetClubStaff returns club staff
func GetClubStaff(id int) (*ClubStaff, error) {
	res := &ClubStaff{}
	err := helper.UrlToStruct(fmt.Sprintf("/clubs/%d/staff", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubRelations struct
type ClubRelations struct {
	Data struct {
		Anime      []common.MalItem `json:"anime"`
		Manga      []common.MalItem `json:"manga"`
		Characters []common.MalItem `json:"characters"`
	} `json:"data"`
}

// GetClubRelations returns club relations
func GetClubRelations(id int) (*ClubRelations, error) {
	res := &ClubRelations{}
	err := helper.UrlToStruct(fmt.Sprintf("/clubs/%d/relations", id), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClubsSearch struct
type ClubsSearch struct {
	Data       []ClubsBase `json:"data"`
	Pagination common.Pagination      `json:"pagination"`
}

// GetClubsSearch returns clubs search
func GetClubsSearch(query url.Values) (*ClubsSearch, error) {
	res := &ClubsSearch{}
	err := helper.UrlToStruct(fmt.Sprintf("/clubs?%s", query.Encode()), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
