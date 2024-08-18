package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
)

type SongService interface {
	GetSongs(songID int) ([]*models.Song, error)
}
