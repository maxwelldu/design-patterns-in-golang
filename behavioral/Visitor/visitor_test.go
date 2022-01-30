package Visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	// 比如是我自己
	e := new(Element)
	// 访问者，这可能跟现实中正好反过来
	e.accept(new(WeiBoVisitor))
	e.accept(new(IQiYiVisitor))
}
