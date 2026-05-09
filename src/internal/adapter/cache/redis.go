package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"

	"sai-server/internal/port"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(addr string) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Redis{client: client}, nil
}

func NewRedisOrNoop(addr string) port.CacheClient {
	r, err := NewRedis(addr)
	if err != nil {
		return &noopCache{}
	}
	return r
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	return n > 0, err
}

func (r *Redis) Publish(ctx context.Context, channel string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, channel, data).Err()
}

func (r *Redis) Subscribe(ctx context.Context, channel string) (<-chan string, error) {
	pubsub := r.client.Subscribe(ctx, channel)
	out := make(chan string, 64)

	go func() {
		defer pubsub.Close()
		defer close(out)

		ch := pubsub.Channel()
		for {
			select {
			case msg := <-ch:
				if msg == nil {
					return
				}
				out <- msg.Payload
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}

type noopCache struct{}

func (n *noopCache) Get(_ context.Context, _ string) (string, error) {
	return "", nil
}

func (n *noopCache) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) error {
	return nil
}

func (n *noopCache) Del(_ context.Context, _ string) error {
	return nil
}

func (n *noopCache) Exists(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (n *noopCache) Publish(_ context.Context, _ string, _ interface{}) error {
	return nil
}

func (n *noopCache) Subscribe(_ context.Context, _ string) (<-chan string, error) {
	ch := make(chan string)
	close(ch)
	return ch, nil
}
