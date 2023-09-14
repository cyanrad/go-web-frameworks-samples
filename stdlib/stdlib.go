package main

import (
	"errors"
	"fmt"
	"net/http"
)

func StdlibMain() {
	http.HandleFunc("/coffee", handleCoffee)

	//-- starting the server
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) { //-- graceful shutdown
		fmt.Printf("server closed\n")
	} else if err != nil { //-- fuck you shutdown
		fmt.Printf("error starting server: %s\n", err)
	}
}

func handleCoffee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleCoffeeGet(w, r)
	case http.MethodPost:
		handleCoffeePost(w, r)
	case http.MethodPut:
		handleCoffeePut(w, r)
	case http.MethodDelete:
		handleCoffeeDelete(w, r)
	}
}

func handleCoffeeGet(w http.ResponseWriter, r *http.Request)    {}
func handleCoffeePost(w http.ResponseWriter, r *http.Request)   {}
func handleCoffeePut(w http.ResponseWriter, r *http.Request)    {}
func handleCoffeeDelete(w http.ResponseWriter, r *http.Request) {}
