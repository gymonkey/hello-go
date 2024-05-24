package main

import (
	"fmt"
	"time"
)

func main() {
	//http.DefaultClient.Timeout = 1 * time.Second
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request.RemoteAddr)
	//	writer.WriteHeader(http.StatusOK)
	//})
	//http.HandleFunc("/inner", func(writer http.ResponseWriter, request *http.Request) {
	//	//log.Println("recv req")
	//	writer.WriteHeader(http.StatusOK)
	//	writer.Write([]byte("inner"))
	//})
	//http.HandleFunc("/outter", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request.RemoteAddr)
	//	writer.WriteHeader(http.StatusOK)
	//	writer.Write([]byte("outter"))
	//})
	//http.ListenAndServe(":8080", nil)
	for {
		fmt.Println(time.Now())
		time.Sleep(10 * time.Second)
	}
}
