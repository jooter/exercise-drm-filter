package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("default port will be used")
		port = "8080"
	}
	addr := ":" + port
	log.Println("listen to", addr)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
