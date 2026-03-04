package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/characters"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/gin-gonic/gin"
)

// CharactersRoutes registers /api/v4/characters endpoints
func CharactersRoutes(r *gin.Engine) {

	charAPI := r.Group("/api/v4/characters")

	// GET /api/v4/characters/:id
	charAPI.GET("/:id", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := characters.GetCharacterById(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/characters/:id/anime
	charAPI.GET("/:id/anime", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := characters.GetCharacterAnime(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/characters/:id/manga
	charAPI.GET("/:id/manga", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := characters.GetCharacterManga(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/characters/:id/voices
	charAPI.GET("/:id/voices", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := characters.GetCharacterVoiceActors(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/characters/:id/pictures
	charAPI.GET("/:id/pictures", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := characters.GetCharacterPictures(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/characters/search?query=...
	charAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := characters.GetCharactersSearch(query)
		helper.HandleResult(ctx, data, err)
	})
}
