// Package utils berisi fungsi-fungsi utilitas
package utils

import (
	"regexp"
	"strings"
)

var (
	// Daftar teks yang tidak diinginkan dalam konten
	unwantedTexts = []string{
		"Baca juga:", "Baca Juga:",
		"Simak breaking news", "Google News",
		"Terus ikuti", "Lebih banyak informasi",
		"Follow", "Instagram", "Twitter", "Facebook",
		"Bagikan:", "Share:", "Read more",
	}

	// Kompilasi regex untuk optimasi
	urlPattern      = regexp.MustCompile(`https?://\S+`)
	emailPattern    = regexp.MustCompile(`\S+@\S+\.\S+`)
	socialPattern   = regexp.MustCompile(`@\S+`)
	specialPattern  = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	numberPattern   = regexp.MustCompile(`\s+\d+\s+`)
	multiSpacePattern = regexp.MustCompile(`\s+`)
)

// CleanContent membersihkan konten artikel dari teks yang tidak diinginkan
func CleanContent(content string) string {
	// Hapus teks yang tidak diinginkan
	for _, text := range unwantedTexts {
		content = strings.ReplaceAll(content, text, "")
	}

	// Hapus URL, email, dan username sosial media
	content = urlPattern.ReplaceAllString(content, "")
	content = emailPattern.ReplaceAllString(content, "")
	content = socialPattern.ReplaceAllString(content, "")

	// Hapus karakter khusus dan angka yang berdiri sendiri
	content = specialPattern.ReplaceAllString(content, " ")
	content = numberPattern.ReplaceAllString(content, " ")

	// Normalisasi spasi dan trim
	content = multiSpacePattern.ReplaceAllString(content, " ")
	return strings.TrimSpace(content)
}

// GetContentPreview menghasilkan preview konten dengan highlight query
func GetContentPreview(content, query string, maxLength int) string {
	cleanedContent := CleanContent(content)
	
	if len(cleanedContent) <= maxLength {
		return cleanedContent
	}

	// Cari posisi query dalam konten
	pos := strings.Index(strings.ToLower(cleanedContent), strings.ToLower(query))
	if pos == -1 {
		return cleanedContent[:maxLength] + "..."
	}

	// Hitung start dan end position untuk preview
	start := pos - 60
	if start < 0 {
		start = 0
	}

	end := start + maxLength
	if end > len(cleanedContent) {
		end = len(cleanedContent)
	}

	result := cleanedContent[start:end]
	if start > 0 {
		result = "..." + result
	}
	if end < len(cleanedContent) {
		result = result + "..."
	}

	return result
}