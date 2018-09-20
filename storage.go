package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type dbClient struct {
	redisClient *redis.Client
}

func getDBClient() *dbClient {
	addr := getRedisAddr()
	logger.Println("connecting to resdis at", addr)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &dbClient{redisClient: client}
}

func getRedisAddr() string {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func (c *dbClient) isReady() bool {
	pong, err := c.redisClient.Ping().Result()
	if err != nil {
		return false
	}
	return pong == "PONG"
}

func (c *dbClient) createUser(user User) error {
	sadd := c.redisClient.SAdd("users", user.UUID)
	if err := sadd.Err(); err != nil {
		logger.Println(err)
		return err
	}

	umap, err := userToMap(user)
	if err != nil {
		return err
	}

	hmset := c.redisClient.HMSet(fmt.Sprintf("users:%s", user.UUID), umap)
	if err := hmset.Err(); err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func (c *dbClient) updateUser(user User) error {
	umap, err := userToMap(user)
	if err != nil {
		logger.Println(err)
		return err
	}

	hmset := c.redisClient.HMSet(fmt.Sprintf("users:%s", user.UUID), umap)
	if err := hmset.Err(); err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func (c *dbClient) deleteUser(uuid string) error {
	srem := c.redisClient.SRem("users", uuid)
	if err := srem.Err(); err != nil {
		logger.Println(err)
		return err
	}

	del := c.redisClient.Del(fmt.Sprintf("users:%s", uuid))
	if err := del.Err(); err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func userToMap(user User) (map[string]interface{}, error) {
	umap := make(map[string]interface{})
	bytes, err := json.Marshal(user)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	if err := json.Unmarshal(bytes, &umap); err != nil {
		logger.Println(err)
		return nil, err
	}

	return umap, nil
}
