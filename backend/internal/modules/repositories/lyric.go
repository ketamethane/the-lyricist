package repo

import models "the-lyricist/backend/internal/modules/models/boilerplate"

type LyricRepository interface {
	GetLyrics(translationID, songID int) ([]*models.Lyric, error)
}
