package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func (rc *RedisClient) Init(ctx context.Context, address, password string) error {
	r := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		return err
	}

	rc.client = r
	return nil
}

func (rc *RedisClient) SaveMessage(ctx context.Context, roomID string, message *Message) error {
	text, err := json.Marshal(message)
	if err != nil {
		return err
	}

	member := &redis.Z{
		Score:  float64(message.Timestamp),
		Member: text,
	}

	_, err = rc.client.ZAdd(ctx, roomID, *member).Result()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) GetMessages(ctx context.Context, roomID string, start, end int64, reverse bool) ([]*Message, error) {
	var (
		rawMessages []string
		messages    []*Message
		err         error
    )

    if reverse {
		rawMessages, err = rc.client.ZRevRange(ctx, roomID, start, end).Result()
    } else {
		rawMessages, err = rc.client.ZRange(ctx, roomID, start, end).Result()
    }

	if err != nil {
		return nil, err
	}
	

    for _, msg := range rawMessages {
       temp := &Message{}
       err := json.Unmarshal([]byte(msg), temp)
       if err != nil {
          return nil, err
       }
       messages = append(messages, temp)
    }

    return messages, nil
}