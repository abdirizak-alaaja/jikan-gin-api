package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/watch"
	"github.com/gin-gonic/gin"
)

// WatchRoutes registers /api/v4/watch endpoints
func WatchRoutes(r *gin.Engine) {
	watchAPI := r.Group("/api/v4/watch")

	// GET /api/v4/watch/episodes
	watchAPI.GET("/episodes", func(ctx *gin.Context) {
		data, err := watch.GetWatchRecentEpisodes()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/watch/episodes/popular
	watchAPI.GET("/episodes/popular", func(ctx *gin.Context) {
		data, err := watch.GetWatchPopularEpisodes()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/watch/promos
	watchAPI.GET("/promos", func(ctx *gin.Context) {
		data, err := watch.GetWatchRecentPromos()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/watch/promos/popular
	watchAPI.GET("/promos/popular", func(ctx *gin.Context) {
		data, err := watch.GetWatchPopularPromos()
		helper.HandleResult(ctx, data, err)
	})
}