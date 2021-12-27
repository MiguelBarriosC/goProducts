package volumes

import "github.com/gin-gonic/gin"

func HandleVolumes(c *gin.Context) {
	c.File("volumen_list.json")
}
