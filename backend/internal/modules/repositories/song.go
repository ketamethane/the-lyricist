package repo

import models "the-lyricist/backend/internal/modules/models/boilerplate"

type SongRepository interface {
	GetSongs(songID int) ([]*models.Song, error)
}
