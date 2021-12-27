package main

import (
	"github.com/MiguelBarriosC/goProducts/middelware"      //middelwares
	"github.com/MiguelBarriosC/goProducts/routes/products" // Route products
	"github.com/MiguelBarriosC/goProducts/routes/volumes"  // Route volumes
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/products")
	{
		v1.GET("/", products.GetEndpoint)
		v1.GET("/:id", products.GetOneEndpoint)
		v1.POST("/", products.CreateEndpoint)
		v1.PUT("/", products.UpdateEndpoint)
		v1.DELETE("/:id", products.DeleteEndpoint)
	}

	r.GET("volumes", middelware.BasicAuth, volumes.HandleVolumes)

	r.Run(":9098") // listen and serve on 0.0.0.0:9098 (for windows "localhost:9098")
}
