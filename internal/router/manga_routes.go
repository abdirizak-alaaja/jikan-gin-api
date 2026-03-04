package router

import (
	"github.com/abdirizak-alaaja/jikan-gin-api/helper"
	"github.com/abdirizak-alaaja/jikan-gin-api/manga"
	"github.com/gin-gonic/gin"
)

// MangaRoutes registers /api/v4/manga endpoints
func MangaRoutes(r *gin.Engine) {
	mangaAPI := r.Group("/api/v4/manga")

	// GET /api/v4/manga/:id
	mangaAPI.GET("/:id", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaById(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/characters
	mangaAPI.GET("/:id/characters", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaCharacters(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/news?page=...
	mangaAPI.GET("/:id/news", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		page := helper.ParsePage(ctx)
		data, err := manga.GetMangaNews(id, page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/forum?filter=...
	mangaAPI.GET("/:id/forum", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		filter := ctx.DefaultQuery("filter", string(manga.MangaForumFilterAll))
		data, err := manga.GetMangaForum(id, manga.MangaForumFilter(filter))
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/pictures
	mangaAPI.GET("/:id/pictures", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaPictures(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/statistics
	mangaAPI.GET("/:id/statistics", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaStatistics(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/moreinfo
	mangaAPI.GET("/:id/moreinfo", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaMoreInfo(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/recommendations
	mangaAPI.GET("/:id/recommendations", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaRecommendations(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/userupdates?page=...
	mangaAPI.GET("/:id/userupdates", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		page := helper.ParsePage(ctx)
		data, err := manga.GetMangaUserUpdates(id, page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/reviews?page=...
	mangaAPI.GET("/:id/reviews", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		page := helper.ParsePage(ctx)
		data, err := manga.GetMangaReviews(id, page)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/relations
	mangaAPI.GET("/:id/relations", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaRelations(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/:id/external
	mangaAPI.GET("/:id/external", func(ctx *gin.Context) {
		id, ok := helper.ParseID(ctx)
		if !ok {
			return
		}
		data, err := manga.GetMangaExternal(id)
		helper.HandleResult(ctx, data, err)
	})

	// GET /api/v4/manga/search?query=...
	mangaAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := manga.GetMangaSearch(query)
		helper.HandleResult(ctx, data, err)
	})
}