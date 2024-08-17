package thelyricist

// note that packages cannot contain hyphens

import (
	services "the-lyricist/backend/internal/modules/services"

	handlers "the-lyricist/backend/internal/modules/handlers"
	repo "the-lyricist/backend/internal/modules/repositories"

	"github.com/google/wire"
)

func InitializeLyricHandler() *handlers.LyricHandler {
	wire.Build(
		handlers.NewLyricHandler,
		services.NewLyricService,
		repo.NewLyricsRepository,
	)
	return nil
}
