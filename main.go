package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	u, _ := url.Parse("http://localhost:9200")
	svc := httputil.NewSingleHostReverseProxy(u)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[INFO] incoming req", r.Method, r.Host, r.URL.Path, "...")
		svc.ServeHTTP(w, r)
	})

	fmt.Println(http.ListenAndServe(":80", nil))
}
