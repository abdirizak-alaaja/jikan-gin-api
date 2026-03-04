package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/recommendations"
	"github.com/gin-gonic/gin"
)

// RecommendationsRoutes registers /api/v4/recommendations endpoints
func RecommendationsRoutes(r *gin.Engine) {
	recAPI := r.Group("/api/v4/recommendations")

	// GET /api/v4/recommendations/anime
	recAPI.GET("/anime", func(ctx *gin.Context) {
		data, err := recommendations.GetRecentAnimeRecommendations()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/recommendations/manga
	recAPI.GET("/manga", func(ctx *gin.Context) {
		data, err := recommendations.GetRecentMangaRecommendations()
		helper.HandleResult(ctx, data, err)
	})
}