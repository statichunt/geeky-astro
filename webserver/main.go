package main

import (
	"errors"
	"log"
	"net"
	"strings"

	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	ws := &Webserver{
		server: http.NewServeMux(),
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		port:   80,
		mw:     &CustomMiddleWare{logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))},
	}
	ws.RegisterMetrics()
	ws.RegisterStaticFileServer()
	ws.RegisterApiEndpoints()

	err := ws.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}
