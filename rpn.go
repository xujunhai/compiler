package main

import (
	"lex"
	"fmt"
	"strconv"
)

//逆波兰式 中缀表达式转后缀表达式求值
/**
思路：
遍历中缀表达式中的数字和符号
对于数字：直接输出
对于符号：
	左括号：进栈
	符号：与栈顶符号进行优先级比较
		若当前运算符>栈顶运算符：进栈
		若当前运算符<=栈顶运算符,出栈运算符直到当前运算符>栈顶运算符或者栈空,进栈当前运算符
	右括号：将栈顶符号弹出并输出，直到匹配左括号
遍历结束：将栈中的所有符号弹出并输出
 */

/**
第二种转成二叉树，后序遍历：


 */


func main(){
	s := "(1+3)*(4+(5-3))"
	fmt.Println(s)
	s = transForm(s)
	fmt.Println(s)
	fmt.Println(calculate(s))
}

const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

func precedence(op string) int {
	switch op {
	case "||":
		return 1
	case "&&":
		return 2
	case "==", "!=", "<", "<=", ">", ">=":
		return 3
	case "+", "-", "|", "^":
		return 4
	case "*", "/", "%", "<<", ">>", "&", "&!":
		return 5
	}
	return LowestPrec
}

func IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '/' || r == '*'
}

//求值
func calculate(expr string) int {
	var stack = lex.NewStack() //操作符栈
	for _,v := range expr {
		if IsOperator(v) {
			r := stack.Pop().(int)
			l := stack.Pop().(int)
			var res int
			switch v {
			case '+':
				res = r+l
			case '-':
				res = l-r
			case '*':
				res = l*r
			case '/':
				res = l/r
			}
			stack.Push(res)
		} else {
			ts,_ := strconv.Atoi(string(v))
			stack.Push(ts)
		}
	}
	return stack.Pop().(int)
}

func transForm(exp string) string {

	var stack = lex.NewStack() //操作符栈
	var buf []rune
	for _,v := range exp {
		if lex.IsDigit(v) {
			buf = append(buf,v)
		} else if IsOperator(v) {
			if stack.Empty() { //栈空push
				stack.Push(v)
			} else {
				p := stack.Peek() //比较栈顶操作符优先级
				for !stack.Empty() && precedence(string(v)) <= precedence(string(p.(rune))) { //若当前运算符<=栈顶运算符,出栈运算符直到当前运算符>栈顶运算符或者栈空
					buf = append(buf,stack.Pop().(rune))
				}
				stack.Push(v) //大于栈顶运算符 直接入栈
			}

		}

		if v == '(' { //uint8 == (
			stack.Push(v)
		}

		if v == ')' { //弹出所有操作符直到（
			for {
				p := stack.Pop().(rune)
				if p == '(' {
					break
				}
				buf = append(buf,p)
			}
		}
	}

	for !stack.Empty() { //栈不为空全部弹出
		buf = append(buf,stack.Pop().(rune))
	}

	return string(buf)
}
