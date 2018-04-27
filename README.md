# HttpRouter

A Frontend Webserver that processes incoming http/https requests and routes 
them to the appropriate backend webserver.

## Getting Started

```go run cmd/configHelper/main.go -o config.yml```

Edit the _config.yml_ according to your needs.

```go run cmd/HttpRouter/main.go -c config.yml```

Open a browser and navigate to the given address and port.

Off course you need to have your DNS entries setup correctly.
A records of your subdomains should point to the listening address configured.

If you use TLS you should have a wildcard certificate, e.g. from Let's Encrypt.
