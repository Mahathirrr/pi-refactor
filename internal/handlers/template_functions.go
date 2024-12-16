// Package handlers berisi handler HTTP dan fungsi template
package handlers

import (
	"html/template"
	"strings"
)

// TemplateFunctions mengembalikan map fungsi-fungsi yang dapat digunakan dalam template
func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		// Membuat slice angka dari start sampai end
		"iterate": func(start, end int) []int {
			result := make([]int, end-start+1)
			for i := start; i <= end; i++ {
				result[i-start] = i
			}
			return result
		},
		// Operasi pengurangan
		"sub": func(a, b int) int {
			return a - b
		},
		// Operasi penambahan
		"add": func(a, b int) int {
			return a + b
		},
		// Mengecek apakah string dimulai dengan prefix tertentu
		"hasPrefix": strings.HasPrefix,
		// Memotong dan memformat path URL
		"trimURLPath": func(url string) string {
			url = strings.TrimPrefix(url, "https://")
			url = strings.TrimPrefix(url, "http://")

			parts := strings.SplitN(url, "/", 2)
			if len(parts) > 1 {
				path := parts[1]
				if len(path) > 50 {
					path = path[:47] + "..."
				}
				return path
			}
			return ""
		},
	}
}