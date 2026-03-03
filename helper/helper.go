package helper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/abdirizak-alaaja/jikan-gin-api/config"
	"github.com/abdirizak-alaaja/jikan-gin-api/jikan"
)

func UrlToStruct(url string, target interface{}) error {
	resp, err := jikan.Client.Get(config.Endpoint + url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(target)
	errClose := resp.Body.Close()
	if errClose != nil {
		return errClose
	}
	if err != nil {
		return err
	}
	return nil
}
