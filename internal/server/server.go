package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	ServerHTTP *fiber.App
}

func NewServer(
	ServerHTTP *fiber.App,
) *Server {
	return &Server{ServerHTTP}
}
