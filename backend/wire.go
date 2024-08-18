package thelyricist

// note that packages cannot contain hyphens

import (
	services "the-lyricist/backend/internal/modules/services"

	handlers "the-lyricist/backend/internal/modules/handlers"
	repo "the-lyricist/backend/internal/modules/repositories"

	"github.com/google/wire"
)

func InitializeSongHandler() *handlers.SongHandler {
	wire.Build(
		handlers.NewSongHandler,
		services.NewSongService,
		repo.NewSongRepository,
	)
	return nil
}

func InitializeTranslationHandler() *handlers.TranslationHandler {
	wire.Build(
		handlers.NewTranslationHandler,
		services.NewTranslationService,
		repo.NewTranslationRepository,
	)
	return nil
}

func InitializeLyricHandler() *handlers.LyricHandler {
	wire.Build(
		handlers.NewLyricHandler,
		services.NewLyricService,
		repo.NewLyricRepository,
	)
	return nil
}
