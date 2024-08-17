package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
	repo "the-lyricist/backend/internal/modules/repositories"
)

type LyricServiceImp struct {
	Repo repo.LyricsRepository
}

func NewLyricService(repo repo.LyricsRepository) LyricService {
	return &LyricServiceImp{Repo: repo}
}

func (s *LyricServiceImp) GetLyrics(translationID, songID int) ([]*models.Lyric, error) {
	// Get activities by emission IDs
	lyrics, err := s.Repo.GetLyrics(translationID, songID)
	if err != nil {
		return nil, err
	}

	return lyrics, nil
}
