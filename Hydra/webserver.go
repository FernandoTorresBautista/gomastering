package main

import (
	"fmt"
	"net/http"
)

//

//
// HandleFunc
// ListenAndServe

func main() {
	fmt.Println("web server hydra example")
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the hydra software system")
}
