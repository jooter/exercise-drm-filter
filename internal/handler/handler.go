package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jooter/exercise-drm-filter/pkg/drmfilter"
)

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func DrmfilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := "un-supported request method"
		log.Println(msg)
		errorResponse(w, msg)
		return
	}

	var req drmfilter.Request
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		fmt.Println(err)
		errorResponse(w, err.Error())
		return
	}

	resp := drmfilter.Filter(&req)
	b, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err)
		errorResponse(w, err.Error())
		return
	}

	fmt.Fprintln(w, string(b))
}

func errorResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	resp := ErrorResponse{Error: msg}
	b, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, string(b))
}
