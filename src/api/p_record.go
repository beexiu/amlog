package main

import (
	"fmt"
	"net/http"
)

func recordAmInfo(w http.ResponseWriter, r *http.Request) {
	if !SS.Validate(r) {
		w.Write([]byte("Invalid user."))
		return
	}

	miles := r.FormValue("miles")
	cash := r.FormValue("cash")

	msg := fmt.Sprintf("Record: %s %s\n", miles, cash)
	w.Write([]byte(msg))
}
