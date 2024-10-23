package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parthvinchhi/url-shortner/pkg/models"
)

// LoadIndexHandler, it loads the index page
func LoadIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// FetchUrlHandler,
func FetchUrlHandler(c *gin.Context) {
	var u models.UrlStruct
	u.OldUrl = c.PostForm("url-to-use")
	fmt.Println(u.OldUrl)
}
