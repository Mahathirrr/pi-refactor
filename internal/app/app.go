// Package app berisi logika utama aplikasi
package app

import (
	"info-retrieval/internal/handlers"
	"info-retrieval/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Run menginisialisasi dan menjalankan server aplikasi
func Run() {
	r := gin.Default()

	// Setup middleware
	middleware.Setup(r)

	// Setup routes
	setupRoutes(r)

	r.Run(":8080")
}

// setupRoutes mendaftarkan semua route aplikasi
func setupRoutes(r *gin.Engine) {
	// Serve static files
	r.Static("/static", "./static")

	// Setup template functions
	r.SetFuncMap(handlers.TemplateFunctions())
	r.LoadHTMLGlob("templates/*")

	// Register routes
	r.GET("/", handlers.IndexHandler)
	r.POST("/search", handlers.SearchHandler)
	r.GET("/search", handlers.SearchHandlerGet)
}

