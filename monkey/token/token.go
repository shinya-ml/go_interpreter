package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子, リテラル
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1, 33, ...

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// keywordsにidentが対応するものがなければ識別子として解釈する
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
