package token

/**
Subset we're going to use to start with our lexer.

let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
*/

// @todo Use int or byte type instead.
type TokenType string // int or byte is more performant but for ease of print, string will do for now.

type Token struct {
	Type    TokenType // IDENTIFIER, OR NUMBER etc.
	Literal string    // literal (actual) value of the input, e.g. "{" or "("
}

const (
	ILLEGAL = "ILLEGAL" // catchall, token we don't know.
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, (var names)
	INT   = "INT"   // 1234..

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
