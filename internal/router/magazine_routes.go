package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/magazine"
	"github.com/gin-gonic/gin"
)

// MagazinesRoutes registers /api/v4/magazines endpoints
func MagazinesRoutes(r *gin.Engine) {
	magAPI := r.Group("/api/v4/magazines")

	// GET /api/v4/magazines?page=1
	magAPI.GET("", func(ctx *gin.Context) {
		page := helper.ParsePage(ctx)
		data, err := magazine.GetMagazines(page)
		helper.HandleResult(ctx, data, err)
	})
}
