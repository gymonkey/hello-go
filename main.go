package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/localip", func(writer http.ResponseWriter, request *http.Request) {
		ifaces, err := net.Interfaces()
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
			return
		}
		// handle err
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(err.Error()))
				return
			}
			ips := make([]string, 0)
			// handle err
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ips = append(ips, ip.String())
				// process IP address
			}
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(strings.Join(ips, ",")))
			return
		}
	})
	fmt.Println("listen port: " + os.Getenv("PORT"))
	http.ListenAndServe(strings.Join([]string{"", os.Getenv("PORT")}, ":"), nil)
}
