package lsp

type SemanticTokensRequest struct {
	Request
	Params SemanticTokensParams `json:"params"`
}

type SemanticTokensParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type SemanticTokensResponse struct {
	Response
	Result SemanticTokensResult `json:"result"`
}

type SemanticTokensResult struct {
	Data []int `json:"data"`
}

type SemanticToken struct {
	TokenType int
	Line      int
	StartChar int
	Length    int
}

func EncodeSemanticTokens(tokens []SemanticToken) []int {
	previous_token := SemanticToken{
		Line:      0,
		StartChar: 0,
	}

	data := []int{}
	for _, token := range tokens {

		delta_line := token.Line - previous_token.Line
		delta_start_char := token.StartChar

		// If the line is the same, calculate the delta start character
		if delta_line == 0 {
			delta_start_char = token.StartChar - previous_token.StartChar
		}

		data = append(data, delta_line)
		data = append(data, delta_start_char)
		data = append(data, token.Length)
		data = append(data, token.TokenType)
		data = append(data, 0) // token modifier (not supported)
		previous_token = token
	}
	return data
}
