package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/people"
	"github.com/gin-gonic/gin"
)

// PeopleRoutes registers /api/v4/people endpoints
func PeopleRoutes(r *gin.Engine) {
	peopleAPI := r.Group("/api/v4/people")

	// GET /api/v4/people/:id
	peopleAPI.GET("/:id", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := people.GetPersonById(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/people/:id/anime
	peopleAPI.GET("/:id/anime", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := people.GetPersonAnime(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/people/:id/manga
	peopleAPI.GET("/:id/manga", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := people.GetPersonManga(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/people/:id/voices
	peopleAPI.GET("/:id/voices", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := people.GetPersonVoices(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/people/:id/pictures
	peopleAPI.GET("/:id/pictures", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := people.GetPersonPictures(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/people/search?query=...
	peopleAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := people.GetPeopleSearch(query)
		helper.HandleResult(ctx, data, err)
	})
}