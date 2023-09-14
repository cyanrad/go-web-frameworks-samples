package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webframeworks/coffee"
)

type stdlibCoffee struct {
	db coffee.CoffeeDB
}

func StdlibMain(cdb coffee.CoffeeDB) {
	cdb.Init()
	slc := stdlibCoffee{db: cdb}

	http.HandleFunc("/coffee", slc.handleCoffee)

	// starting the server
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) { // graceful shutdown
		fmt.Printf("server closed\n")
	} else if err != nil { //-- fuck you shutdown
		fmt.Printf("error starting server: %s\n", err)
	}
}

func (slc stdlibCoffee) handleCoffee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		slc.handleCoffeeGet(w, r)
	case http.MethodPost:
		slc.handleCoffeePost(w, r)
	case http.MethodPut:
		slc.handleCoffeePut(w, r)
	case http.MethodDelete:
		slc.handleCoffeeDelete(w, r)
	}
}

func (slc stdlibCoffee) handleCoffeeGet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		sendJson(w, slc.db)

	case "/avg":
		avgCoffee := slc.db.Avg()
		sendJson(w, avgCoffee)

	default:
		ID, err := strconv.Atoi(path[1:])
		if err != nil {
			log.Fatal(err)
		}
		c, ok := slc.db.Get(ID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		}
		sendJson(w, c)
	}
}
func (slc stdlibCoffee) handleCoffeePost(w http.ResponseWriter, r *http.Request)   {}
func (slc stdlibCoffee) handleCoffeePut(w http.ResponseWriter, r *http.Request)    {}
func (slc stdlibCoffee) handleCoffeeDelete(w http.ResponseWriter, r *http.Request) {}

func sendJson(w http.ResponseWriter, data any) {
	if bytes, err := json.Marshal(data); err != nil {
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
