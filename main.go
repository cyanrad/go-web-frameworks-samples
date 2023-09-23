package main

import (
	"webframeworks/coffee"
	"webframeworks/gin"
)

func main() {
	var cdb coffee.CoffeeDB
	cdb.Init()

	// stdlib.Main(&cdb)
	gin.Main(&cdb)
}
