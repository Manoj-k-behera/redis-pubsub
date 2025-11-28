package redis

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Subscriber struct {
	Client *redis.Client
}

func NewSubscriber(client *redis.Client) *Subscriber {
	return &Subscriber{Client: client}
}

func (s *Subscriber) Subscribe(channel string, handler func(map[string]interface{})) error {
	sub := s.Client.Subscribe(Ctx, channel)

	ch := sub.Channel()

	go func() {
		for msg := range ch {
			var data map[string]interface{}
			json.Unmarshal([]byte(msg.Payload), &data)
			handler(data)
		}
	}()

	return nil
}
