// Package search berisi implementasi mesin pencari
package search

import "html/template"

// Article merepresentasikan artikel berita properti
type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

// SearchResult merepresentasikan hasil pencarian
type SearchResult struct {
	Title              string
	Content            string
	URL                string
	Score              float64
	HighlightedContent template.HTML
	Favicon            string
}

// InvertedIndex adalah struktur data untuk inverted index
type InvertedIndex struct {
	Index map[string]*PostingList
}

// PostingList menyimpan informasi dokumen untuk setiap term
type PostingList struct {
	DocFrequency int
	Postings     map[int]*Posting
}

// Posting menyimpan informasi term dalam dokumen
type Posting struct {
	DocID     int
	Frequency int
	Positions []int
}