package Memento

import "testing"

func TestMember_ReinstateMemento(t *testing.T) {
	e := NewNumber(7)
	e.Double()
	e.Double()

	// 记录在一个备忘录上
	memento := e.CreateMemento()
	e.Half()

	//从备忘录中拿出来
	e.ReinstateMemento(memento)
	t.Log(e.value)
}