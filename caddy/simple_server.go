package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping)
	addr := "0.0.0.0:8000"
	fmt.Printf("starting http server on %s\n", addr)
	http.ListenAndServe(addr, r)
}
