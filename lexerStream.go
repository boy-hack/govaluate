package govaluate

type lexerStream struct {
	source         []rune
	position       int
	length         int
	functionStatus string // 是否为函数参数状态 idle pre(匹配'(' ) pre2(匹配到双引号) pre3(到这个状态，所有除了"(没有转译的) 都包含进来)
}

func newLexerStream(source string) *lexerStream {

	var ret *lexerStream
	var runes []rune

	for _, character := range source {
		runes = append(runes, character)
	}

	ret = new(lexerStream)
	ret.source = runes
	ret.length = len(runes)
	ret.functionStatus = "idle"
	return ret
}

func (this *lexerStream) readCharacter() rune {

	var character rune

	character = this.source[this.position]
	this.position += 1
	return character
}

func (this *lexerStream) rewind(amount int) {
	this.position -= amount
}

func (this lexerStream) canRead() bool {
	return this.position < this.length
}
