package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(response, "method not supported for hello endpoint", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "this is hello endpoint and we are handling get method for it")
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "error while passing the form%s", err)
		return
	}

	fmt.Fprintf(response, "POST request was parsed successfully")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "\nname: %s", name)
	fmt.Fprintf(response, "\naddress: %s", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("starting server at localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
