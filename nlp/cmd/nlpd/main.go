package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/weitecklee/nlp"
	"github.com/weitecklee/nlp/stemmer"
)

var (
	numTok = expvar.NewInt("tokenize.calls")
)

func main() {

	// Create server
	logger := log.New(log.Writer(), "nlp ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger, // dependency injection
	}
	// routing
	// /health is exact match
	// /health/ is prefix match
	// http.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/tokenize", tokenizeHandler)
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// run server
	addr := os.Getenv("NLPD_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	s.logger.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)

}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Before gorilla mux
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	*/

	numTok.Add(1)

	// Step 1: Get, convert & validate the data
	// defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1_000_000) // limit size of reader for safety purposes
	data, err := io.ReadAll(rdr)
	if err != nil {
		s.logger.Printf("Error: Can't read - %s", err)
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

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check
	fmt.Fprintln(w, "OK")
}

// hey -n 10000 http://localhost:8080/health

// with expvar, http://localhost:8080/debug/vars
