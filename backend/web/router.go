package web

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// KIV where to place logger

	// Initialize handlers
	songHandler := InitializeSongHandler()
	translationHandler := InitializeTranslationHandler()
	lyricHandler := InitializeLyricHandler()

	routerSongs := router.Group("/v1")
	{
		routerSongs.GET("/lyrics", songHandler.GetSongs)
		routerSongs.GET("/lyrics", translationHandler.GetTranslations)
		routerSongs.GET("/lyrics", lyricHandler.GetLyrics)
	}

	return router
}
