package main

import (
	"redis-pubsub/internal/config"
	r "redis-pubsub/internal/redis"
	"time"
)

func main() {
	cfg := config.Load()
	client := r.NewRedisClient(cfg)
	channel := "example-channel"

	publisher := r.NewPublisher(client)
	subscriber := r.NewSubscriber(client)
	defer client.Close()

	error := subscriber.Subscribe(channel, func(message map[string]interface{}) {
		println("Received message:", message["text"].(string))
	})

	if error != nil {
		println("Subscription error:", error.Error())
		return
	}

	err := publisher.Publish(channel, map[string]interface{}{
		"hello": "world",
		"text":  "Hello from Go Publisher!",
	})

	if err != nil {
		println("Publish error:", err.Error())
		return
	}

	time.Sleep(2 * time.Second)
}
