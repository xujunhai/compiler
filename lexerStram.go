package lex

type lexerStream struct {
	src   []rune
	offset int
	rdOffset 	 int //当前读的位置
	length   int
	ch rune //当前读到的 character
}

func NewLexerStream(source string) *lexerStream {
	var ret *lexerStream
	var runes []rune

	for _, character := range source {
		runes = append(runes, character)
	}

	ret = new(lexerStream)
	ret.src = runes
	ret.length = len(runes)
	return ret
}

//移动读取位置
func (l *lexerStream) next() {
	if l.rdOffset < len(l.src) {
		l.offset = l.rdOffset

		r, w := rune(l.src[l.rdOffset]), 1
		l.ch = r
		l.rdOffset += w
	} else {
		l.offset = len(l.src)
		l.ch = -1
	}
}

//预读下一个字符 没有增加scanner前进
func (l *lexerStream) peek() rune {
	if l.rdOffset < len(l.src) {
		return l.src[l.rdOffset]
	}
	return 0
}

func (l *lexerStream) rewind(amount int) {
	l.rdOffset -= amount
}

func (l lexerStream) canRead() bool {
	return l.rdOffset < l.length
}