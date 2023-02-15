package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var port = flag.String("port", "8080", "the port to listen on")

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
		w.Write(jsonData)
	})

	addr := ":" + *port
	fmt.Printf("Serving on port %s\n", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func init() {
	if p := os.Getenv("HNIS_PORT"); p != "" {
		port = &p
	}
}
