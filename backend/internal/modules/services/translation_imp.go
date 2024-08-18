package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
	repo "the-lyricist/backend/internal/modules/repositories"
)

type TranslationServiceImp struct {
	Repo repo.TranslationRepository
}

func NewTranslationService(repo repo.TranslationRepository) TranslationService {
	return &TranslationServiceImp{Repo: repo}
}

func (s *TranslationServiceImp) GetTranslations(translationID int) ([]*models.Translation, error) {
	// Get activities by emission IDs
	translations, err := s.Repo.GetTranslations(translationID)
	if err != nil {
		return nil, err
	}

	return translations, nil
}
