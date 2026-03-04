package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/random"
	"github.com/gin-gonic/gin"
)

// RandomRoutes registers /api/v4/random endpoints
func RandomRoutes(r *gin.Engine) {
	randAPI := r.Group("/api/v4/random")

	randAPI.GET("/anime", func(ctx *gin.Context) {
		data, err := random.GetRandomAnime()
		helper.HandleResult(ctx, data, err)
	})

	randAPI.GET("/manga", func(ctx *gin.Context) {
		data, err := random.GetRandomManga()
		helper.HandleResult(ctx, data, err)
	})

	randAPI.GET("/characters", func(ctx *gin.Context) {
		data, err := random.GetRandomCharacters()
		helper.HandleResult(ctx, data, err)
	})

	randAPI.GET("/people", func(ctx *gin.Context) {
		data, err := random.GetRandomPeople()
		helper.HandleResult(ctx, data, err)
	})

	randAPI.GET("/users", func(ctx *gin.Context) {
		data, err := random.GetRandomUsers()
		helper.HandleResult(ctx, data, err)
	})
}