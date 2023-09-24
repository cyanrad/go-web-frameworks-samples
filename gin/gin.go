package gin

import (
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

	r.PATCH("/coffee/:id", ginc.handleCoffeePatch)

	r.DELETE("/coffee/:id", ginc.handleCoffeeDelete)

	r.Run(":8080")
}

func (ginc ginCoffee) handleCoffeeGet(c *gin.Context) {
	c.JSON(http.StatusOK, ginc.db)
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
	c.JSON(http.StatusOK, ginc.db.Avg())
}

func (ginc ginCoffee) handleCoffeePost(c *gin.Context) {
	newCoffee := coffee.Coffee{}
	err := c.ShouldBindJSON(&newCoffee)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ginc.db.Create(newCoffee)
	c.JSON(http.StatusCreated, newCoffee)
}

func (ginc ginCoffee) handleCoffeePatch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	newCoffee := coffee.CoffeePatch{}
	err = c.ShouldBindJSON(&newCoffee)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	patchedCoffee, ok := ginc.db.Patch(id, newCoffee)
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, patchedCoffee)
}

func (ginc ginCoffee) handleCoffeeDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	ok := ginc.db.Delete(id)
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}
