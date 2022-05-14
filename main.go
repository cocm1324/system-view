package main

import (
	"log"
	"net/http"
)

const PORT string = "3000"

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("web"))
	mux.Handle("/app", http.StripPrefix("/app", fs))
	mux.Handle("/app/", http.StripPrefix("/app/", fs))
	log.Printf("server on %s", PORT)
	http.ListenAndServe(":"+PORT, mux)
}
