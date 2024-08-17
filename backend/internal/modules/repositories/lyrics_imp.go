package repo

import (
	"context"
	"database/sql"

	models "the-lyricist/backend/internal/modules/models/boilerplate"

	db "the-lyricist/backend/internal/database"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LyricsRepositoryImp struct {
	db *sql.DB
}

func NewLyricsRepository() LyricsRepository {
	database, err := db.GetDB()
	if err != nil {
		return nil
	}
	return &LyricsRepositoryImp{db: database}
}

// get all lyrics based on translation and song id
func (r *LyricsRepositoryImp) GetLyrics(translationID, songID int) ([]*models.Lyric, error) {
	lyrics, err := models.Lyrics(qm.Where("translation_id = ? and song_id = ? ", translationID, songID)).All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	return lyrics, nil
}
