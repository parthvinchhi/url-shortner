package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/url-shortner/pkg/handlers"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("pkg/templates/*")

	r.GET("/", handlers.LoadIndexHandler)
	r.POST("/url", handlers.FetchUrlHandler)

	return r
}
