package datastructs

type SubdomainRoute struct {
	Subdomain string
	Backend   string
}

type Config struct {
	Addr   string
	Port   string
	TLS    bool
	Cert   string
	Key    string
	Routes []SubdomainRoute
}
