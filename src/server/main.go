package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>Welcome to the future site of another blog</h2>")
}

func contact(w http.ResponseWriter, r * *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
    "to <a href=\"mailto:somewhere@nowhere.com\">"+
    "thisgoesnowhere@lost.confused</a>.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}