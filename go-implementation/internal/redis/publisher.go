package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type Publisher struct {
	Client RedisAdapter
}

type RedisAdapter interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Close() error
}

func NewPublisher(client RedisAdapter) *Publisher {
	return &Publisher{Client: client}	
}

func (p *Publisher) Publish(channel string, message interface{}) error {
	jsonData, err := json.Marshal(message)

	if err != nil {
		return err
	}
	return p.Client.Publish(Ctx, channel, jsonData).Err()
}

func (p *Publisher) Close() error {
	return p.Client.Close()
}