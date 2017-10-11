package main

import (
	"log/syslog"
	"os"
	"time"
)

func main() {
	hostname, _ := os.Hostname()
	syslog.Info("Starting business application")
	for {
		time.Sleep(10 * time.Second)
		syslog.Info("hello from %s", hostname)
	}
}
