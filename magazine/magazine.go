package magazine

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// Magazines struct
type Magazines struct {
	Data       []common.MalItemCount `json:"data"`
	Pagination common.Pagination     `json:"pagination"`
}

// GetMagazines returns magazines
func GetMagazines(page int) (*Magazines, error) {
	res := &Magazines{}
	err := helper.UrlToStruct(fmt.Sprintf("/magazines?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
