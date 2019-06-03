package server

import "github.com/go-redis/redis"

var redisClient *redis.ClusterClient

type Redis struct {
	DB				int
	Password		string
	PoolSize 		int
	Cluster 		[]string
}

func (r *Redis) initRedisCluster() error {
	//
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: r.Cluster,
	})

	err := client.Ping().Err()
	if err != nil {
		return err
	}
	redisClient = client

	return nil
}

func RC() *redis.ClusterClient {
	return redisClient
}
