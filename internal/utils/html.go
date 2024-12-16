// Package utils berisi fungsi-fungsi utilitas
package utils

import (
	"regexp"
	"strings"
)

// HighlightText menyorot teks query dalam konten
func HighlightText(text, query string) string {
	if query == "" {
		return text
	}

	// Buat pattern untuk highlight
	queryTokens := strings.Fields(query)
	highlighted := text

	for _, token := range queryTokens {
		if len(token) < 2 {
			continue
		}
		pattern := `(?i)\b[\wа-я]*` + regexp.QuoteMeta(token) + `[\wа-я]*\b`
		re := regexp.MustCompile(pattern)
		highlighted = re.ReplaceAllString(highlighted, `<em>$0</em>`)
	}

	return highlighted
}

// GetFaviconPath mengembalikan path favicon berdasarkan URL
func GetFaviconPath(url string) string {
	switch {
	case strings.HasPrefix(url, "https://artikel.rumah123.com/"):
		return "/static/rumah123.png"
	case strings.HasPrefix(url, "https://propertiterkini.com/"):
		return "/static/propertiterkini.png"
	case strings.HasPrefix(url, "https://propertyandthecity.com/"):
		return "/static/propertyandthecity.png"
	default:
		return "/static/favicon.svg"
	}
}