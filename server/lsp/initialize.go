package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResponseResult `json:"result"`
}

type InitializeResponseResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync TextDocumentSyncOptions `json:"textDocumentSync"`
}

type TextDocumentSyncOptions struct {
	OpenClose bool                 `json:"openClose"`
	SyncKind  TextDocumentSyncKind `json:"change"`
}

type TextDocumentSyncKind int

const (
	None        TextDocumentSyncKind = 0
	Full        TextDocumentSyncKind = 1
	Incremental TextDocumentSyncKind = 2
)

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResponseResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: TextDocumentSyncOptions{
					OpenClose: true,
					SyncKind:  Full,
				},
			},
			ServerInfo: ServerInfo{
				Name:    "Logos",
				Version: "0.1",
			},
		},
	}
}
