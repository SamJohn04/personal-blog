package services

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

func MarkdownToHTML(markdownContent string) (string, error) {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(markdownContent), &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

var sanitizer = bluemonday.UGCPolicy()

func SanitizeHTML(htmlContent string) string {
	return sanitizer.Sanitize(htmlContent)
}
