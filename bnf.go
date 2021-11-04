import (
	"lex"
	"fmt"
)

/**
词法分析：
	数字：[0-9]+
	左括号: (
	右括号: )
	操作符: + - * /
上下文无关文法：

第一种简单文法BNF转成AST树:
	addExpr: mulExpr | mulExpr op1 addExpr
	op1: +|-
	mulExpr: num | num op2 mulExpr
	op2: *|/
	num: [0-9] | (addExpr)

第二种pratt parser


*/

func main()  {
	src := "(2+3)*(5-4)+5"
	tp := lex.NewTokenParser([]byte(src))
	tokens,_ := tp.Parse()
	fmt.Println(tokens)

	//语法解析
	s := lex.NewTokenStream(tokens)
	node := addExpr(s)
	//深度优先visitor遍历
	walk(node)
	rs := eval(node,"")
	fmt.Println(rs)
}



// Binary ASTNode A BinaryExpr node represents a binary expression.
type Binary struct {
	X     IExpr        // left operand
	Op    lex.Token // operator
	Y     IExpr        // right operand
}

func (b Binary) Type() lex.Token {
	return b.Op
}

type IExpr interface {
	Type() lex.Token
}

type BasicLit struct {
	Kind lex.Token
	Value interface{}
}

func (b BasicLit) Type() lex.Token {
	return b.Kind
}


func walk(expr IExpr) {

}

// -------------------------------
//递归求值
func eval(expr IExpr,step string) int  {
	switch expr.(type) {
	case Binary:
		lhs := eval(expr.(Binary).X,"")
		rhs := eval(expr.(Binary).Y,"")
		switch expr.(Binary).Op {
		case lex.SUB:
			fmt.Printf("%d-%d\n",lhs,rhs)
			return lhs - rhs
		case lex.ADD:
			fmt.Printf("%d+%d\n",lhs,rhs)
			return lhs + rhs
		case lex.MUL:
			fmt.Printf("%d*%d\n",lhs,rhs)
			return lhs * rhs
		case lex.QUO:
			fmt.Printf("%d/%d\n",lhs,rhs)
			return lhs / rhs
		}
	case BasicLit:
		switch expr.(BasicLit).Kind {
		case lex.INT:
			return expr.(BasicLit).Value.(int)
		}
	}
	return 0
}

//返回一个AST树，其根是一个“+”或“-”二进制运算符
func addExpr(stream *lex.TokenStream) IExpr {
	lhs := mulExpr(stream)
	node := lhs

	t := stream.Peek()	//是否匹配第二种文法
	if lhs.Type() != lex.EOF && t.Kind != lex.EOF {
		if t.Kind == lex.ADD || t.Kind == lex.SUB { //op2
			stream.Next()
			rhs := addExpr(stream)
			node = Binary{X: lhs,Op: t.Kind,Y: rhs}
		}
	}
	return node
}

//返回一个AST树，其根是一个“*”或“/”二进制运算符
func mulExpr(stream *lex.TokenStream) IExpr {
	lhs := primary(stream) //第一种文法
	node := lhs

	t := stream.Peek()	//是否匹配第二种文法
	if lhs.Type() != lex.EOF && t.Kind != lex.EOF {
		if t.Kind == lex.MUL || t.Kind == lex.QUO { //op2
			stream.Next()
			rhs := mulExpr(stream)
			node = Binary{X: lhs,Op: t.Kind,Y: rhs}
		}
	}
	return node
}

//基础表达式
func primary(stream *lex.TokenStream) IExpr  {
	var node IExpr
	t := stream.Peek()
	if t.Kind.IsLiteral() {
		stream.Next()
		node = BasicLit{Kind: t.Kind,Value: t.Value}
	}

	//左阔号
	if t.Kind == lex.LPAREN {
		stream.Next() //消耗token
		node = addExpr(stream) //递归得出(X)
		if node.Type() != lex.EOF {
			t := stream.Peek()
			if t.Kind != lex.RPAREN {
				panic("expect )")
			}
			stream.Next()
		}
	}
	return node
}
