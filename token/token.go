package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT"
	INT = "INT"
	STRING  = "STRING"

	ASSIGN = "="
	PLUS = "+"
	BANG = "!"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	COMMA = ","   
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"
	
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LockupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
