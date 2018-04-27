package main

import (
	"flag"
	"github.com/scusi/HttpRouter/datastructs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var outFile string

func init() {
	flag.StringVar(&outFile, "o", "exampleConfig.yml", "file to write example config to")
}

func main() {
	flag.Parse()
	c := defaultConfig()
	err := MarshalConfig(outFile, c)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}

func defaultConfig() *datastructs.Config {
	var s []datastructs.SubdomainRoute
	s = append(s, datastructs.SubdomainRoute{Subdomain: "one.scusi.io", Backend: "http://box.scusi.io/"})
	s = append(s, datastructs.SubdomainRoute{Subdomain: "two.scusi.io", Backend: "http://blog.fefe.de"})
	dConfig := &Config{
		Addr:   "127.0.0.1",
		TLS:    true,
		Port:   "443",
		Cert:   "cert.pem",
		Key:    "key.pem",
		Routes: s,
	}
	return dConfig
}

func MarshalConfig(filename string, c *datastructs.Config) (err error) {
	out, err := yaml.Marshal(c)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, out, 0600)
	if err != nil {
		return
	}
	return
}
