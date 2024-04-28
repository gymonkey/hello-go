package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format(time.RFC3339))
		time.Sleep(10 * time.Second)
	}
}
