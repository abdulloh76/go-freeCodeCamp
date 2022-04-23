package main

import (
	"book-store/pkg/routes"
	"book-store/pkg/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := utils.PORT
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
