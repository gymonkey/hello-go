package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// 添加全局变量来存储进程启动时间
var startTime time.Time

func main() {
	// 记录进程启动时间
	startTime = time.Now()

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
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s\n", r.URL.Path)
		if time.Now().After(startTime.Add(5 * time.Minute)) {
			time.Sleep(10 * time.Second)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World1111!"))
	})

	// 添加新的处理函数来返回进程启动时间
	http.HandleFunc("/start-time", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Process started at: %s", startTime.Format(time.RFC3339))))
	})

	// 添加处理函数来返回所有环境变量
	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		envs := os.Environ()
		envString := strings.Join(envs, "\n")
		w.Write([]byte(fmt.Sprintf("Environment Variables:\n%s", envString)))
	})

	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
