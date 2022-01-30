package Interpreter

import (
	"testing"

	"github.com/golib/assert"
)

// client
func TestInterpret(t *testing.T) {
	// 待解释的表达式
	expression := "x y +"
	// 初始化解释器
	sentence := NewEvaluator(expression)

	// 准备好变量与映射的值，映射的值也是一个表达式
	variables := make(map[string]Expression)
	variables["x"] = &Integer{10}
	variables["y"] = &Integer{41}

	// 开始解释执行
	result := sentence.Interpret(variables)
	assert.Equal(t, 51, result)
}
