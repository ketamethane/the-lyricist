package repo

import models "the-lyricist/backend/internal/modules/models/boilerplate"

type LyricsRepository interface {
	GetLyrics(translationID, songID int) ([]*models.Lyric, error)
}
