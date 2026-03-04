package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/clubs"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/gin-gonic/gin"
)

// ClubsRoutes registers /api/v4/clubs endpoints
func ClubsRoutes(r *gin.Engine) {
	clubAPI := r.Group("/api/v4/clubs")

	// GET /api/v4/clubs/:id
	clubAPI.GET("/:id", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := clubs.GetClubsById(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/clubs/:id/members?page=1
	clubAPI.GET("/:id/members", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		page := helper.ParsePage(ctx)
		data, err := clubs.GetClubMembers(id, page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/clubs/:id/staff
	clubAPI.GET("/:id/staff", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := clubs.GetClubStaff(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/clubs/:id/relations
	clubAPI.GET("/:id/relations", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := clubs.GetClubRelations(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/clubs/search?query=...
	clubAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := clubs.GetClubsSearch(query)
		helper.HandleResult(ctx, data, err)
	})
}