package handler

import (
	"context"

	"location/internal/dataaccess/repo"
	"location/internal/model"
	"location/pkg/logger"

	"github.com/sirupsen/logrus"
)

type handler struct {
	repo repo.LocationRepo
	log  *logrus.Logger
}

func NewHandler(
	repo repo.LocationRepo,
) *handler {
	return &handler{
		repo: repo,
		log:  logger.GetLogger(),
	}
}

func (h *handler) ProcessLocationUpdate(ctx context.Context, loc *model.LocationUpdate) error {
	if err := h.repo.UpdateLocation(context.Background(), loc); err != nil {
		h.log.Errorln(err)
		return err
	}

	return nil
}
