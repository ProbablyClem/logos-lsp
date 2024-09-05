package utils_test

import (
	"logos-lsp/utils"
	"testing"
)

func TestToMarkdownQuote(t *testing.T) {
	t.Run("Single line", func(t *testing.T) {
		text := "Hello, World!"
		want := "> Hello, World!"
		got := utils.ToMarkdownQuote(text)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Multiple lines", func(t *testing.T) {
		text := `Hello, World!
How are you?`
		want := "> Hello, World!\n> How are you?"
		got := utils.ToMarkdownQuote(text)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
