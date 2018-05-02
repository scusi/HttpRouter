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

## Example Config

An example config could look like this:
```
addr: 127.0.0.1
port: "9090"
tls: true
cert: /home/scusi/.scusi.io.pki/root/certs/test.scusi.io.crt
key: /home/scusi/.scusi.io.pki/root/keys/test.scusi.io.key
routes:
- subdomain: one.scusi.io
  backend: https://one.my-backend.intern/
- subdomain: two.scusi.io
  backend: http://two.my-backend.intern/
```

The above example config would start a HTTPS webserver at the _localhost_ interface on port _9090_.

It knows two subdomains, _one.scusi.io_ and _two.scusi.io_.
