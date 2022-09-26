package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const spinProxyPrefix = "/api"
const spinitronBaseURL = "https://spinitron.com/api"
const port = "8080"

func spinProxyHandler(w http.ResponseWriter, r *http.Request) {

	// API is read-only
	if r.Method != http.MethodGet {
		log.Printf("unsupported HTTP method: %s", r.Method)
		return
	}

	path := r.RequestURI[len(spinProxyPrefix):]
	client := &http.Client{
		Timeout: time.Minute,
	}

	spinReq, err := http.NewRequest(http.MethodGet, spinitronBaseURL+path, nil)
	if err != nil {
		log.Println(err)
	}

	spinReq.Header = http.Header{
		"Authorization": {"Bearer " + os.Getenv("SPINITRON_API_KEY")},
		"accept":        {"application/json"},
	}

	// API is read-only
	spinReq.Method = http.MethodGet

	resp, err := client.Do(spinReq)
	if err != nil {
		log.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode >= 400 {
		log.Printf("unexpected status code %d: %s", resp.StatusCode, r.RequestURI)
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}

func main() {
	if os.Getenv("SPINITRON_API_KEY") == "" {
		panic("SPINITRON_API_KEY is empty")
	}

	http.HandleFunc(spinProxyPrefix+"/", spinProxyHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
