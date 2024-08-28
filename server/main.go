package main

import (
	"bufio"
	"log"
	"logos-lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("logos-lsp.log")
	logger.Println("Starting Logos LSP")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		text := scanner.Text()
		handleMessage(logger, text)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Printf("Received message: %s", msg)
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[logos-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
