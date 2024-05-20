package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = 1 * time.Second
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr)
		writer.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/inner", func(writer http.ResponseWriter, request *http.Request) {
		resp, err := http.Get("http://10.116.8.153:8000/abc")
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(string(data))
			writer.WriteHeader(resp.StatusCode)
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			writer.Write(data)
		}
	})
	http.HandleFunc("/outter", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("outter"))
	})
	go http.ListenAndServe(":8080", nil)
	for {
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(10 * time.Second)
	}
}
