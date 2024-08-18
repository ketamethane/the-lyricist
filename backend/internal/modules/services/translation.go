package services

import (
	models "the-lyricist/backend/internal/modules/models/boilerplate"
)

type TranslationService interface {
	GetTranslations(translationID int) ([]*models.Translation, error)
}
