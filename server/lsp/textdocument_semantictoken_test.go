package lsp_test

import (
	"logos-lsp/lsp"
	"testing"
)

func TestEncodeSemanticTokens(t *testing.T) {

	tokens := []lsp.SemanticToken{
		{TokenType: 0, Line: 2, StartChar: 5, Length: 3},
		{TokenType: 0, Line: 2, StartChar: 10, Length: 4},
		{TokenType: 0, Line: 5, StartChar: 2, Length: 7},
	}

	data := lsp.EncodeSemanticTokens(tokens)

	expected := []int{
		2, 5, 3, 0, 0,
		0, 5, 4, 0, 0,
		3, 2, 7, 0, 0,
	}

	if len(data) != len(expected) {
		t.Errorf("Expected %d tokens, got %d", len(expected), len(data))
	}

	for i := range data {
		if data[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], data[i])
		}
	}
}
