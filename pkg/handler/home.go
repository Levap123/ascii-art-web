package handler

import (
	"net/http"

	"ascii-art-web/internal/logs"
	"ascii-art-web/pkg/handler/static"
)

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		logs.LoggerErr(r.URL.Path, r.Method, "Home Page", http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		logs.LoggerErr(r.URL.Path, r.Method, "Home Page", http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := static.ExecuteTemplate("home.html", nil, w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		logs.LoggerErr(r.URL.Path, r.Method, "Home Page", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	logs.LoggerOk(r.URL.Path, r.Method, "Home Page", http.StatusOK)
}
