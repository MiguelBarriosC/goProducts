package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelBarriosC/goProducts/models"
	"github.com/gin-gonic/gin"
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
	//us := new(User)
	u, err := GetOne(id)

	if err != nil {
		fmt.Print(err)
		http.Error(c.Writer, "No existe!", http.StatusInternalServerError)
		return
	}

	c.JSON(200, u)
}

func CreateEndpoint(c *gin.Context) {
	var product models.Product

	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	e := Create(product)

	if e != nil {
		http.Error(c.Writer, e.Error(), http.StatusInternalServerError)
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
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	e := Update(p)

	if e != nil {
		http.Error(c.Writer, e.Error(), http.StatusInternalServerError)
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
		http.Error(c.Writer, "Error al eliminar producto con id: "+id, http.StatusInternalServerError)
		return
	}
	c.JSON(200, gin.H{
		"message": "DELETE OK",
	})
}
