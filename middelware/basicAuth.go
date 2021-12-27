package middelware

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	// Obtiene Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()

	if !hasAuth || !checkUserAndPass(user, password) {
		c.Header("WWW-Authenticate", `Basic realm="Account invalid"`)
		c.JSON(401, gin.H{
			"error":   true,
			"message": "No authorizado!",
		})
		c.Abort()
		return
	}
	c.Next() // Continua al handler
}

func checkUserAndPass(username, password string) bool {
	return username == "admin" && password == "123"
}
