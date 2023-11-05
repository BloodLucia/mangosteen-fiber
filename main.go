package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	emailPKG "github.com/jordan-wright/email"
	"github.com/kalougata/bookkeeping/cmd/wire"
	"github.com/kalougata/bookkeeping/pkg/http"
	"log"
	"net/smtp"
)

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
