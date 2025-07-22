package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for key, val := range request.Header {
			fmt.Printf("key: %s values: %v\n", key, val)
		}
		writer.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8081", nil)
}
