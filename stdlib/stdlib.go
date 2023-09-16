package stdlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"webframeworks/coffee"
)

type stdlibCoffee struct {
	db *coffee.CoffeeDB
}

func Main(cdb *coffee.CoffeeDB) {
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
	case http.MethodPatch:
		slc.handleCoffeePatch(w, r)
	case http.MethodDelete:
		slc.handleCoffeeDelete(w, r)
	}
}

func (slc stdlibCoffee) handleCoffeeGet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/coffee/")

	switch path {
	case "":
		err := writeJson(w, slc.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "avg":
		err := writeJson(w, slc.db.Avg())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		ID, err := strconv.Atoi(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		coffee, ok := slc.db.Get(ID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = writeJson(w, coffee)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (slc stdlibCoffee) handleCoffeePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/coffee/")

	switch path {
	case "":
		defer r.Body.Close()
		newCoffee := coffee.Coffee{}
		err := json.NewDecoder(r.Body).Decode(&newCoffee)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		slc.db.Create(newCoffee)
		err = writeJson(w, newCoffee)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (slc stdlibCoffee) handleCoffeeDelete(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/coffee/")

	switch path {
	default:
		ID, err := strconv.Atoi(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		ok := slc.db.Delete(ID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

}

func (slc stdlibCoffee) handleCoffeePatch(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/coffee/")

	switch path {
	default:
		ID, err := strconv.Atoi(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		coffee, ok := slc.db.Get(ID)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body.Close()

		// this is bad cuz if we ever add anyting that's not an int it will error
		newCoffeeValues := map[string]int{}
		err = json.Unmarshal(bytes, &newCoffeeValues)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if val, ok := newCoffeeValues["instant_coffee"]; ok {
			coffee.InstantCoffee = val
		}
		if val, ok := newCoffeeValues["coffee_mate"]; ok {
			coffee.CoffeeMate = val
		}
		if val, ok := newCoffeeValues["powdered_milk"]; ok {
			coffee.PowderedMilk = val
		}
		if val, ok := newCoffeeValues["evaporated_milk"]; ok {
			coffee.EvaporatedMilk = val
		}
		if val, ok := newCoffeeValues["water"]; ok {
			coffee.Water = val
		}
		if val, ok := newCoffeeValues["rating"]; ok {
			coffee.Rating = val
		}

		slc.db.Set(ID, coffee)
		writeJson(w, coffee)
	}
}

func writeJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
