package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/weitecklee/nlp"
)

func main() {
	// routing
	// /health is exact match
	// /health/ is prefix match
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	// run server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Step 1: Get, convert & validate the data
	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1_000_000) // limit size of reader for safety purposes
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "Can't read", http.StatusBadRequest)
		return
	}
	if len(data) == 0 {
		http.Error(w, "Missing data", http.StatusBadRequest)
		return
	}
	text := string(data)

	// Step 2: Work
	tokens := nlp.Tokenize(text)

	// Step 3: Encode & emit output
	resp := map[string]any{
		"tokens": tokens,
	}
	//
	// err = json.NewEncoder(w).Encode(resp)
	// One downside of above method is response is sent
	// before error is caught
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "Can't encode", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

// shortcut "hand" for http handler declaration

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check
	fmt.Fprintln(w, "OK")
}

// hey -n 10000 http://localhost:8080/health
