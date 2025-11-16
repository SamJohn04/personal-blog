package services

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func MarkdownToHTML(markdown_content string) (string, error) {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(markdown_content), &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
