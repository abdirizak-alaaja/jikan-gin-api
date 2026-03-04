package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/reviews"
	"github.com/gin-gonic/gin"
)

// ReviewsRoutes registers /api/v4/reviews endpoints
func ReviewsRoutes(r *gin.Engine) {
	reviewAPI := r.Group("/api/v4/reviews")

	// GET /api/v4/reviews/anime
	reviewAPI.GET("/anime", func(ctx *gin.Context) {
		data, err := reviews.GetRecentAnimeReviews()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/reviews/manga
	reviewAPI.GET("/manga", func(ctx *gin.Context) {
		data, err := reviews.GetRecentMangaReviews()
		helper.HandleResult(ctx, data, err)
	})
}