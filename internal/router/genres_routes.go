package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/genres"
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/gin-gonic/gin"
)

// GenresRoutes registers /api/v4/genres endpoints
func GenresRoutes(r *gin.Engine) {
	genresAPI := r.Group("/api/v4/genres")

	// GET /api/v4/genres/anime?page=1&limit=25&filter=genres
	genresAPI.GET("/anime", func(ctx *gin.Context) {
		page := helper.ParsePage(ctx)
		limit := helper.ParseLimit(ctx)
		filter := ctx.DefaultQuery("filter", string(genres.GenresFilterGenres))
		data, err := genres.GetAnimeGenres(page, limit, genres.GenreFilter(filter))
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/genres/manga?page=1&limit=25&filter=genres
	genresAPI.GET("/manga", func(ctx *gin.Context) {
		page := helper.ParsePage(ctx)
		limit := helper.ParseLimit(ctx)
		filter := ctx.DefaultQuery("filter", string(genres.GenresFilterGenres))
		data, err := genres.GetMangaGenres(page, limit, genres.GenreFilter(filter))
		helper.HandleResult(ctx, data, err)
	})
}