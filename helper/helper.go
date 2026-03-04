package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/abdirizak-alaaja/jikan-gin-api/config"
	"github.com/abdirizak-alaaja/jikan-gin-api/jikan"
	"github.com/gin-gonic/gin"
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

// helper to parse id param
func ParseID(ctx *gin.Context) (int, bool) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

// helper to parse page query
func ParsePage(ctx *gin.Context) int {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 1
	}
	return page
}

// helper to handle results
func HandleResult(ctx *gin.Context, data any, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// ParseLimit parses the ?limit= query param (defaults to 25)
func ParseLimit(ctx *gin.Context) int {
	limitStr := ctx.DefaultQuery("limit", "25")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return 25
	}
	return limit
}
