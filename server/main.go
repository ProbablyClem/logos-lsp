package main

import (
	"bufio"
	"encoding/json"
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
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			println("Error decoding message: %s", err)
			continue
		}

		handleMessage(state, method, content)
	}
}

func handleMessage(state *analysis.State, method string, msg []byte) {
	println("Received message with method", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(msg, &request); err != nil {
			println("Error unmarshalling initialize request %s", err)
		}
		println("Connected to", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		writer := os.Stdout
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer.Write([]byte(reply))

		println("Sent initialize response")

	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(msg, &notification); err != nil {
			println("Error unmarshalling didOpen notification %s", err)
		}
		println("Opened document", notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)

	}
}
