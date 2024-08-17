package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
)

type LyricService interface {
	GetLyrics(translationID, songID int) ([]*models.Lyric, error)
}
