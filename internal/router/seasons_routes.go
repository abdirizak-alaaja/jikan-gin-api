package router

import (
	"strconv"

	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/seasons"
	"github.com/gin-gonic/gin"
)

// SeasonsRoutes registers /api/v4/seasons endpoints
func SeasonsRoutes(r *gin.Engine) {
	seasonAPI := r.Group("/api/v4/seasons")

	// GET /api/v4/seasons/now
	seasonAPI.GET("/now", func(ctx *gin.Context) {
		data, err := seasons.GetSeasonNow()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/seasons/upcoming
	seasonAPI.GET("/upcoming", func(ctx *gin.Context) {
		data, err := seasons.GetSeasonUpcoming()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/seasons/list
	seasonAPI.GET("/list", func(ctx *gin.Context) {
		data, err := seasons.GetSeasonsList()
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/seasons/:year/:season
	seasonAPI.GET("/:year/:season", func(ctx *gin.Context) {
		yearStr := ctx.Param("year")
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid year"})
			return
		}
		season := ctx.Param("season")
		data, err := seasons.GetSeason(year, season)
		helper.HandleResult(ctx, data, err)
	})
}
