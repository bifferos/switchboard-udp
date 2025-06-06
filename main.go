package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	StateDir string `json:"stateDir"`
	Port     int    `json:"port"`
}

var (
	stateDir = filepath.Join("/var/lib/switchboard")
)

func main() {

	// Some defaults, for missing config.
	cfg := Config{
		StateDir: stateDir,
		Port:     6060,
	}

	configPath := flag.String("config", filepath.Join("/etc/switchboard-udp", "config.json"), "Path to config file")
	flag.Parse()

	f, err := os.Open(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file not found: %v\n", err)
	} else {
		if err := json.NewDecoder(f).Decode(&cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding config: %v\n", err)
			os.Exit(1)
		}
	}
	defer f.Close()

	fmt.Println("Using configuration:")
	fmt.Printf("  StateDir: %s\n", cfg.StateDir)

	addr := net.UDPAddr{
		Port: cfg.Port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	fmt.Println("UDP server listening on", addr.String())

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		received := strings.TrimSpace(string(buf[:n]))
		// Sanitize the received data to avoid directory traversal
		received = filepath.Base(received)
		filePath := filepath.Join(cfg.StateDir, received)

		_, err = os.Stat(filePath)
		var response []byte
		if err == nil {
			response = []byte("true")
		} else if os.IsNotExist(err) {
			response = []byte("false")
		} else {
			response = []byte("false")
		}

		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Println("Write error:", err)
		}
	}
}
