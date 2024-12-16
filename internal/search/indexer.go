// Package search berisi implementasi mesin pencari
package search

// buildInvertedIndex membangun inverted index dari kumpulan artikel
func buildInvertedIndex(articles []Article) *InvertedIndex {
	idx := NewInvertedIndex()

	for docID, article := range articles {
		// Gabungkan title dan content untuk indexing
		tokens := textProcessor.ProcessText(article.Title + " " + article.Content)

		// Track posisi untuk setiap term
		for pos, token := range tokens {
			if _, exists := idx.Index[token]; !exists {
				idx.Index[token] = &PostingList{
					DocFrequency: 0,
					Postings:     make(map[int]*Posting),
				}
			}

			if _, exists := idx.Index[token].Postings[docID]; !exists {
				idx.Index[token].Postings[docID] = &Posting{
					DocID:     docID,
					Frequency: 0,
					Positions: make([]int, 0),
				}
				idx.Index[token].DocFrequency++
			}

			posting := idx.Index[token].Postings[docID]
			posting.Frequency++
			posting.Positions = append(posting.Positions, pos)
		}
	}

	return idx
}

// NewInvertedIndex membuat instance baru dari InvertedIndex
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index: make(map[string]*PostingList),
	}
}