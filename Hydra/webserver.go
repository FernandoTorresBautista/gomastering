package main

import (
	"GoMastering/Hydra/hlogger"
	"fmt"
	"net/http"
)

//

//
// HandleFunc
// ListenAndServe

//
// loogger with singleton pattern

func main() {
	logger := *hlogger.GetInstance()
	logger.Println("Starting hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := *hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the hydra software system")

	logger.Println("Received an http request on root url")
}
