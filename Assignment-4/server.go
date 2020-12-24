package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	gender := r.FormValue("gender")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Phone Number = %s\n", phone)
	fmt.Fprintf(w, "Gender = %s\n", gender)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fileServer := http.FileServer(http.Dir("./")) // New code
	router.Handle("/", fileServer)                // New code

	router.HandleFunc("/hello", hellohandler)
	router.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
