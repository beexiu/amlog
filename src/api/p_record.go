package main

import (
	"fmt"
	"net/http"
)

func recordAmInfo(w http.ResponseWriter, r *http.Request) {
	miles := r.FormValue("miles")
	cash := r.FormValue("cash")

	fmt.Printf("Record: %s %s\n", miles, cash)
}
