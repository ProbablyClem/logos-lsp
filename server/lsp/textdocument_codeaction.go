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
	IsPreferred bool           `json:"isPreferred"`
	Diagnostics []Diagnostic   `json:"diagnostics"`
	Edit        WorkspaceEdit  `json:"edit"`
}

type WorkspaceEdit struct {
	Changes map[string][]TextEdit `json:"changes"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}
