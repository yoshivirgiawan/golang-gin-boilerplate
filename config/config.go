package config

import (
	"boilerplate/helper"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

type DBConfig struct {
	Type     string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func InitDB(dbConfig DBConfig) {
	var (
		dialect gorm.Dialector
		dsn     string
		err     error
	)

	switch dbConfig.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		dialect = mysql.Open(dsn)
	case "pgsql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Database,
		)
		dialect = postgres.Open(dsn)
	default:
		helper.LogError(fmt.Errorf("Unknown database type: %s", dbConfig.Type))
	}

	DB, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		helper.LogError(err)
	}
}

func InitRedis(redisConfig RedisConfig) {
	ctx := context.Background() // Konteks untuk operasi-operasi Redis

	if redisConfig.Password == "null" {
		redisConfig.Password = ""
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		helper.LogError(fmt.Errorf("Error connecting to Redis: \n%s", err))
	}
}
