// Package search berisi implementasi mesin pencari
package search

import (
	"math"
)

// calculateTFIDF menghitung skor TF-IDF untuk semua term dalam corpus
func calculateTFIDF(invertedIndex *InvertedIndex, totalDocs int) map[string]map[int]float64 {
	tfidfScores := make(map[string]map[int]float64)

	for term, postingList := range invertedIndex.Index {
		tfidfScores[term] = make(map[int]float64)

		// Hitung IDF untuk term
		idf := math.Log(float64(totalDocs) / float64(postingList.DocFrequency))

		// Hitung TF-IDF untuk setiap dokumen
		for docID, posting := range postingList.Postings {
			tf := float64(posting.Frequency)
			tfidfScores[term][docID] = tf * idf
		}
	}

	return tfidfScores
}

// calculateTermFrequency menghitung frekuensi term dalam dokumen
func calculateTermFrequency(term string, docID int, invertedIndex *InvertedIndex) float64 {
	if postingList, exists := invertedIndex.Index[term]; exists {
		if posting, exists := postingList.Postings[docID]; exists {
			return float64(posting.Frequency)
		}
	}
	return 0
}

// calculateIDF menghitung Inverse Document Frequency untuk term
func calculateIDF(term string, totalDocs int, invertedIndex *InvertedIndex) float64 {
	if postingList, exists := invertedIndex.Index[term]; exists {
		return math.Log(float64(totalDocs) / float64(postingList.DocFrequency))
	}
	return 0
}