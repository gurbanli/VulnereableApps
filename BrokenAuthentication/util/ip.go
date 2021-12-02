package util

import (
	"log"
	"net"
	"net/http"
)

func GetIpAddressOfClient(req *http.Request) string{
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil{
		log.Fatal(err)
	}
	return ip
}

func GetIpAddressOfClientVulnerable(req *http.Request) string {
	var err error
	ip := req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, err = net.SplitHostPort(req.RemoteAddr)
		if err != nil{
			log.Fatal(err)
		}
	}
	return ip
}