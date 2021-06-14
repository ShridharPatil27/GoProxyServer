package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
)

// These constant is used to define server config
const (
	// keeping same url for both proxy_url_1 and proxy_url_2 as this just to demonstrate
	proxy_url_1        = "http://localhost:5000"
	proxy_url_2        = "http://localhost:5000"
	port               = "3000"
	cacheTimeInMinutes = 5
)

var (
	reqHeaders = map[string]string{
		"custom_request_header_by_proxy_1": "1111111111111",
		"custom_request_header_by_proxy_2": "1111111111111",
	}
	resHeaders = map[string]string{
		"custom_response_header_by_proxy_1": "000000000000",
		"custom_response_header_by_proxy_2": "000000000000",
	}
	responseReplacers = map[string]string{
		"pen": "pencil",
		"mug": "bottle",
	}
	severCount                   = 0
	_          http.RoundTripper = &transport{}
	cache                        = &bigcache.BigCache{}
)

type transport struct {
	http.RoundTripper
}

func init() {
	cacheObj, initErr := bigcache.NewBigCache(bigcache.DefaultConfig(cacheTimeInMinutes * time.Minute))
	if initErr != nil {
		log.Fatal(initErr)
	}
	cache = cacheObj
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	for k, v := range responseReplacers {
		b = bytes.Replace(b, []byte(k), []byte(v), -1)
	}
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	// set custom response headers
	addCustomResponseHeaders(resp)
	// print response headers
	printResHeaders(resp)

	// set cache for the response
	cache.Set(req.RequestURI, b)

	return resp, nil
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// user http.RoundTripper interface
	proxy.Transport = &transport{http.DefaultTransport}

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Balance returns one of the servers based using round-robin algorithm
func getProxyURL() string {
	var servers = []string{proxy_url_1, proxy_url_2}

	server := servers[severCount]
	severCount++

	// reset the counter and start from the beginning
	if severCount >= len(servers) {
		severCount = 0
	}

	return server
}

// Add custom header to the req object
func addCustomRequestHeaders(req *http.Request) {
	for k, v := range reqHeaders {
		req.Header.Set(k, v)
	}
}

// Add custom header to the res object
func addCustomResponseHeaders(res *http.Response) {
	for k, v := range resHeaders {
		res.Header.Set(k, v)
	}
}

// Add custom header to the response writer object
func addCustomResponseHeadersRW(res http.ResponseWriter) {
	for k, v := range resHeaders {
		res.Header().Set(k, v)
	}
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

	entry, err := cache.Get(req.RequestURI)
	if err == nil && len(entry) != 0 {
		// return from cache
		addCustomResponseHeadersRW(res)
		fmt.Println("Served from cache")
		printResHeadersRW(res)
		fmt.Fprintln(res, string(entry))
		return
	}

	url := getProxyURL()

	// set custom request headers
	addCustomRequestHeaders(req)

	serveReverseProxy(url, res, req)
}

func printResHeaders(res *http.Response) {
	for name, values := range res.Header {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
}

func printResHeadersRW(res http.ResponseWriter) {
	for name, values := range res.Header() {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
