package main

import (
	"fmt"
	"net/http"
)

func recordAmInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			audit.Error("", "run record error, err=%v", err)
		}
	}()

	if !SS.Validate(r) {
		w.Write([]byte("Invalid user."))
		return
	}

	miles := r.FormValue("miles")
	cash := r.FormValue("cash")

	msg := fmt.Sprintf("Record: %s %s\n", miles, cash)
	w.Write([]byte(msg))
}
