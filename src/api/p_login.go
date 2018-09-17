package main

import "net/http"

func loginByName(w http.ResponseWriter, r *http.Request) {
	err := SS.Login(w, r)

	if err == nil {
		w.Write([]byte("login success."))
	} else {
		w.Write([]byte("login failed."))
	}
}
