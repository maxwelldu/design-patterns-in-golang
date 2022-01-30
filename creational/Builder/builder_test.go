package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	builder := NewConcreteBuilder() //builder.built = false
	director := NewDirector(&builder) // 将builder注入到Director类中
	director.Construct() //调用builder.Build方法，将 builder.built = true
	result := builder.GetResult() //返回 builder.buit 的值
	fmt.Println(result.Built, result.Head)
}