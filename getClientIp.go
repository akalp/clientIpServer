package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	port = flag.String("port", "8080", "the port to listen on")
	cert = flag.String("cert", "", "the path of cert pem file")
	key  = flag.String("key", "", "the path of key file")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tip := r.Header.Get("X-Real-IP")
		if tip == "" {
			tip = r.Header.Get("X-Forwarded-For")
			if tip == "" {
				tip = r.RemoteAddr
			}
		}

		ip, _, err := net.SplitHostPort(tip)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"ip": ip,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		setupCORS(&w, r)
		w.Write(jsonData)
	})

	addr := ":" + *port
	fmt.Printf("Serving on port %s\n", *port)
	if len(*cert) == 0 || len(*key) == 0 {
		if err := http.ListenAndServe(addr, nil); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServeTLS(addr, *cert, *key, nil); err != nil {
			panic(err)
		}
	}
}

func init() {
	if p := os.Getenv("HNIS_PORT"); p != "" {
		port = &p
	}
	if k := os.Getenv("HNIS_TLS_KEY"); k != "" {
		key = &k
	}
	if c := os.Getenv("HNIS_TLS_CERT"); c != "" {
		cert = &c
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
