package data

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	db := NewMysql(conf)
	cache := NewRedis(conf)
	data := &Data{
		DB:    db,
		Cache: cache,
	}

	db.ShowSQL(true)

	return data, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database: %s \n", err)
		}
		if err := cache.Close(); err != nil {
			log.Fatalf("Failed to close redis client: %s \n", err)
		}
	}, nil
}

func NewMysql(conf *config.Config) *xorm.Engine {
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

	return db
}

func NewRedis(conf *config.Config) *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RDB.Host, conf.RDB.Port),
		Username: conf.RDB.User,
		Password: conf.RDB.Passwd,
		DB:       conf.RDB.Db,
	})

	if err := db.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect redis client %s", err)
	}

	return db
}
