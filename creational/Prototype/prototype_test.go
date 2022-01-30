package Prototype

import (
	"testing"

	"github.com/golib/assert"
)

//访问类：使用具体原型类中的 clone() 方法来复制新的对象。
func TestConcretePrototype_Clone(t *testing.T) {
	name := "Give me five"
	e := ConcretePrototype{name: name}
	newProto := e.Clone()
	assert.Equal(t, name, newProto.Name())
}
