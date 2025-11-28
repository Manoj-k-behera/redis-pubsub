package tests

import (
	"context"
	"encoding/json"
	"testing"

	r "redis-pubsub/internal/redis"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	args := m.Called(ctx, channel, message)
	return args.Get(0).(*redis.IntCmd)
}

func (m *MockRedis) Close() error {
	return nil
}

func TestPublisherPublishesMessage(t *testing.T) {
	mockClient := new(MockRedis)

	expectedPayload, _ := json.Marshal(map[string]string{"msg": "hello"})

	mockCmd := redis.NewIntCmd(context.Background())
	mockCmd.SetVal(1)
	mockCmd.SetErr(nil)

	mockClient.
		On("Publish", r.Ctx, "test-channel", expectedPayload).
		Return(mockCmd)

	publisher := r.NewPublisher(mockClient)

	err := publisher.Publish("test-channel", map[string]string{"msg": "hello"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mockClient.AssertExpectations(t)
}
