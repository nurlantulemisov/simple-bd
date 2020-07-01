package scanner

import (
	"fmt"
	"io"
	"statement"
	"token"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok token.Token // last read token
		lit string      // last read literal
		n   int         // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (statement.Statement, error) {
	tok, _ := p.scanIgnoreWhitespace()

	var stmt statement.Statement
	var err error

	switch tok {
	case token.SELECT:
		stmt, err = p.parceSelect()
	case token.INSERT:
		stmt, err = p.parceInsert()
	}

	if err != nil {
		return nil, err
	}

	// Next we should see the "FROM" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != token.FROM {
		return nil, fmt.Errorf("found %q, expected FROM", lit)
	}

	// Finally we should read the table name.
	tokTbl, litTbl := p.scanIgnoreWhitespace()
	if tokTbl != token.IDENT {
		return nil, fmt.Errorf("found %q, expected table name", litTbl)
	}
	stmt.SetTable(litTbl)

	// Return the successfully parsed statement.
	return stmt, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok token.Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok token.Token, lit string) {
	tok, lit = p.scan()

	if tok == token.WS {
		tok, lit = p.scan()
	}

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) parceSelect() (*statement.Select, error) {
	stmt := &statement.Select{}

	// Next we should loop over all our comma-delimited fields.
	for {
		// Read a field.
		tok, lit := p.scanIgnoreWhitespace()
		if tok != token.IDENT && tok != token.ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", lit)
		}

		stmt.Fields = append(stmt.Fields, lit)

		// If the next token is not a comma then break the loop.
		if tok, _ := p.scanIgnoreWhitespace(); tok != token.COMMA {
			p.unscan()
			break
		}
	}
	return stmt, nil
}

func (p *Parser) parceInsert() (*statement.Insert, error) {
	stmt := &statement.Insert{}
	stmt.Fields = make(map[string]string)
	// Next we should loop over all our comma-delimited fields.
	for {
		// Read a field.
		tokField, litField := p.scanIgnoreWhitespace()
		if tokField != token.IDENT {
			return nil, fmt.Errorf("found %q, expected field", litField)
		}

		tokEqual, litEqual := p.scanIgnoreWhitespace()
		if tokEqual != token.EQUALL {
			return nil, fmt.Errorf("found %q, expected =", litEqual)
		}

		tokValue, litValue := p.scanIgnoreWhitespace()
		if tokValue != token.IDENT {
			return nil, fmt.Errorf("found %q, expected field", litValue)
		}

		stmt.Fields[litField] = litValue

		// If the next token is not a comma then break the loop.
		if tok, _ := p.scanIgnoreWhitespace(); tok != token.COMMA {
			p.unscan()
			break
		}
	}
	return stmt, nil
}
