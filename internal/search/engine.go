// Package search berisi implementasi mesin pencari
package search

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"info-retrieval/pkg/logger"
)

// SearchEngine adalah struct utama mesin pencari
type SearchEngine struct {
	Articles      []Article
	InvertedIndex *InvertedIndex
	TFIDFScores   map[string]map[int]float64
	mutex         sync.RWMutex
}

var (
	searchEngine *SearchEngine
	once         sync.Once
)

// GetEngine mengembalikan instance singleton SearchEngine
func GetEngine() *SearchEngine {
	once.Do(func() {
		searchEngine = &SearchEngine{}
		if err := searchEngine.Initialize(); err != nil {
			logger.Fatal("Failed to initialize search engine: %v", err)
		}
	})
	return searchEngine
}

// Initialize menginisialisasi SearchEngine
func (se *SearchEngine) Initialize() error {
	articles, err := loadArticles()
	if err != nil {
		return err
	}

	invertedIndex := buildInvertedIndex(articles)
	tfidfScores := calculateTFIDF(invertedIndex, len(articles))

	se.mutex.Lock()
	se.Articles = articles
	se.InvertedIndex = invertedIndex
	se.TFIDFScores = tfidfScores
	se.mutex.Unlock()

	return nil
}

// Refresh memperbarui data SearchEngine
func (se *SearchEngine) Refresh() error {
	return se.Initialize()
}

// loadArticles memuat artikel dari file JSON
func loadArticles() ([]Article, error) {
	data, err := ioutil.ReadFile("articles.json")
	if err != nil {
		logger.Error("Error reading articles.json: %v", err)
		return nil, err
	}

	var articles []Article
	if err := json.Unmarshal(data, &articles); err != nil {
		logger.Error("Error parsing JSON from articles.json: %v", err)
		return nil, err
	}

	return articles, nil
}