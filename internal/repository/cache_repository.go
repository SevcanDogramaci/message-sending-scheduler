package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/SevcanDogramaci/message-sending-scheduler/pkg/redis"
)

type CacheRepository struct {
	redis *redis.Redis
}

func NewCacheRepository(redis *redis.Redis) *CacheRepository {
	return &CacheRepository{redis: redis}
}

func (r *CacheRepository) SetMessage(transferMetadata *model.TransferMetadata) error {
	ttl := time.Duration(r.redis.Config.DefaultTTLSecs) * time.Second

	transferMetadataBytes, err := json.Marshal(transferMetadata)
	if err != nil {
		return err
	}

	status := r.redis.Client.Set(
		context.Background(),
		transferMetadata.ID,
		transferMetadataBytes,
		ttl,
	)

	return status.Err()
}
