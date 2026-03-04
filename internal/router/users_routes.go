package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/users"
	"github.com/gin-gonic/gin"

)

// UsersRoutes registers /api/v4/users endpoints
func UsersRoutes(r *gin.Engine) {
	userAPI := r.Group("/api/v4/users")

	// GET /api/v4/users/search?query=...
	userAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := users.GetUsersSearch(query)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username
	userAPI.GET("/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserProfile(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/statistics
	userAPI.GET("/:username/statistics", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserStatistics(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/favorites
	userAPI.GET("/:username/favorites", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserFavorites(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/userupdates
	userAPI.GET("/:username/userupdates", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserUpdates(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/about
	userAPI.GET("/:username/about", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserAbout(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/history?type=anime|manga
	userAPI.GET("/:username/history", func(ctx *gin.Context) {
		username := ctx.Param("username")
		filter := users.UserHistoryFilter(ctx.Query("type"))
		data, err := users.GetUserHistory(username, filter)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/friends
	userAPI.GET("/:username/friends", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserFriends(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/reviews
	userAPI.GET("/:username/reviews", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserReviews(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/recommendations
	userAPI.GET("/:username/recommendations", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserRecommendations(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/clubs
	userAPI.GET("/:username/clubs", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserClubs(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/animelist
	userAPI.GET("/:username/animelist", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserAnimelist(username)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/users/:username/mangalist
	userAPI.GET("/:username/mangalist", func(ctx *gin.Context) {
		username := ctx.Param("username")
		data, err := users.GetUserMangalist(username)
		helper.HandleResult(ctx, data, err)
	})
}