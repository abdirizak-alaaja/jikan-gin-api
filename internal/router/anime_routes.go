package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/abdirizak-alaaja/jikan-gin-api/anime"
)

// AnimeRoutes registers all /api/v4/anime endpoints
func AnimeRoutes(r *gin.Engine) {

	animeAPI := r.Group("/api/v4/anime")

	// GET /api/v4/anime/:id
	animeAPI.GET("/:id", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}

		data, err := anime.GetAnimeById(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/characters
	animeAPI.GET("/:id/characters", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}

		data, err := anime.GetAnimeCharacters(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/staff
	animeAPI.GET("/:id/staff", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}

		data, err := anime.GetAnimeStaff(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/episodes?page=1
	animeAPI.GET("/:id/episodes", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}

		page := parsePage(ctx)
		data, err := anime.GetAnimeEpisodes(id, page)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/episodes/:episode
	animeAPI.GET("/:id/episodes/:episode", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}

		epParam := ctx.Param("episode")
		episode, err := strconv.Atoi(epParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid episode id"})
			return
		}

		data, err := anime.GetAnimeEpisodeById(id, episode)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/news?page=1
	animeAPI.GET("/:id/news", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		page := parsePage(ctx)
		data, err := anime.GetAnimeNews(id, page)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/forum?filter=all
	animeAPI.GET("/:id/forum", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		filter := ctx.Query("filter")
		data, err := anime.GetAnimeForum(id, anime.AnimeForumFilter(filter))
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/videos
	animeAPI.GET("/:id/videos", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeVideos(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/pictures
	animeAPI.GET("/:id/pictures", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimePictures(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/statistics
	animeAPI.GET("/:id/statistics", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeStatistics(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/moreinfo
	animeAPI.GET("/:id/moreinfo", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeMoreInfo(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/recommendations
	animeAPI.GET("/:id/recommendations", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeRecommendations(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/userupdates?page=1
	animeAPI.GET("/:id/userupdates", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		page := parsePage(ctx)
		data, err := anime.GetAnimeUserUpdates(id, page)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/reviews?page=1
	animeAPI.GET("/:id/reviews", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		page := parsePage(ctx)
		data, err := anime.GetAnimeReviews(id, page)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/relations
	animeAPI.GET("/:id/relations", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeRelations(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/themes
	animeAPI.GET("/:id/themes", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeThemes(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/:id/external
	animeAPI.GET("/:id/external", func(ctx *gin.Context) {
		id, ok := parseID(ctx)
		if !ok {
			return
		}
		data, err := anime.GetAnimeExternal(id)
		handleResult(ctx, data, err)
	})

	// GET /api/v4/anime/search?query=...
	animeAPI.GET("/search", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		data, err := anime.GetAnimeSearch(query)
		handleResult(ctx, data, err)
	})
}

// helper to parse id param
func parseID(ctx *gin.Context) (int, bool) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

// helper to parse page query
func parsePage(ctx *gin.Context) int {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 1
	}
	return page
}

// helper to handle results
func handleResult(ctx *gin.Context, data any, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}