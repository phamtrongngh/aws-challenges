package mongorepo

import (
	"context"
	"location/pkg/logger"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"location/internal/model"
)

type locationMongoRepo struct {
	client *mongo.Client
	log    *logrus.Logger
}

func NewMongoLocationRepo(uri string) *locationMongoRepo {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &locationMongoRepo{
		client: client,
		log:    logger.GetLogger(),
	}
}

func (r *locationMongoRepo) UpdateLocation(ctx context.Context, locUpdate *model.LocationUpdate) error {
	deviceLocModels := make([]*Location, len(locUpdate.Values))

	for i := range locUpdate.Values {
		lat, _ := strconv.ParseFloat(locUpdate.Values[i].Latitude, 64)
		long, _ := strconv.ParseFloat(locUpdate.Values[i].Longitude, 64)
		timestamp := time.UnixMilli(locUpdate.Values[i].Timestamp)

		deviceLocModels[i] = &Location{
			Metadata: Metadata{
				CarSku: locUpdate.Info.CarSku,
				Entity: Entity{
					Id:   locUpdate.Info.Entity.Id,
					Type: locUpdate.Info.Entity.Type,
				},
			},
			Timestamp: timestamp,
			Coords: Coordinates{
				Type:        "Point",
				Coordinates: []float64{lat, long},
			},
			Speed: locUpdate.Values[i].SpeedOverGround,
		}
	}

	deviceLocInterfaces := make([]any, len(deviceLocModels))
	for i, v := range deviceLocModels {
		deviceLocInterfaces[i] = v
	}

	collection := r.client.Database("locations").Collection("locations")
	_, err := collection.InsertMany(ctx, deviceLocInterfaces)
	if err != nil {
		r.log.Errorf("[MongoDB] Error inserting location data: %v", err)
		return err
	}

	return nil
}

func (r *locationMongoRepo) GetLatestLocation(ctx context.Context, deviceId string, filters model.LocationFilter) ([]*model.LocationUpdateValue, error) {
	filter := map[string]any{
		"metadata.entity.id":   deviceId,
		"metadata.entity.type": Device,
	}

	opts := options.Find().SetSort(map[string]int{"timestamp": -1})

	if filters.Limit != nil {
		opts.SetLimit(*filters.Limit)
	}

	collection := r.client.Database("locations").Collection("locations")
	result, err := collection.Find(ctx, filter, opts)
	if err != nil {
		r.log.Errorf("[MongoDB] Error finding location data: %v", err)
		return nil, err
	}
	defer result.Close(ctx)

	var locs []*Location
	if err = result.All(ctx, &locs); err != nil {
		r.log.Errorf("[MongoDB] Error decoding location data: %v", err)
		return nil, err
	}

	locUpdateValues := []*model.LocationUpdateValue{}
	for _, loc := range locs {
		locUpdateValues = append(locUpdateValues, &model.LocationUpdateValue{
			Timestamp:       loc.Timestamp.UnixMilli(),
			Longitude:       strconv.FormatFloat(loc.Coords.Coordinates[1], 'f', -1, 64),
			Latitude:        strconv.FormatFloat(loc.Coords.Coordinates[0], 'f', -1, 64),
			SpeedOverGround: loc.Speed,
		})
	}

	return locUpdateValues, nil
}
