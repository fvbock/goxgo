// Request and Response structures for - in this first case - python services
// providing NLP/scientific math functions

package goxgo

import (
	"fmt"
)

/*
CallTarget contains the target information for the payload by naming services to
be called and the desired version of that service

Services - a <[]string> of "service name/procedure name" pairs that the payload
should be "piped through". ie:

 	[ "NLTK/stem" ] or
 	[ "NLTK/tokenize", "NLTK/stem" ]

If you list more than one the result from the first gets passed as arguments
into the next. Make sure they are passable - and raise exceptions on the
server-side if they don't.

NB: this pipelining is not yet implemented

Version - a <string> representing the API version
*/
type CallTarget struct {
	Services []string
	Version  string
}

/* TokenizeRequest - request structure
 */
type TokenizeRequest struct {
	Target *CallTarget
	Body   string
	Locale string
}

/* TokenizeResponse - response structure
 */
type TokenizeResponse struct {
	Locale string   `json:"locale"`
	Tokens []string `json:"tokens"`
}

func (s *TokenizeResponse) String() string {
	return fmt.Sprintf("Tokens\nLanguage: %s\n%v\n", s.Locale, s.Tokens)
}

/* StemRequest - request structure
 */
type StemRequest struct {
	Target *CallTarget
	Words  []string
	Locale string
}

/* StemResponse - response structure
 */
type StemResponse struct {
	Locale string   `json:"locale"`
	Words  []string `json:"words"`
}

func (s *StemResponse) String() string {
	return fmt.Sprintf("Stemmed Tokens\nLanguage: %s\n%v\n", s.Locale, s.Words)
}
