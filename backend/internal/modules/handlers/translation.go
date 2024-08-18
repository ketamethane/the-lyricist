package handlers

import (
	"net/http"
	"strconv"

	services "the-lyricist/backend/internal/modules/services"

	"github.com/gin-gonic/gin"
)

type TranslationHandler struct {
	TranslationService services.TranslationService
}

func NewTranslationHandler(TranslationService services.TranslationService) *TranslationHandler {
	return &TranslationHandler{
		TranslationService: TranslationService,
	}
}

func (h *TranslationHandler) GetTranslations(c *gin.Context) {
	songIDStr := c.Query("song-id")
	translationIDStr := c.Query("translation-id")

	if songIDStr == "" {
		httpError(c, http.StatusBadRequest, "Missing required query parameters: company id")
		return
	}

	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		httpError(c, http.StatusBadRequest, "Invalid company id")
		return
	}

	if translationIDStr == "" {
		httpError(c, http.StatusBadRequest, "Missing required query parameters: translation id")
		return
	}

	translationID, err := strconv.Atoi(translationIDStr)
	if err != nil {
		httpError(c, http.StatusBadRequest, "Invalid company id")
		return
	}

	lyrics, err := h.TranslationService.GetTranslations(songID, translationID)
	if err != nil {
		httpError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lyrics)
}
