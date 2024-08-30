package analysis_test

import (
	"logos-lsp/analysis"
	"logos-lsp/lsp"
	"testing"
)

func TestFindBibleQuotes(t *testing.T) {
	text := `Here are some Bible quotes: 
	Jn 3 16, Jean 3:16, and Jean 3-16. 
	Some other formats might include John 3:16 or Jn 3-16.
	Don't forget about Gen 1:1 or Rom 8 28.`

	uri := "file:///path/to/file.txt"
	bibleQuotes := analysis.FindBibleQuotesWithPosition(uri, text)

	if len(bibleQuotes) != 7 {
		t.Errorf("Expected 6 Bible quotes, got %d", len(bibleQuotes))
	}
	// Check the first quote
	if bibleQuotes[0].Book != "Jean" {
		t.Errorf("Expected book to be Jn, got %s", bibleQuotes[0].Book)
	}
	if bibleQuotes[0].Chapter != 3 {
		t.Errorf("Expected chapter to be 3, got %d", bibleQuotes[0].Chapter)
	}
	if bibleQuotes[0].Verse != 16 {
		t.Errorf("Expected verse to be 16, got %d", bibleQuotes[0].Verse)
	}
	if bibleQuotes[0].Uri != uri {
		t.Errorf("Expected URI to be %s, got %s", uri, bibleQuotes[0].Uri)
	}
	if bibleQuotes[0].Range.Start.Line != 2 {
		t.Errorf("Expected line to be 1, got %d", bibleQuotes[0].Range.Start.Line)
	}
	if bibleQuotes[0].Range.Start.Character != 2 {
		t.Errorf("Expected character to be 2, got %d", bibleQuotes[0].Range.Start.Character)
	}
	if bibleQuotes[0].Range.End.Line != 2 {
		t.Errorf("Expected line to be 2, got %d", bibleQuotes[0].Range.End.Line)
	}
	if bibleQuotes[0].Range.End.Character != 9 {
		t.Errorf("Expected character to be 9, got %d", bibleQuotes[0].Range.End.Character)
	}

	// Check the last quote
	if bibleQuotes[6].Book != "Romains" {
		t.Errorf("Expected book to be Romains, got %s", bibleQuotes[6].Book)
	}
	if bibleQuotes[6].Chapter != 8 {
		t.Errorf("Expected chapter to be 8, got %d", bibleQuotes[6].Chapter)
	}
	if bibleQuotes[6].Verse != 28 {
		t.Errorf("Expected verse to be 28, got %d", bibleQuotes[6].Verse)
	}
	if bibleQuotes[6].Uri != uri {
		t.Errorf("Expected URI to be %s, got %s", uri, bibleQuotes[6].Uri)
	}
	if bibleQuotes[6].Range.Start.Line != 4 {
		t.Errorf("Expected line to be 4, got %d", bibleQuotes[6].Range.Start.Line)
	}
	if bibleQuotes[6].Range.Start.Character != 32 {
		t.Errorf("Expected character to be 32, got %d", bibleQuotes[6].Range.Start.Character)
	}
	if bibleQuotes[6].Range.End.Line != 4 {
		t.Errorf("Expected line to be 4, got %d", bibleQuotes[6].Range.End.Line)
	}
	if bibleQuotes[6].Range.End.Character != 40 {
		t.Errorf("Expected character to be 37, got %d", bibleQuotes[6].Range.End.Character)
	}
}

func TestFindRangeBibleQuotes(t *testing.T) {
	text := `qdqzdzqd Mt 20 1-10 qzdqdqzd`

	uri := "file:///path/to/file.txt"
	bibleQuotes := analysis.FindBibleQuotesWithPosition(uri, text)

	// Check the first quote
	if bibleQuotes[0].Book != "Matthieu" {
		t.Errorf("Expected book to be Matthieu, got %s", bibleQuotes[0].Book)
	}
	if bibleQuotes[0].Chapter != 20 {
		t.Errorf("Expected chapter to be 20, got %d", bibleQuotes[0].Chapter)
	}
	if bibleQuotes[0].Verse != 1 {
		t.Errorf("Expected verse to be 1, got %d", bibleQuotes[0].Verse)
	}
	if bibleQuotes[0].Range.Start.Line != 0 {
		t.Errorf("Expected line to be 0, got %d", bibleQuotes[0].Range.Start.Line)
	}
	if bibleQuotes[0].Range.Start.Character != 9 {
		t.Errorf("Expected character to be 9, got %d", bibleQuotes[0].Range.Start.Character)
	}
	if bibleQuotes[0].Range.End.Line != 0 {
		t.Errorf("Expected line to be 0, got %d", bibleQuotes[0].Range.End.Line)
	}
	if bibleQuotes[0].Range.End.Character != 19 {
		t.Errorf("Expected character to be 19, got %d", bibleQuotes[0].Range.End.Character)
	}
}

func TestQuoteIsInRange(t *testing.T) {
	quote := analysis.Quote{
		Range: lsp.Range{
			Start: lsp.Position{Line: 2, Character: 1},
			End:   lsp.Position{Line: 2, Character: 5},
		},
	}

	// Test when the position is before the quote
	if quote.IsInRange(1, 4) {
		t.Errorf("Expected position to not be in range")
	}

	// Test when the position is within the quote
	if !quote.IsInRange(2, 3) {
		t.Errorf("Expected position to be in range")
	}

	// Test when the position is after the quote
	if quote.IsInRange(2, 6) {
		t.Errorf("Expected position to be in range")
	}
}
