package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/top"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TopRoutes registers /api/v4/top endpoints
func TopRoutes(r *gin.Engine) {
	topAPI := r.Group("/api/v4/top")

	// GET /api/v4/top/anime
	topAPI.GET("/anime", func(ctx *gin.Context) {
		subType := ctx.Query("type")
		filter := ctx.Query("filter")
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

		data, err := top.GetTopAnime(top.TopAnimeType(subType), top.TopAnimeFilter(filter), page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/top/manga
	topAPI.GET("/manga", func(ctx *gin.Context) {
		subType := ctx.Query("type")
		filter := ctx.Query("filter")
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

		data, err := top.GetTopManga(top.TopMangaType(subType), top.TopMangaFilter(filter), page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/top/people
	topAPI.GET("/people", func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		data, err := top.GetTopPeople(page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/top/characters
	topAPI.GET("/characters", func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		data, err := top.GetTopCharacters(page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/top/reviews
	topAPI.GET("/reviews", func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		data, err := top.GetTopReviews(page)
		helper.HandleResult(ctx, data, err)
	})
}