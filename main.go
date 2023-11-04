package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	emailPKG "github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
	"xorm.io/xorm"
)

var DB *xorm.Engine

type VerificationCodeIn struct {
	Email string `json:"email"`
}

type From struct {
	Address string
	Name    string
}

type EmailOptions struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

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

func MailSendCode(mail, code string) error {
	e := emailPKG.NewEmail()
	e.From = "kalougata@111.com"
	e.To = []string{mail}
	e.Subject = "邮箱验证码"
	e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	err := e.Send(
		"smtp.111.com:25",
		smtp.PlainAuth(
			"",
			"",
			"",
			"smtp.111.com",
		),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Post("/sendVerificationCode", func(ctx *fiber.Ctx) error {
		var body VerificationCodeIn
		if err := ctx.BodyParser(&body); err != nil {
			return ctx.SendString(err.Error())
		}

		if err := MailSendCode(body.Email, "123456"); err != nil {
			ctx.SendStatus(http.StatusInternalServerError)
			return ctx.SendString("发送邮箱验证码失败")
		}

		return ctx.SendString("ok!")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to listen Serve err: %s \n", err)
	}
}
