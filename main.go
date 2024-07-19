package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = 1 * time.Second
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr)
		conn, err := net.Dial("tcp", "192.168.0.24:8080")
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
		} else {
			conn.Write([]byte("hello world"))
			writer.WriteHeader(http.StatusOK)
		}
	})
	http.HandleFunc("/inner", func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://hello-java-qrid5mtbrq-uc.a.run.app", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
		} else {
			defer resp.Body.Close()
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(resp.Status))
			data, _ := io.ReadAll(resp.Body)
			writer.Write(data)
		}
	})
	http.HandleFunc("/outter", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Println(request.RemoteAddr)
		writer.WriteHeader(http.StatusOK)
		data, _ := io.ReadAll(request.Body)
		writer.Write([]byte(request.Header.Get("Authorization")))
		writer.Write([]byte(data))
	})
	http.ListenAndServe(":8080", nil)
	//for {
	//	fmt.Println("6" + strconv.FormatInt(time.Now().Unix(), 10))
	//	time.Sleep(10 * time.Second)
	//}
}
