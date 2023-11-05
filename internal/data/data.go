package data

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kalougata/bookkeeping/pkg/config"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type Data struct {
	DB    *xorm.Engine
	Cache *redis.Client
}

func NewData(conf *config.Config) (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.DB.User,
		conf.DB.Passwd,
		conf.DB.Host,
		conf.DB.Post,
		conf.DB.DbName,
	)

	db, err := xorm.NewEngine(conf.DB.Driver, dsn)

	if err != nil {
		log.Fatalf("Failed to use NewEngine with xorm: %s \n", err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("Failed to connect database: %s \n", err)
	}

	cache, err := newRedis()
	if err != nil {
		log.Fatalf("Failed to connect redis client %s \n", err)
	}

	data := &Data{
		DB:    db,
		Cache: cache,
	}

	return data, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database: %s \n", err)
		}
		if err := cache.Close(); err != nil {
			log.Fatalf("Failed to close redis client: %s \n", err)
		}
	}, nil
}

func newRedis() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "root",
		DB:       0,
	})

	if err := db.Ping(context.TODO()).Err(); err != nil {
		log.Fatalf("Failed to connect redis client %s", err)
	}

	if err := db.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return db, nil
}
