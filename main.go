package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	u, _ := url.Parse("http://localhost:8080")
	svc := httputil.NewSingleHostReverseProxy(u)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[INFO] incoming req", r.Method, r.Host, r.URL.Path, "...")
		svc.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
