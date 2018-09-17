package main

import (
	"net/http"
	"path"
	"strings"
)

func getCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func setCookie(w http.ResponseWriter, name, value string, age int) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: age,
		Secure: false,
	})
}

func setContentType(w http.ResponseWriter, filePath string) {
	switch strings.ToLower(path.Ext(filePath)) {
	case ".css":
		w.Header().Add("Content-Type", "text/css; charset=utf-8")
	case ".js":
		w.Header().Add("Content-Type", "application/javascript")
	case ".jpg", ".jpeg":
		w.Header().Add("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Add("Content-Type", "image/png")
	case ".txt", ".log":
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	default:
		w.Header().Add("Content-Type", "application/octet-stream")
	}
}
