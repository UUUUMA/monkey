package token

const (
	// 制御文字
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 & 　整数リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = ">"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
