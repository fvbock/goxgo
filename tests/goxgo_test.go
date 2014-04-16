package goxgo

import (
	"testing"

	gxg "github.com/fvbock/goxgo"
)

var (
	dsn              gxg.DSN
	tokenizeResponse gxg.TokenizeResponse
	stemResponse     gxg.StemResponse
)

func init() {
	dsn = gxg.DSN{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     4563,
	}

}

func TestGoXGoTokenize(t *testing.T) {
	tokenizePayload := gxg.TokenizeRequest{
		Target: &gxg.CallTarget{
			Services: []string{"NLTK/tokenize"},
			Version:  "0.1"},
		Body:   "Give me a tokenized version of this body of text. Testing trying embodiment embodied",
		Locale: "en",
	}

	gxg.Call(&dsn, &tokenizePayload, &tokenizeResponse)
	t.Logf("tokenizeResponse:\n%v\n", tokenizeResponse)
}

func TestGoXGoStem(t *testing.T) {
	stemPayload := gxg.StemRequest{
		Target: &gxg.CallTarget{
			Services: []string{"NLTK/stem"},
			Version:  "0.1"},
		Words:  tokenizeResponse.Tokens,
		Locale: tokenizeResponse.Locale,
	}

	gxg.Call(&dsn, &stemPayload, &stemResponse)
	t.Logf("stemResponse:\n%v\n", stemResponse)
}
