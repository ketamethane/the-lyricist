package repo

import (
	"context"
	"database/sql"

	models "the-lyricist/backend/internal/modules/models/boilerplate"

	db "the-lyricist/backend/internal/database"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SongRepositoryImp struct {
	db *sql.DB
}

func NewSongRepository() SongRepository {
	database, err := db.GetDB()
	if err != nil {
		return nil
	}
	return &SongRepositoryImp{db: database}
}

// get all lyrics based on translation and song id
func (r *SongRepositoryImp) GetSongs(songID int) ([]*models.Song, error) {
	songs, err := models.Songs(qm.Where("id = ? ", songID)).All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
