package repo

import models "the-lyricist/backend/internal/modules/models/boilerplate"

type TranslationRepository interface {
	GetTranslations(translationID int) ([]*models.Translation, error)
}
