package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/starlight-go/starlight"
)

func main() {
	http.HandleFunc("/", handle)
	port := ":8080"
	fmt.Printf("runing web server on http://localhost%v?name=starlight&repeat=3\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handing request ", r.URL)

	globals := map[string]interface{}{
		"r":       r,
		"w":       w,
		"Fprintf": fmt.Fprintf,
	}

	_, err := starlight.Eval("handle.star", globals, nil)
	if err != nil {
		fmt.Println(err)
	}
}
