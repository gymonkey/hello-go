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
		ip, err := getLocalIp()
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(ip))
		return
	})
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		addr := request.Form.Get("addr")
		pinger, err := ping.NewPinger(addr)
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

func getLocalIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// handle err
	ips := make([]string, 0)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
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
	return ips[0], nil
}
