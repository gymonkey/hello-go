package main

import "net/http"

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("hello"))
	})
	http.ListenAndServe(":9999", nil)
}
