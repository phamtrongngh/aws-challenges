package rediscache

import (
	"context"
	"encoding/json"
	"location/internal/dataaccess/repo"
	"location/internal/model"
	"location/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type locationRedisCache struct {
	client  *redis.Client
	log     *logrus.Logger
	locRepo repo.LocationRepo
}

func NewLocationRedisCache(
	addr string,
	password string,
	db int,
	locRepo repo.LocationRepo,
) *locationRedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &locationRedisCache{
		client:  client,
		log:     logger.GetLogger(),
		locRepo: locRepo,
	}
}

func (r *locationRedisCache) UpdateLocation(ctx context.Context, location *model.LocationUpdate) error {
	return r.locRepo.UpdateLocation(ctx, location)
}

func (r *locationRedisCache) GetLatestLocation(ctx context.Context, deviceId string, filters model.LocationFilter) ([]*model.LocationUpdateValue, error) {
	cachedData, err := r.client.Get(ctx, deviceId).Result()
	if err == redis.Nil {
		// Cache miss, fetch data from the database
		locUpdateValues, err := r.locRepo.GetLatestLocation(ctx, deviceId, filters)
		if err != nil {
			r.log.Errorf("[Redis Cache] Error fetching location data from database: %v", err)
			return nil, err
		}
		// Store the fetched data in the cache
		data, err := json.Marshal(locUpdateValues)
		if err != nil {
			r.log.Errorf("[Redis Cache] Error marshaling location data: %v", err)
			return nil, err
		}
		err = r.client.Set(ctx, deviceId, data, 24*time.Hour).Err()
		if err != nil {
			r.log.Errorf("[Redis Cache] Error storing location data in cache: %v", err)
			return nil, err
		}
		return locUpdateValues, nil
	} else if err != nil {
		r.log.Errorf("[Redis Cache] Error fetching location data from cache: %v", err)
		return nil, err
	}
	// Cache hit, return the data from the cache
	var locUpdateValues []*model.LocationUpdateValue
	err = json.Unmarshal([]byte(cachedData), &locUpdateValues)
	if err != nil {
		r.log.Errorf("[Redis Cache] Error unmarshaling location data: %v", err)
		return nil, err
	}
	return locUpdateValues, nil
}
