// Package search berisi implementasi mesin pencari
package search

import (
	"html/template"
	"sort"

	"info-retrieval/internal/utils"
)

// Search melakukan pencarian dengan query dan metode tertentu
func Search(query string, method string) []SearchResult {
	engine := GetEngine()
	engine.mutex.RLock()
	defer engine.mutex.RUnlock()

	// Proses query
	queryTokens := GetProcessor().ProcessText(query)
	queryVector := make(map[string]float64)
	for _, token := range queryTokens {
		queryVector[token]++
	}

	var results []SearchResult

	// Hitung similarity untuk setiap artikel
	for i, article := range engine.Articles {
		var score float64
		switch method {
		case "cosine":
			score = cosineSimilarityWithTFIDF(queryVector, engine.TFIDFScores, i)
		case "jaccard":
			score = jaccardSimilarityWithTFIDF(queryVector, engine.TFIDFScores, i)
		default:
			score = cosineSimilarityWithTFIDF(queryVector, engine.TFIDFScores, i)
		}

		if score > 0 {
			contentPreview := utils.GetContentPreview(article.Content, query, 160)
			highlightedContent := utils.HighlightText(contentPreview, query)

			results = append(results, SearchResult{
				Title:              article.Title,
				Content:            contentPreview,
				URL:                article.URL,
				Score:              score,
				HighlightedContent: template.HTML(highlightedContent),
				Favicon:            utils.GetFaviconPath(article.URL),
			})
		}
	}

	// Urutkan hasil berdasarkan score
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}