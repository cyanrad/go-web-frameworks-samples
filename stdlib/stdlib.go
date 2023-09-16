package stdlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webframeworks/coffee"
)

type stdlibCoffee struct {
	db coffee.CoffeeDB
}

func Main(cdb coffee.CoffeeDB) {
	cdb.Init()
	slc := stdlibCoffee{db: cdb}

	http.HandleFunc("/coffee/", slc.handleCoffee)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
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
	path := strings.TrimPrefix(r.URL.Path, "/coffee/")

	switch path {
	case "":
		writeJson(w, slc.db)

	case "avg":
		writeJson(w, slc.db.Avg())

	default:
		ID, err := strconv.Atoi(path)
		if err != nil {
			log.Fatal(err)
		}
		coffee, ok := slc.db.Get(ID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		}
		writeJson(w, coffee)
	}
}

func (slc stdlibCoffee) handleCoffeePost(w http.ResponseWriter, r *http.Request) {}

func (slc stdlibCoffee) handleCoffeePut(w http.ResponseWriter, r *http.Request) {}

func (slc stdlibCoffee) handleCoffeeDelete(w http.ResponseWriter, r *http.Request) {}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
