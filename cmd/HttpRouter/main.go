// creates a webserver that reverse proxies based on subdomain
//
// TODO: currently the host variable in the http request proxied is not made to match the backend host.
//       depending on the backend host configuration the backend server does not answer with the right content.
//       possible fix: https://stackoverflow.com/questions/21270945/how-to-read-the-response-from-a-newsinglehostreverseproxy
//
package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/scusi/HttpRouter/datastructs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var configFile string
var cfg *datastructs.Config

func init() {
	flag.StringVar(&configFile, "c", "config.yml", "config file in yaml format")
	// configure logging
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	flag.Parse()
	// read config
	log.Info("reading config file")
	ymlData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	// unmarshal config data
	err = yaml.Unmarshal(ymlData, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("starting up")

	router := mux.NewRouter()
	for _, route := range cfg.Routes {
		router.Host(route.Subdomain).Handler(http.Handler(NewReverseProxy(route.Backend)))
		log.WithFields(log.Fields{"Host": route.Subdomain, "Backend": route.Backend}).Info("Router added")
	}

	if cfg.TLS != true {
		log.Info("TLS is not configured")
		listenAddr := cfg.Addr + ":" + cfg.Port
		log.Info("start listening on " + listenAddr)
		log.Fatal(http.ListenAndServe(listenAddr, router))
	} else {
		log.Info("TLS is configured")
		listenAddr := cfg.Addr + ":" + cfg.Port
		log.Info("start listening on " + listenAddr)
		log.Fatal(http.ListenAndServeTLS(listenAddr, cfg.Cert, cfg.Key, router))
	}
}

// NewReverseProyx returns a HTTPHandler that is a reverseProxy for the given URL
// It ajusts the Host header via a custom ReverseProxy.Director function.
func NewReverseProxy(URL string) *httputil.ReverseProxy {
	// parse the supplied URL
	rpURL, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	// creating a Director Function for reverseProxy
	// this is neccessary to rewrite the Host Header in order to get the correct page.
	director := func(r *http.Request) {
		r.URL.Scheme = rpURL.Scheme
		r.URL.Host = rpURL.Host
		r.Host = rpURL.Host
		log.Debug("relay request to " + r.URL.String() + " for " + r.RemoteAddr)
	}
	// create the reverseProxy
	rproxy := httputil.NewSingleHostReverseProxy(rpURL)
	// apply the director function
	rproxy.Director = director
	return rproxy
}
