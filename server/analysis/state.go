package analysis

import (
	"fmt"
	"log"
	"logos-lsp/lsp"
)

type State struct {
	// Map file name to content
	Documents map[string]string
	// Map quotes by file name
	Quotes map[string][]Quote
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
		Quotes:    make(map[string][]Quote),
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
			return lsp.HoverResult{
				Contents: fmt.Sprintf("%s Chapitre %d Verset %d", quote.Book, quote.Chapter, quote.Verse),
			}
		}
	}

	return lsp.HoverResult{
		Contents: "No quote found",
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

func (s *State) searchQuotes(uri string) {
	text, ok := s.Documents[uri]
	if !ok {
		return
	}

	quotes := FindBibleQuotesWithPosition(uri, text)
	s.Quotes[uri] = quotes

	for _, quote := range quotes {
		log.Printf("Found quote: %s %d:%d", quote.Book, quote.Chapter, quote.Verse)
		log.Printf("Quote start at %d:%d", quote.Range.Start.Line, quote.Range.Start.Character)
		log.Printf("Quote end at %d:%d", quote.Range.End.Line, quote.Range.End.Character)
	}
}
