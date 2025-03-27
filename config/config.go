package config

type Config struct {
	RedisAddr     string `envconfig:"REDIS_ADDR" default:"localhost:6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisDB       int    `envconfig:"REDIS_DB" default:"0"`
	Port          string `envconfig:"PORT" default:"8080"`
	Debug         bool   `envconfig:"DEBUG" default:"false"`
}
