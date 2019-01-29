package presentation

import (
	"net/http"

	"github.com/KendoCross/kendocqrs/application"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	signProtocolAPI := "/api/sign_protocol"
	router.POST(signProtocolAPI, handles(signProtocolAPI))

	return router
}

func handles(command string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		result := application.HandCommand(data, command)
		c.JSON(http.StatusOK, result)
	}
}
