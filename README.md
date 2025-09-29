# tinyip

small golang https server that echos the client ip.

## prerequisites

- Go 1.16 or higher
- OpenSSL (or tls certificates from somewhere)

## test it out

### generate TLS certificates for testing:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

### start the server

```bash
# Default settings (0.0.0.0:8443)
go run main.go

# Custom IP and port
go run main.go --ip 127.0.0.1 --port 8443

# Custom certificate and key files
go run main.go --cert mycert.pem --key mykey.pem
```
### connect to the server

```bash
# Ignore certificate validation for self-signed certs
curl -k https://localhost:8443
```

Or visit `https://localhost:8443` in your browser (you'll need to accept the security warning for self-signed certificates).

## command line options

- `--ip`: IP address to bind to (default: "0.0.0.0")
- `--port`: Port to listen on (default: 8443)
- `--cert`: TLS certificate file (default: "cert.pem")
- `--key`: TLS private key file (default: "key.pem")

## authorship

copilot prompted by mamcgove at microsoft