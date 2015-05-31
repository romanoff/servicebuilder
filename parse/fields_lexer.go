package parse

import (
	"strings"
)

func (self *Lexer) scanFieldIdentifier() Lexeme {
	lexeme := self.scanIdentifier()
	switch strings.ToLower(lexeme.Value) {
	case "string":
		lexeme.Token = STRING
	case "int":
		lexeme.Token = INT
	case "double":
		lexeme.Token = DOUBLE
	case "date":
		lexeme.Token = DATE
	case "datetime":
		lexeme.Token = DATETIME
	}
	return lexeme
}

func (self *Lexer) scanFields() {
	self.nextFunc = nil
	originalParenDepth := self.parenDepth
ScanLoop:
	for {
		if self.nextFunc != nil {
			self.nextFunc()
		}
		ch := self.read()
		if isWhitespace(ch) {
			self.unread()
			self.bus <- self.scanWhitespace()
			continue
		} else if isLetter(ch) {
			self.unread()
			self.bus <- self.scanFieldIdentifier()
			continue
		}

		switch ch {
		case eof:
			self.unread()
			break ScanLoop
		case '{':
			self.parenDepth++
			self.bus <- Lexeme{LEFTBRACE, string(ch), *self.Position}
			continue
		case ':':
			self.bus <- Lexeme{COLON, string(ch), *self.Position}
			continue
		case '}':
			self.parenDepth--
			self.bus <- Lexeme{RIGHTBRACE, string(ch), *self.Position}
			if self.parenDepth == originalParenDepth {
				break ScanLoop
			}
			continue
		}
	}
}