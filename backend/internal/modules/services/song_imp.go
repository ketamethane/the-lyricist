package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
	repo "the-lyricist/backend/internal/modules/repositories"
)

type SongServiceImp struct {
	Repo repo.SongRepository
}

func NewSongService(repo repo.SongRepository) SongService {
	return &SongServiceImp{Repo: repo}
}

func (s *SongServiceImp) GetSongs(songID int) ([]*models.Song, error) {
	// Get activities by emission IDs
	songs, err := s.Repo.GetSongs(songID)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
