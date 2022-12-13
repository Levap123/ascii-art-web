package main

import (
	"ascii-art-web/internal/logs"
	"ascii-art-web/internal/server"
	"ascii-art-web/pkg/handler"
	"ascii-art-web/pkg/service"
)

const port = "80"

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)
	routes := handler.InitRoutes()
	server := new(server.Server)
	logs.LoggerOut.Printf("server is listening on: http://localhost:%s\n", port)
	if err := server.Run(port, routes); err != nil {
		logs.ErrLoggerOut.Fatal(err)
	}
}
