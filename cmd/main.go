package main

import (
	"fmt"
	"net/http"
)

func main() {
	StartServerMultiplexer()
}

func index(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func StartServerMultiplexer()  {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/",files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8082",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}