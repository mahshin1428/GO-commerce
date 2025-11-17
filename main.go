package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, " Tui amre chinbi na ami tor baper khaloto bhai er chachato bhai er sumondir polar chacha")
}

func getProducts(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "")
}


func main() {

	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000",  mux)

	if err != nil {
		fmt.Println("Error starting the server")
	}
}
