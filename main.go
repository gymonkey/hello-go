package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
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
		ips := make([]string, 0)
		for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
				writer.WriteHeader(http.StatusOK)
				writer.Write([]byte(err.Error()))
				return
			}
			// handle err
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip.To4() != nil {
					ips = append(ips, ip.String())
				}
				// process IP address
			}
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(strings.Join(ips, ",")))
		return
	})
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		pinger, err := ping.NewPinger("www.google.com")
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
			return
		}
		pinger.Count = 3
		pinger.Timeout = 1 * time.Minute
		err = pinger.Run()
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
			return
		}
		stats := pinger.Statistics()
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("%s %s %s %f\n", stats.MaxRtt, stats.MinRtt, stats.AvgRtt, stats.PacketLoss)))
	})
	//os.Setenv("PORT", "18080")
	fmt.Println("listen port: " + os.Getenv("PORT"))
	http.ListenAndServe(strings.Join([]string{"", os.Getenv("PORT")}, ":"), nil)
}
