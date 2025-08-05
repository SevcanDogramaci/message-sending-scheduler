package repository

import (
	"context"
	"time"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	redisClient *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{redisClient: client}
}

func (r *CacheRepository) SetMessage(transferMetadata model.TransferMetadata) error {
	status := r.redisClient.Set(
		context.Background(),
		transferMetadata.ID,
		transferMetadata,
		time.Duration(1),
	)

	return status.Err()
}
