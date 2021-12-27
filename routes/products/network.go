package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelBarriosC/goProducts/models" // Package models
	"github.com/gin-gonic/gin"                    // Gin
)

func GetEndpoint(c *gin.Context) {
	u, err := GetAll()

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(200, u)
}

func GetOneEndpoint(c *gin.Context) {
	id := c.Param("id")
	p, err := GetOne(id)

	if err != nil {
		fmt.Print(err)
		c.JSON(404, gin.H{
			"error":   true,
			"message": "No existe!",
		})
		return
	}

	c.JSON(200, p)
}

func CreateEndpoint(c *gin.Context) {
	var product models.Product

	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	e := Create(product)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": e.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "INSERTED OK",
	})
}

func UpdateEndpoint(c *gin.Context) {
	var p models.Product

	err := json.NewDecoder(c.Request.Body).Decode(&p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	e := Update(p)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": e.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "UPDATE OK",
	})
}

func DeleteEndpoint(c *gin.Context) {
	id := c.Param("id")

	err := Delete(id)
	if !err {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error al eliminar producto con id: " + id,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "DELETE OK",
	})
}
