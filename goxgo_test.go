package goxgo

import (
	"testing"
)

var (
	dsn              DSN
	tokenizeResponse TokenizeResponse
	stemResponse     StemResponse
)

func init() {
	dsn = DSN{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     4563,
	}

}

func TestGoXGoTokenize(t *testing.T) {
	tokenizePayload := TokenizeRequest{
		Target: &CallTarget{
			Services: []string{"NLTK/tokenize"},
			Version:  "0.1"},
		Body:   "Give me a tokenized version of this body of text. Testing trying embodiment embodied",
		Locale: "en",
	}

	Call(&dsn, &tokenizePayload, &tokenizeResponse)
	t.Logf("tokenizeResponse:\n%v\n", tokenizeResponse)
}

func TestGoXGoStem(t *testing.T) {
	stemPayload := StemRequest{
		Target: &CallTarget{
			Services: []string{"NLTK/stem"},
			Version:  "0.1"},
		Words:  tokenizeResponse.Tokens,
		Locale: tokenizeResponse.Locale,
	}

	Call(&dsn, &stemPayload, &stemResponse)
	t.Logf("stemResponse:\n%v\n", stemResponse)
}
