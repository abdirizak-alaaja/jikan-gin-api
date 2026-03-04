package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/producers"
	"github.com/gin-gonic/gin"
)

// ProducersRoutes registers /api/v4/producers endpoints
func ProducersRoutes(r *gin.Engine) {
	prodAPI := r.Group("/api/v4/producers")

	// GET /api/v4/producers?page=1
	prodAPI.GET("", func(ctx *gin.Context) {
		page := helper.ParsePage(ctx)

		data, err := producers.GetProducers(page)
		helper.HandleResult(ctx, data, err)
	})
}
