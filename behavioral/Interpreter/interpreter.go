/*
场景:
	规则引擎
	当语言的文法较为简单，且执行效率不是关键问题时。
	当问题重复出现，且可以用一种简单的语言来进行表达时。
	当一个语言需要解释执行，并且语言中的句子可以表示为一个抽象语法树的时候，如 XML 文档解释。
 */
package Interpreter

import "strings"

//抽象表达式（Abstract Expression）角色：定义解释器的接口，约定解释器的解释操作，主要包含解释方法 interpret()。
type Expression interface {
	Interpret(variables map[string]Expression) int
}

//终结符表达式（Terminal Expression）角色：是抽象表达式的子类，用来实现文法中与终结符相关的操作，文法中的每一个终结符都有一个具体终结表达式与之相对应。
//加法表达式
type Plus struct {
	leftOperand Expression
	rightOperand Expression
}
//加法解释器将两个表达式的结果相加
func (e *Plus) Interpret(variables map[string]Expression) int {
	// 左操作符解释的结果，加上右表达式解释后的结果
	return e.leftOperand.Interpret(variables) + e.rightOperand.Interpret(variables)
}

//非终结符表达式（Nonterminal Expression）角色：也是抽象表达式的子类，用来实现文法中与非终结符相关的操作，文法中的每条规则都对应于一个非终结符表达式。
// 整数表达式
type Integer struct {
	integer int
}
// 整数表达式直接返回整数值
func (e *Integer) Interpret(variables map[string]Expression) int {
	return e.integer
}

//非终结符表达式（Nonterminal Expression）角色：也是抽象表达式的子类，用来实现文法中与非终结符相关的操作，文法中的每条规则都对应于一个非终结符表达式。
type Variable struct {
	name string
}
// 变量表达式解释器，将变量中的表达式取出来并解放，然后返回解释后的结果
func (e *Variable) Interpret(variables map[string]Expression) int {
	value, found := variables[e.name]
	if !found {
		return 0
	}
	return value.Interpret(variables)
}

//单链表节点
type Node struct {
	value interface{}
	next *Node
}

//使用单链表实现栈结构
type Stack struct {
	top *Node //栈顶元素
	size int //维持栈的大小
}
func (s *Stack) Push(value interface{}) {
	s.top = &Node{
		value: value,
		next: s.top,
	}
	s.size++
}
func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

//最后，定义一个环境（Context）类，它包含解释器需要的数据，完成对终结符表达式的初始化，并定义一个方法 freeRide(String info) 调用表达式对象的解释方法来对被分析的字符串进行解释。其结构图如图 3 所示。
type Evaluator struct {
	syntaxTree Expression
}
// 解释这个表达式, 参数是表达式中的参数对应的值; 这里调用加法的解释器
func (e Evaluator) Interpret(variables map[string]Expression) int {
	return e.syntaxTree.Interpret(variables)
}
//解释加法运算，格式："x y +"
func NewEvaluator(expression string)*Evaluator  {
	expressionStack := new(Stack)
	//将表达式的值取出来，以空格为分隔符
	for _, token := range strings.Split(expression, " ") {
		switch token {
		// 发现是加号
		case "+":
			// 取出前面一个表达式
			right := expressionStack.Pop().(Expression)
			// 取出前面第二个表达式
			left := expressionStack.Pop().(Expression)
			// 组合成一个新的加法表达式
			subExpression := &Plus{left, right}
			// 将加法表达式放入栈中
			expressionStack.Push(subExpression)
		default:
			// 除了+号之外的变量标识符放入表达式的栈中
			expressionStack.Push(&Variable{token})
		}
	}
	// 取出一个表达式
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree: syntaxTree}
}