package main

import (
	"net/http"
	"./app"
)



func main() {


	http.ListenAndServe("127.0.0.1:8000", app.NewHttpHandler())
}