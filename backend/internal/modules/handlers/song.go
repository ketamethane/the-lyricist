package handlers

import (
	"net/http"
	"strconv"

	services "the-lyricist/backend/internal/modules/services"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	SongService services.SongService
}

func NewSongHandler(SongService services.SongService) *SongHandler {
	return &SongHandler{
		SongService: SongService,
	}
}

// improve database search for items
func (h *SongHandler) GetSongs(c *gin.Context) {
	songIDStr := c.Query("song-id")

	if songIDStr == "" {
		httpError(c, http.StatusBadRequest, "Missing required query parameters: company id")
		return
	}

	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		httpError(c, http.StatusBadRequest, "Invalid company id")
		return
	}

	lyrics, err := h.SongService.GetSongs(songID)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lyrics)
}
