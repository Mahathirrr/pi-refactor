// Package search berisi implementasi mesin pencari
package search

import (
	"math"
)

// CosineSimilarity menghitung similarity antara query dan dokumen menggunakan cosine similarity
func cosineSimilarityWithTFIDF(queryVector map[string]float64, tfidfScores map[string]map[int]float64, docID int) float64 {
	docVector := make(map[string]float64)
	for term, scores := range tfidfScores {
		if score, exists := scores[docID]; exists {
			docVector[term] = score
		}
	}

	normalizedQuery := normalizeVector(queryVector)
	normalizedDoc := normalizeVector(docVector)

	var dotProduct, queryMagnitude, docMagnitude float64

	for term, queryWeight := range normalizedQuery {
		queryMagnitude += queryWeight * queryWeight
		if docWeight, exists := normalizedDoc[term]; exists {
			dotProduct += queryWeight * docWeight
		}
	}

	for _, docWeight := range normalizedDoc {
		docMagnitude += docWeight * docWeight
	}

	queryMagnitude = math.Sqrt(queryMagnitude)
	docMagnitude = math.Sqrt(docMagnitude)

	if queryMagnitude == 0 || docMagnitude == 0 {
		return 0
	}

	return dotProduct / (queryMagnitude * docMagnitude)
}

// JaccardSimilarity menghitung similarity antara query dan dokumen menggunakan Jaccard similarity
func jaccardSimilarityWithTFIDF(queryVector map[string]float64, tfidfScores map[string]map[int]float64, docID int) float64 {
	querySet := make(map[string]bool)
	docSet := make(map[string]bool)

	for term := range queryVector {
		querySet[term] = true
	}

	for term, scores := range tfidfScores {
		if _, exists := scores[docID]; exists {
			docSet[term] = true
		}
	}

	intersection := 0
	for term := range querySet {
		if docSet[term] {
			intersection++
		}
	}

	union := len(querySet) + len(docSet) - intersection
	if union == 0 {
		return 0
	}

	return float64(intersection) / float64(union)
}

// normalizeVector menormalisasi vektor untuk perhitungan similarity
func normalizeVector(vector map[string]float64) map[string]float64 {
	normalized := make(map[string]float64)
	var magnitude float64

	for _, weight := range vector {
		magnitude += weight * weight
	}
	magnitude = math.Sqrt(magnitude)

	if magnitude > 0 {
		for term, weight := range vector {
			normalized[term] = weight / magnitude
		}
	}

	return normalized
}