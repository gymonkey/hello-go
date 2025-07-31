package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//config := volcengine.NewConfig().WithRegion(region).WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	//s, err := session.NewSession(config)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//c := cr.New(s)
	//input := &cr.GetUserInput{
	//	Registry: volcengine.String("acc-test-cr"),
	//}

	//vefaas.ListRevisionsOutput{}

	// 添加HTTP服务器监听8080端口
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World1111!"))
	})

	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
