package handler

import (
	"net/http"

	"ascii-art-web/pkg/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	css := http.FileServer(http.Dir("static"))
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", css))
	router.HandleFunc("/", h.HomePage)
	router.HandleFunc("/ascii-art", h.AsciiArt)
	return router
}
