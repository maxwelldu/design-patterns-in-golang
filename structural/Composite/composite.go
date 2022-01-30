/**
场景:
	在需要表示一个对象整体与部分的层次结构的场合。
	要求对用户隐藏组合对象与单个对象的不同，用户可以用统一的接口使用组合结构中的所有对象的场合。
 */
package Composite

import "fmt"

//抽象构件（Component）角色：它的主要作用是为树叶构件和树枝构件声明公共接口，并实现它们的默认行为。在透明式的组合模式中抽象构件还声明访问和管理子类的接口；在安全式的组合模式中不声明访问和管理子类的接口，管理工作由树枝构件完成。（总的抽象类或接口，定义一些通用的方法，比如新增、删除）
type Component interface {
	Traverse()
}

//树叶构件（Leaf）角色：是组合中的叶节点对象，它没有子节点，用于继承或实现抽象构件。
type Leaf struct {
	value int
}
func NewLeaf(value int) *Leaf {
	return &Leaf{value: value}
}
func (e *Leaf) Traverse() {
	fmt.Println(e.value)
}

//树枝构件（Composite）角色 / 中间构件：是组合中的分支节点对象，它有子节点，用于继承和实现抽象构件。它的主要作用是存储和管理子部件，通常包含 Add()、Remove()、GetChild() 等方法。
type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}
func (e *Composite) Add(component Component) {
	e.children = append(e.children, component)
}
func (e *Composite) Traverse() {
	for idx, _ := range e.children {
		e.children[idx].Traverse()
	}
}