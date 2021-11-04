package lex

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

// TokenExpr 通用token表示
type TokenExpr struct {
	Kind  Token
	Value interface{}
}

type TokenParser struct {
	lexer *lexerStream
}

func NewTokenParser(src []byte) *TokenParser {
	return &TokenParser{
		lexer: NewLexerStream(string(src)),
	}
}

func (t *TokenParser) Parse() ([]TokenExpr,error)  {
	var tokens []TokenExpr

	//初始化
	t.lexer.next()
	//跳过空格等字符
	for {
		tok := t.scan()
		if tok.Kind == EOF {
			break
		}
		tokens = append(tokens,tok)
	}
	return tokens,nil
}

func (t TokenParser) scan() TokenExpr {
	t.skipWhiteSpace()

	switch ch := t.lexer.ch; {
	case isLetter(ch):
		//t := t.scanIdentifier()
		//是否是GROUP,状态转义

		//浮点数||整数 今日数字状态
	case isDecimal(ch) || ch == '.' && isDecimal(t.lexer.peek()):
		return t.scanNumber()
	default:
		t.lexer.next() // always make progress
		switch ch {
		case -1:
			return TokenExpr{Kind: EOF}
		case '+':
			return TokenExpr{Kind: ADD,Value: ch}
		case '-':
			return TokenExpr{Kind: SUB,Value: ch}
		case '*':
			return TokenExpr{Kind: MUL,Value: ch}
		case '/':
			return TokenExpr{Kind: QUO,Value: ch}
		case '(':
			return TokenExpr{Kind: LPAREN}
		case ')':
			return TokenExpr{Kind: RPAREN}
		}
	}
	return TokenExpr{Kind: EOF}
}

func (t TokenParser) skipWhiteSpace()  {
	for t.lexer.ch == ' ' || t.lexer.ch == '\t' || t.lexer.ch == '\n' || t.lexer.ch == '\r' {
		t.lexer.next()
	}
}

func (t TokenParser) scanNumber() TokenExpr {
	offset := t.lexer.offset //包含当前位置

	for isDecimal(t.lexer.ch) { //当前依然是数字
		t.lexer.next() //移动
	}

	str := string(t.lexer.src[offset:t.lexer.offset])
	i,_ := strconv.Atoi(str)
	return TokenExpr{Kind: INT,Value: i}
}

//标识符
func (t *TokenParser) scanIdentifier() TokenExpr {
	for t.lexer.canRead() {

	}
	return TokenExpr{}
}

//是否是字符
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

// IsDigit 是否是数字
func IsDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}
func lower(ch rune) rune     { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isHex(ch rune) bool     { return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f' }

//逻辑运算符 比较运算

//词法解析器

//自定义函数+变量()

//语法解析器
