// Package handlers berisi handler HTTP dan fungsi template
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler menangani request halaman utama
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}