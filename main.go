package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func init() {
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

	DB = db
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database: %s \n", err)
		}
	}()
}

func main() {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to listen Serve err: %s \n", err)
	}
}
