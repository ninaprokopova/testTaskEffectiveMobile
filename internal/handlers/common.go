package handlers

import (
	"log/slog"
	"testEffMobile/internal/service"

	"gorm.io/gorm"
)

type PersonHandler struct {
	DB       *gorm.DB
	enricher *service.EnricherService
	logger   *slog.Logger
}

func NewPersonHandler(db *gorm.DB, enricher *service.EnricherService, logger *slog.Logger) *PersonHandler {
	return &PersonHandler{
		DB:       db,
		enricher: enricher,
		logger:   logger,
	}
}
