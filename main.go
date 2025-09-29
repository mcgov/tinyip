// SPDX-License-Identifier: MIT
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Define command-line flags
	ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
	port := flag.Int("port", 8443, "Port to listen on")
	certFile := flag.String("cert", "cert.pem", "TLS certificate file")
	keyFile := flag.String("key", "key.pem", "TLS private key file")
	flag.Parse()

	// Define the address to listen on
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	// Create a handler that echoes back the client's IP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the client IP address
		clientIP := r.RemoteAddr
		// If the request is coming through a proxy that sets X-Forwarded-For,
		// use that instead (optional)
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			clientIP = forwardedFor
		}

		log.Printf("Request from %s", clientIP)
		IpOnly := strings.Split(clientIP, ":")[0]
		fmt.Fprintf(w, "%s\n", IpOnly)
	})

	// Set up the TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	// Create the HTTPS server
	server := &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
	}

	// Start the server
	log.Printf("Starting HTTPS server on %s", addr)
	log.Printf("Generate self-signed certificates if needed: openssl req -x509 -newkey rsa:4096 -keyout %s -out %s -days 365 -nodes", *keyFile, *certFile)
	err := server.ListenAndServeTLS(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("Error starting HTTPS server: %v", err)
	}
}
