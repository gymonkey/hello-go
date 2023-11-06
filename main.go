package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("hello"))
	})
	fmt.Println("listen port: " + os.Getenv("PORT"))
	http.ListenAndServe(strings.Join([]string{"", os.Getenv("PORT")}, ":"), nil)
}
