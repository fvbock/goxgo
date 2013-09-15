package goxgo

import (
	gxg "github.com/fvbock/goxgo"
	"testing"
)

func TestGoXGo(t *testing.T) {
	dsn := gxg.DSN{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     4563,
	}

	tokenizePayload := gxg.TokenizeRequest{
		Target: &gxg.CallTarget{Services: []string{"NLTK/tokenize"}, Version: "0.1"},
		Body:   "Give me a tokenized version of this unoptimzed body of text pls. Once successfully done we will try to stem the words too. Testing trying embodiment embodied",
		Locale: "en",
	}

	var tokenizeResponse gxg.TokenizeResponse
	gxg.Call(&dsn, &tokenizePayload, &tokenizeResponse)
	t.Log("tokenizeResponse", tokenizeResponse)

	stemPayload := gxg.StemRequest{
		Target: &gxg.CallTarget{Services: []string{"NLTK/stem"}, Version: "0.1"},
		Words:  tokenizeResponse.Tokens,
		Locale: tokenizeResponse.Locale,
	}

	var stemResponse gxg.StemResponse
	gxg.Call(&dsn, &stemPayload, &stemResponse)
	t.Log("stemResponse", stemResponse)
}
