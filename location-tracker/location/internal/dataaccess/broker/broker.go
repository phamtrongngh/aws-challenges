package broker

import (
	"context"

	"location/internal/model"
)

type LocationProducer interface {
	Produce(ctx context.Context, msg *model.LocationUpdate) error
}

type LocationConsumer interface {
	Consume(ctx context.Context, handler func(ctx context.Context, msg *model.LocationUpdate) error) error
}
