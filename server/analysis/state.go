package analysis

import (
	"fmt"
	"logos-lsp/lsp"
)

type State struct {
	// Map file name to content
	Documents map[string]string
}

func NewState() *State {
	return &State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(uri string, position lsp.Position) lsp.HoverResult {
	return lsp.HoverResult{
		Contents: fmt.Sprintf("You are at line %d, character %d", position.Line, position.Character),
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
