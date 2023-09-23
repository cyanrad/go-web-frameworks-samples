package gin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webframeworks/coffee"

	"github.com/gin-gonic/gin"
)

type ginCoffee struct {
	db *coffee.CoffeeDB
}

func Main(cdb *coffee.CoffeeDB) {
	ginc := ginCoffee{db: cdb}
	r := gin.Default()
	r.GET("/coffee", ginc.handleCoffeeGet)
	r.GET("/coffee/:id", ginc.handleCoffeeGetById)
	r.GET("/coffee/avg", ginc.handleCoffeeGetAvg)

	r.POST("/coffee", ginc.handleCoffeePost)

	r.PATCH("/coffee", ginc.handleCoffeePatch)

	r.DELETE("/coffee", ginc.handleCoffeeDelete)

	r.Run(":8080")
}

func (ginc ginCoffee) handleCoffeeGet(c *gin.Context) {
	c.JSON(http.StatusOK, *ginc.db)
}

func (ginc ginCoffee) handleCoffeeGetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	coffee, ok := ginc.db.Get(id)
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, coffee)
}

func (ginc ginCoffee) handleCoffeeGetAvg(c *gin.Context) {
}

func (ginc ginCoffee) handleCoffeePost(c *gin.Context) {
}

func (ginc ginCoffee) handleCoffeeDelete(c *gin.Context) {
}

func (ginc ginCoffee) handleCoffeePatch(c *gin.Context) {
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
