package repo

import (
	"context"

	"location/internal/model"
)

type LocationRepo interface {
	UpdateLocation(ctx context.Context, location *model.LocationUpdate) error
	GetLatestLocation(ctx context.Context, deviceId string, filters model.LocationFilter) ([]*model.LocationUpdateValue, error)
}
