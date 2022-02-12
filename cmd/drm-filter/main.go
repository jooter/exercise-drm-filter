package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jooter/exercise-drm-filter/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.DrmfilterHandler)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("default port will be used")
		port = "8080"
	}
	addr := ":" + port
	log.Println("listen to", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
