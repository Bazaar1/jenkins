package app

import (
	"apiserver/domain"
	"apiserver/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/test", test)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}