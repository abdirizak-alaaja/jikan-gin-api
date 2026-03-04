package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/schedules"
	"github.com/gin-gonic/gin"
)

// SchedulesRoutes registers /api/v4/schedules endpoints
func SchedulesRoutes(r *gin.Engine) {
	scheduleAPI := r.Group("/api/v4/schedules")

	// GET /api/v4/schedules?filter=<day>
	scheduleAPI.GET("", func(ctx *gin.Context) {
		filter := ctx.Query("filter") // can be monday, tuesday, etc.
		data, err := schedules.GetSchedules(schedules.ScheduleFilter(filter))
		helper.HandleResult(ctx, data, err)
	})
}