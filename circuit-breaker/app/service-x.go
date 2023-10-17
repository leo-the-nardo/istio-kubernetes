package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", run)
	http.ListenAndServe(":8000", nil)
}

func run(res http.ResponseWriter, req *http.Request) {
	if os.Getenv("error") == "yes" {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		res.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("OK"))
}
