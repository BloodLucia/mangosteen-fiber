package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	emailPKG "github.com/jordan-wright/email"
	"github.com/kalougata/bookkeeping/cmd/wire"
	"github.com/kalougata/bookkeeping/pkg/http"
	"log"
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
	server, cleanup, err := wire.NewApp()

	if err != nil {
		log.Panicln(err)
	}

	http.Run(server.ServerHTTP, ":8888")

	defer cleanup()
}
