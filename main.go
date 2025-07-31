package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

	// 创建Gin路由器
	router := gin.Default()

	// 添加HTTP服务器监听8080端口
	router.GET("/api/ping", func(c *gin.Context) {
		fmt.Printf("Received request: %s\n", c.Request.URL.Path)
		if time.Now().After(startTime.Add(5 * time.Minute)) {
			time.Sleep(10 * time.Second)
		}
		c.String(http.StatusOK, "Hello, World1111!")
	})

	// 添加新的处理函数来返回进程启动时间
	router.GET("/start-time", func(c *gin.Context) {
		c.String(http.StatusOK, "Process started at: %s", startTime.Format(time.RFC3339))
	})

	// 添加处理函数来返回所有环境变量
	router.GET("/env", func(c *gin.Context) {
		envs := os.Environ()
		envString := strings.Join(envs, "\n")
		c.String(http.StatusOK, "Environment Variables:\n%s", envString)
	})

	fmt.Println("Server starting on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}