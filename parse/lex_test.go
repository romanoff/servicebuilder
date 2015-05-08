package parse

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestLexerScanning(t *testing.T) {
	cases := []struct {
		Input  string
		Output []Token
	}{
		{
			"simple.sb",
			[]Token{MODEL, WS, IDENT, WS, LEFTBRACE, WS, FIELDS, WS, LEFTBRACE, WS, IDENT, COLON, WS, STRING, WS, RIGHTBRACE, WS, RIGHTBRACE, WS, EOF},
		},
	}
	for _, tc := range cases {
		content, err := ioutil.ReadFile(filepath.Join("test-fixtures", tc.Input))
		if err != nil {
			t.Errorf("Expected to read content from %v fixture, but got %v", tc.Input, err)
		}
		lexer := NewLexer(tc.Input, bytes.NewBuffer([]byte(content)))
		scanner := lexer.Scan()
		i := 0
		for {
			lexeme := <-scanner
			if lexeme.Token != tc.Output[i] {
				t.Errorf("Expected to get %v token, but got %v", tc.Output[i], lexeme.Token)
			}
			i++
			if lexeme.Token == EOF {
				break
			}
		}

	}
}
