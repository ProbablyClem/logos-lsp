package rpc_test

import (
	"logos-lsp/rpc"
	"testing"
)

type EncodingExampleStruct struct {
	Hello string
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 17\r\n\r\n{\"Hello\":\"world\"}"
	actual := rpc.EncodeMessage(EncodingExampleStruct{Hello: "world"})

	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incoming := "Content-Length: 17\r\n\r\n{\"Method\":\"test\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	length := len(content)
	expectedLength := 17
	expectedMethod := "test"

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if length != expectedLength {
		t.Errorf("Expected length %d, got %d", expectedLength, length)
	}

	if method != expectedMethod {
		t.Errorf("Expected method %s, got %s", expectedMethod, method)
	}
}
