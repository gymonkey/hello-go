package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/inner", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("inner"))
	})
	http.HandleFunc("/outter", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("outter"))
	})
	go http.ListenAndServe(":8080", nil)
	for {
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(10 * time.Second)
	}
}
