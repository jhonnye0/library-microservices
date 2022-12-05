package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiHost   = "localhost:8080"
	watermark = "localhost:8081"
	hello     = "localhost:8082"
)

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te", // canonicalized version of "TE"
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func delHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}

func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

type Proxy struct {
	Servers map[string]string
}

func (p *Proxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	log.Println(req.RemoteAddr, req.Method, req.URL)

	if req.URL.Scheme != "http" && req.URL.Scheme != "https" && req.Host != apiHost {
		log.Printf("Invalid scheme %s", req.URL.Scheme)
		log.Printf("Host: %s", req.Host)
		msg := "unsupported protocal scheme " + req.URL.Scheme
		http.Error(wr, msg, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	req.RequestURI = ""

	delHopHeaders(req.Header)

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		appendHostToXForwardHeader(req.Header, clientIP)
		fmt.Println("clientIP: ", clientIP)

		if strings.Contains(req.URL.String(), "/api/v1") {

			str := strings.SplitN(req.URL.Path[8:], "/", 2)
			service := str[0]
			path := "/" + str[1]

			if host, ok := p.Servers[service]; ok {
				req.URL = &url.URL{
					Scheme: "http",
					Host:   host,
					Path:   path,
				}

				resp, err := client.Do(req)
				if err != nil {
					http.Error(wr, "Server Error", http.StatusInternalServerError)
					log.Fatal("ServeHTTP: ", err)
				}
				defer resp.Body.Close()

				log.Println(req.RemoteAddr, " ", resp.Status)

				delHopHeaders(resp.Header)

				copyHeader(wr.Header(), resp.Header)
				wr.WriteHeader(resp.StatusCode)
				io.Copy(wr, resp.Body)

			} else {
				log.Fatal("Invalid path")
			}

			return
		}
	} else {
		http.Error(wr, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP: ", "Invalid HostPort")
	}

}

func main() {
	servers := map[string]string{"watermark": watermark, "hello": hello}

	var addr = flag.String("addr", "127.0.0.1:8080", "The addr of the application.")
	flag.Parse()

	handler := &Proxy{
		Servers: servers,
	}

	log.Println("Starting proxy server on", *addr)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
