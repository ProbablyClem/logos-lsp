package main

import (
	"bufio"
	"encoding/json"
	"io"
	"logos-lsp/analysis"
	"logos-lsp/lsp"
	"logos-lsp/rpc"
	"os"
)

func main() {
	println("Starting Logos Server")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			println("Error decoding message: %s", err)
			continue
		}

		handleMessage(writer, state, method, content)
	}
}

func handleMessage(writer io.Writer, state *analysis.State, method string, msg []byte) {
	println("Received message with method", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(msg, &request); err != nil {
			println("Error unmarshalling initialize request %s", err)
		}
		println("Connected to", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		println("Sent initialize response")

	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(msg, &notification); err != nil {
			println("Error unmarshalling didOpen notification %s", err)
		}
		println("Opened document", notification.Params.TextDocument.URI)
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)

	case "textDocument/didChange":
		var notification lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(msg, &notification); err != nil {
			println("Error unmarshalling didChange notification %s", err)
		}
		println("Changed document", notification.Params.TextDocument.URI)
		for _, change := range notification.Params.ContentChanges {
			state.UpdateDocument(notification.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(msg, &request); err != nil {
			println("Error unmarshalling hover request %s", err)
		}

		result := state.Hover(
			request.Params.TextDocumentPositionParams.TextDocument.URI,
			request.Params.TextDocumentPositionParams.Position,
		)

		response := lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &request.ID,
			},
			Result: result,
		}

		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
