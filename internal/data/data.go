package data

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData() (*Data, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		"root",
		"123456",
		"127.0.0.1",
		3306,
		"bookkeeping_dev",
	)

	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		log.Fatalf("Failed to use NewEngine with xorm: %s \n", err)
	}

	if err := db.PingContext(context.TODO()); err != nil {
		log.Fatalf("Failed to connect database: %s \n", err)
	}

	data := &Data{
		DB: db,
	}

	return data, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database: %s \n", err)
		}
	}, nil
}
