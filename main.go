package main

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"net/http"
)

func main() {
	//http.DefaultClient.Timeout = 1 * time.Second
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	resp, err := http.DefaultClient.Do(request)
	//	if err != nil {
	//		writer.WriteHeader(http.StatusBadGateway)
	//		writer.Write([]byte(err.Error()))
	//	} else {
	//		defer resp.Body.Close()
	//		writer.WriteHeader(resp.StatusCode)
	//		for key, val := range resp.Header {
	//			writer.Header().Set(key, val[0])
	//		}
	//		io.Copy(writer, resp.Body)
	//	}
	//})
	//http.HandleFunc("/inner", func(writer http.ResponseWriter, request *http.Request) {
	//	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	//	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://ip.sb/", nil)
	//	req.Header.Set("User-Agent", "curl/8.7.1")
	//	resp, err := http.DefaultClient.Do(req)
	//	if err != nil {
	//		writer.WriteHeader(http.StatusOK)
	//		writer.Write([]byte(err.Error()))
	//	} else {
	//		defer resp.Body.Close()
	//		writer.WriteHeader(http.StatusOK)
	//		writer.Write([]byte(resp.Status))
	//		data, _ := io.ReadAll(resp.Body)
	//		writer.Write(data)
	//	}
	//})
	//http.HandleFunc("/outter", func(writer http.ResponseWriter, request *http.Request) {
	//	//fmt.Println(request.RemoteAddr)
	//	writer.WriteHeader(http.StatusOK)
	//	data, _ := io.ReadAll(request.Body)
	//	fmt.Println(request.Header.Get("Authorization"))
	//	fmt.Println(string(data))
	//})
	//http.ListenAndServe(":8080", nil)
	//for {
	//	fmt.Println("6" + strconv.FormatInt(time.Now().Unix(), 10))
	//	time.Sleep(10 * time.Second)
	//}
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello"})
	})
	r.Run()
}
