package main

import "net/http"

func loginByName(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			audit.Error("", "run login error, err=%v", err)
		}
	}()

    err := SS.Login(w, r)

	if err == nil {
		w.Write([]byte("login success."))
	} else {
		w.Write([]byte("login failed."))
	}
}
