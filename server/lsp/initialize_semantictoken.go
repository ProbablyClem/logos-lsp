package lsp

type SemanticTokensOptions struct {
	Legend *SemanticTokensLegend `json:"legend"`
	Full   bool                  `json:"full"`
}

type SemanticTokensLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}

const (
	TokenTypesFunction = 0 // The number is the index of the token type in the legend (defined by the spec)
)

func NewSemanticTokensOptions() SemanticTokensOptions {
	return SemanticTokensOptions{
		Legend: &SemanticTokensLegend{
			TokenTypes:     []string{"function"},
			TokenModifiers: []string{},
		},
		Full: true,
	}
}
