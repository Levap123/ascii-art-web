package handler

import (
	"net/http"

	"ascii-art-web/internal/logs"
	"ascii-art-web/pkg/handler/static"
)

func (h *Handler) AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		logs.LoggerErr(r.URL.Path, r.Method, "ASCII art", http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	target := r.FormValue("ascii")
	banner := r.FormValue("banner")
	ascii, err := h.Service.Ascii.GenerateWords(target, banner)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		logs.LoggerErr(r.URL.Path, r.Method, "ASCII art", http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := static.ExecuteTemplate("home.html", ascii, w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		logs.LoggerErr(r.URL.Path, r.Method, "ASCII art", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	logs.LoggerOk(r.URL.Path, r.Method, "ASCII art", http.StatusOK)
}
