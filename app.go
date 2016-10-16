package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	fmt.Printf("Server running on port: %d", port);
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/", Index)
	router.HandleFunc("/auth/signin", Auth)
	router.HandleFunc("/auth/fbverify", Auth)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func fbVerify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
