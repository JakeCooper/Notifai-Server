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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func fbVerify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("FUCK")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
