package lex

type TokenStream struct {
	tokens      []TokenExpr
	index       int
	tokenLength int
}

func NewTokenStream(tokens []TokenExpr) *TokenStream {

	var ret *TokenStream

	ret = new(TokenStream)
	ret.tokens = tokens
	ret.tokenLength = len(tokens)
	return ret
}

// Peek lookahead预测下一个运算符
func (t *TokenStream) Peek() TokenExpr {
	if t.index < t.tokenLength {
		return t.tokens[t.index]
	}
	return TokenExpr{Kind: EOF}
}

func (t *TokenStream) Rewind() {
	t.index -= 1
}

func (t *TokenStream) Next() TokenExpr {
	var token TokenExpr
	token = t.tokens[t.index]
	t.index += 1
	return token
}

func (t TokenStream) HasNext() bool {
	return t.index < t.tokenLength
}
