package main

import (
	"webframeworks/coffee"
	"webframeworks/stdlib"
)

func main() {
	var cdb coffee.CoffeeDB
	cdb.Init()

	stdlib.Main(&cdb)
}
