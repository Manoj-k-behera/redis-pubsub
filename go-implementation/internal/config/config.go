package config

import "os"

type RedisConfig struct {
	Host string
	Port string
}

func Load() RedisConfig{
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "6379"
	}

	return RedisConfig{
		Host: host,
		Port: port,
	}
}