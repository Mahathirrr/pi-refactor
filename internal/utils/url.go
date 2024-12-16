// Package utils berisi fungsi-fungsi utilitas
package utils

import "strings"

// TrimURLPath memformat URL untuk tampilan yang lebih baik
func TrimURLPath(url string) string {
	// Hapus protokol
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")

	// Ambil path setelah domain
	parts := strings.SplitN(url, "/", 2)
	if len(parts) > 1 {
		path := parts[1]
		if len(path) > 50 {
			return path[:47] + "..."
		}
		return path
	}
	return ""
}

// GetDomainFromURL mengekstrak domain dari URL
func GetDomainFromURL(url string) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	parts := strings.Split(url, "/")
	return parts[0]
}