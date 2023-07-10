package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type LinkStore struct {
	rdb *redis.Client
	ctx context.Context
}

func NewStore() *LinkStore {
	redis_port := os.Getenv("REDIS_PORT")
	redis_host := os.Getenv("REDIS_HOST")
	if redis_port == "" {
		redis_port = "6379"
	}
	if redis_host == "" {
		redis_host = "localhost"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	return &LinkStore{rdb: rdb, ctx: ctx}
}

func (ls *LinkStore) checkLink(url string) (*PreviewInfo, error) {
	val, err := ls.rdb.Get(ls.ctx, url).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var d PreviewInfo
	err = json.Unmarshal([]byte(val), &d)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &d, nil

}

func (ls *LinkStore) saveLink(url string, preview PreviewInfo) error {

	p, err := json.Marshal(preview)
	if err != nil {
		fmt.Println("PARCE ERROR", err.Error())
		return err
	}
	err = ls.rdb.Set(ls.ctx, url, p, 0).Err()
	if err != nil {
		fmt.Println("SAVE ERROR", err.Error())
		return err
	}
	ls.rdb.Expire(ls.ctx, url, time.Duration(5*time.Hour))

	return nil
}
