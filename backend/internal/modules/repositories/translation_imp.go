package repo

import (
	"context"
	"database/sql"

	models "the-lyricist/backend/internal/modules/models/boilerplate"

	db "the-lyricist/backend/internal/database"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TranslationRepositoryImp struct {
	db *sql.DB
}

func NewTranslationRepository() TranslationRepository {
	database, err := db.GetDB()
	if err != nil {
		return nil
	}
	return &TranslationRepositoryImp{db: database}
}

// get all lyrics based on translation and song id
func (r *TranslationRepositoryImp) GetTranslations(translationID int) ([]*models.Translation, error) {
	translations, err := models.Translations(qm.Where("id = ? ", translationID)).All(context.Background(), r.db)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
