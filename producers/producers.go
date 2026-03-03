package producers

import (
	"fmt"

	"github.com/abdirizak-alaaja/jikan-gin-api/common"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
)

// Producers struct
type Producers struct {
	Data       []common.MalItemCount `json:"data"`
	Pagination common.Pagination     `json:"pagination"`
}

// GetProducers returns producers
func GetProducers(page int) (*Producers, error) {
	res := &Producers{}
	err := helper.UrlToStruct(fmt.Sprintf("/producers?page=%d", page), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
