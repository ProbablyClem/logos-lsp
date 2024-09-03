package lsp

type CodeActionRequest struct {
	Request
	Params CodeActionParams `json:"params"`
}

type CodeActionParams struct {
	TextDocumentIdentifier TextDocumentIdentifier `json:"textDocument"`
	Range                  Range                  `json:"range"`
}

type CodeActionKind string

const (
	RefactorInline CodeActionKind = "refactor.inline"
)

type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeAction struct {
	Title       string         `json:"title"`
	Kind        CodeActionKind `json:"kind"`
	Command     Command        `json:"command"`
	IsPreferred bool           `json:"isPreferred"`
	Diagnostics []Diagnostic   `json:"diagnostics"`
}

type Command struct {
	Command string `json:"command"`
	Title   string `json:"title"`
}
