package scanner

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"token"
	"utils"
)

var eof = rune(0)

// Scanner type
type Scanner struct {
	r *bufio.Reader
}

// NewScanner new instance
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()

	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok token.Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if utils.IsWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if utils.IsLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return token.EOF, ""
	case '*':
		return token.ASTERISK, string(ch)
	case ',':
		return token.COMMA, string(ch)
	}

	return token.ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		ch := s.read()
		if ch == eof {
			break
		} else if !utils.IsWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return token.WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok token.Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof || utils.IsWhitespace(ch) || ch == ',' {
			s.unread()
			break
		} else if !utils.IsLetter(ch) && utils.IsDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return token.SELECT, buf.String()
	case "FROM":
		return token.FROM, buf.String()
	}

	// Otherwise return as a regular identifier.
	return token.IDENT, buf.String()
}
