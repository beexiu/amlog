package main

import "net/http"

func main() {

	http.HandleFunc("/login", loginByName)
	http.HandleFunc("/record", recordAmInfo)

	http.ListenAndServe(":26031", nil)
}
