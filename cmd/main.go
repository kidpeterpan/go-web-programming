package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	_ = http.ListenAndServe(":8082", nil)
}

func handler(writer http.ResponseWriter, request *http.Request)  {
	_, _ = fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}