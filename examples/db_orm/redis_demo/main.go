package main

import "github.com/go-redis/redis"

type Config2 struct {
	Redis 	RedisOpt
}

type RedisOpt struct {
	DB				int
	Password		string
	PoolSize 		int
	Cluster 		[]string
}

var redisClient *redis.ClusterClient

func (r *RedisOpt) initRedisCluster() error {
	//
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: r.Cluster,
		PoolSize: r.PoolSize,
	})

	err := client.Ping().Err()
	if err != nil {
		return err
	}
	redisClient = client

	return nil
}

func main() {
	//filePath
}
