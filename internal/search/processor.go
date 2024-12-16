// Package search berisi implementasi mesin pencari
package search

import (
	"regexp"
	"strings"
	"sync"

	"github.com/RadhiFadlillah/go-sastrawi"
)

// TextProcessor menangani pemrosesan teks
type TextProcessor struct {
	stemmer     sastrawi.Stemmer
	stopwords   sastrawi.Dictionary
	punctuation *regexp.Regexp
	numbers     *regexp.Regexp
}

var (
	textProcessor *TextProcessor
	procOnce     sync.Once
)

// GetProcessor mengembalikan instance singleton TextProcessor
func GetProcessor() *TextProcessor {
	procOnce.Do(func() {
		textProcessor = NewTextProcessor()
	})
	return textProcessor
}

// NewTextProcessor membuat instance TextProcessor baru
func NewTextProcessor() *TextProcessor {
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)
	stopwords := sastrawi.DefaultStopword()

	return &TextProcessor{
		stemmer:     stemmer,
		stopwords:   stopwords,
		punctuation: regexp.MustCompile(`[^\w\s]`),
		numbers:     regexp.MustCompile(`\b\d+\b`),
	}
}

// ProcessText melakukan pemrosesan teks lengkap
func (tp *TextProcessor) ProcessText(text string) []string {
	cleaned := tp.removePunctuationsAndNumbers(text)
	folded := tp.caseFolding(cleaned)
	tokens := tp.tokenize(folded)
	return tp.processWords(tokens)
}

// Fungsi-fungsi pembantu
func (tp *TextProcessor) removePunctuationsAndNumbers(text string) string {
	text = tp.punctuation.ReplaceAllString(text, " ")
	text = tp.numbers.ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}

func (tp *TextProcessor) caseFolding(text string) string {
	return strings.ToLower(text)
}

func (tp *TextProcessor) tokenize(text string) []string {
	return sastrawi.Tokenize(text)
}

func (tp *TextProcessor) processWords(tokens []string) []string {
	result := make([]string, 0)
	for _, token := range tokens {
		if tp.stopwords.Contains(token) {
			continue
		}
		stemmed := tp.stemmer.Stem(token)
		if stemmed != "" {
			result = append(result, stemmed)
		}
	}
	return result
}