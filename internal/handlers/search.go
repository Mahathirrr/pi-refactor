// Package handlers berisi handler HTTP
package handlers

import (
	"net/http"
	"strconv"

	"info-retrieval/internal/config"
	"info-retrieval/internal/search"

	"github.com/gin-gonic/gin"
)

// SearchHandler menangani request POST pencarian
func SearchHandler(c *gin.Context) {
	query := c.PostForm("query")
	method := c.PostForm("method")
	c.Redirect(http.StatusFound, "/search?q="+query+"&method="+method+"&page=1")
}

// SearchHandlerGet menangani request GET pencarian dengan pagination
func SearchHandlerGet(c *gin.Context) {
	cfg := config.LoadConfig()
	query := c.Query("q")
	method := c.Query("method")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	allResults := search.Search(query, method)
	totalResults := len(allResults)
	totalPages := (totalResults + cfg.ItemsPerPage - 1) / cfg.ItemsPerPage

	if page < 1 {
		page = 1
	} else if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	var pagedResults []search.SearchResult
	if totalResults > 0 {
		start := (page - 1) * cfg.ItemsPerPage
		end := start + cfg.ItemsPerPage
		if end > totalResults {
			end = totalResults
		}
		pagedResults = allResults[start:end]
	}

	c.HTML(http.StatusOK, "results.html", gin.H{
		"results":      pagedResults,
		"query":        query,
		"method":       method,
		"currentPage":  page,
		"totalPages":   totalPages,
		"totalResults": totalResults,
		"previousPage": page - 1,
		"nextPage":     page + 1,
		"showPrevious": page > 1,
		"showNext":     page < totalPages,
	})
}