package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CounterStore interface {
	IncrementCounters(ctx context.Context, targetURL string) (int64, int64, error)
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, password string, db int) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisStore{
		client: client,
	}, nil
}

func (rs *RedisStore) IncrementCounters(ctx context.Context, targetURL string) (int64, int64, error) {
	today := time.Now().Format("2006-01-02")

	todayKey := fmt.Sprintf("hits:%s:%s", targetURL, today)
	totalKey := fmt.Sprintf("hits:%s:total", targetURL)

	todayCount, err := rs.client.Incr(ctx, todayKey).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to increment today's counter: %w", err)
	}

	rs.client.Expire(ctx, todayKey, time.Hour*24)

	totalCount, err := rs.client.Incr(ctx, totalKey).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to increment total counter: %w", err)
	}

	return todayCount, totalCount, nil
}
