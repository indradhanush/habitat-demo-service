package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type config struct {
	Addr string `json:"addr"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No config file given\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <path to config>\n", os.Args[0])
		os.Exit(1)
	}

	configBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config from %q: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	var config config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal config: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006")

		fmt.Fprintf(w, "Hello from habitat demo service. The time is: %s\n", message)
	})

	if err = http.ListenAndServe(config.Addr, nil); err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "HTTP server error: %v\n", err)
		os.Exit(1)
	}
}
