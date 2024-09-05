package analysis

import (
	"fmt"
	"log"
	"logos-lsp/bible"
	"logos-lsp/lsp"
	"logos-lsp/utils"
)

type State struct {
	// Map file name to content
	Documents map[string]string
	// Map quotes by file name
	Quotes map[string][]Quote
	Bible  bible.Bible
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
		Quotes:    make(map[string][]Quote),
		Bible:     bible.LoadFromFile(),
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
	s.searchQuotes(uri)

}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
	s.searchQuotes(uri)
}

func (s *State) Hover(uri string, position lsp.Position) lsp.HoverResult {
	println("Hovering at", position.Line, position.Character)
	quotes := s.Quotes[uri]
	for _, quote := range quotes {
		if quote.IsInRange(position.Line, position.Character) {
			content := s.Bible.GetQuoteContent(quote.Book, quote.Chapter, quote.StartVerse, quote.EndVerse)
			versets := ""
			if quote.EndVerse != quote.StartVerse {
				versets = fmt.Sprintf("%d-%d", quote.StartVerse, quote.EndVerse)
			} else {
				versets = fmt.Sprintf("%d", quote.StartVerse)
			}
			return lsp.HoverResult{
				Contents: lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: fmt.Sprintf("### 📖 %s Chapitre %d Verset %s   \n%s", quote.Book, quote.Chapter, versets, content),
				},
			}
		}
	}

	return lsp.HoverResult{
		Contents: lsp.MarkupContent{
			Kind:  lsp.PlainText,
			Value: "No quote found",
		},
	}
}

func (s *State) Definition(uri string, position lsp.Position) lsp.Location {
	return lsp.Location{
		URI: uri,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      position.Line,
				Character: 0,
			},
			End: lsp.Position{
				Line:      position.Line,
				Character: 10,
			},
		},
	}
}

func (s *State) SemanticTokens(uri string) []lsp.SemanticToken {
	quotes := s.Quotes[uri]
	tokens := []lsp.SemanticToken{}
	for _, quote := range quotes {
		tokens = append(tokens, lsp.SemanticToken{
			TokenType: lsp.TokenTypesFunction,
			Line:      quote.Range.Start.Line,
			StartChar: quote.Range.Start.Character,
			Length:    quote.Range.End.Character - quote.Range.Start.Character,
		})
	}
	return tokens
}

func (s *State) CodeAction(uri string, selectedRange lsp.Range) []lsp.CodeAction {
	quotes := s.Quotes[uri]
	codeActions := []lsp.CodeAction{}

	for _, quote := range quotes {
		if !quote.IsInRange(selectedRange.Start.Line, selectedRange.Start.Character) {
			continue
		}

		quoteRange := lsp.Range{
			Start: lsp.Position{
				Line:      quote.Range.Start.Line,
				Character: quote.Range.Start.Character,
			},
			End: lsp.Position{
				Line:      quote.Range.End.Line,
				Character: quote.Range.End.Character,
			},
		}

		content := s.Bible.GetQuoteContent(quote.Book, quote.Chapter, quote.StartVerse, quote.EndVerse)
		reference := ""
		if quote.EndVerse != quote.StartVerse {
			reference = fmt.Sprintf("%s %d : %d-%d", quote.Book, quote.Chapter, quote.StartVerse, quote.EndVerse)
		} else {
			reference = fmt.Sprintf("%s %d : %d", quote.Book, quote.Chapter, quote.StartVerse)
		}

		content = fmt.Sprintf("### 📖 %s\n%s", reference, content)

		quotedContent := utils.ToMarkdownQuote(content)
		quotedContent = quotedContent + "\n"

		textEdit := lsp.TextEdit{
			Range:   quoteRange,
			NewText: quotedContent,
		}

		codeActions = append(codeActions, lsp.CodeAction{
			Title: fmt.Sprintf("Inline %s %d:%d", quote.Book, quote.Chapter, quote.StartVerse),
			Kind:  lsp.RefactorInline,
			Edit: lsp.WorkspaceEdit{
				Changes: map[string][]lsp.TextEdit{
					uri: {textEdit},
				},
			},
			IsPreferred: true,
			Diagnostics: []lsp.Diagnostic{
				{
					Range:    quoteRange,
					Severity: lsp.Hint,
					Message:  "Inline quote",
				},
			},
		})
	}
	return codeActions
}

func (s *State) searchQuotes(uri string) {
	text, ok := s.Documents[uri]
	if !ok {
		return
	}

	quotes := FindBibleQuotesWithPosition(uri, text)
	s.Quotes[uri] = quotes

	for _, quote := range quotes {
		log.Printf("Found quote: %s %d:%d", quote.Book, quote.Chapter, quote.StartVerse)
		log.Printf("Quote start at %d:%d", quote.Range.Start.Line, quote.Range.Start.Character)
		log.Printf("Quote end at %d:%d", quote.Range.End.Line, quote.Range.End.Character)
	}
}
