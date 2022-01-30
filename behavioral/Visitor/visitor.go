/*
当系统中存在类型数量稳定（固定）的一类数据结构时，可以使用访问者模式方便地实现对该类型所有数据结构的不同操作，而又不会对数据产生任何副作用（脏数据）。

简而言之，就是当对集合中的不同类型数据（类型数量稳定）进行多种操作时，使用访问者模式。

场景：
	对象结构相对稳定，但其操作算法经常变化的程序。
	对象结构中的对象需要提供多种不同且不相关的操作，而且要避免让这些操作的变化影响对象的结构。
	对象结构包含很多类型的对象，希望对这些对象实施一些依赖于其具体类型的操作。
 */
package Visitor

import "fmt"

//抽象访问者（Visitor）角色：定义一个访问具体元素的接口，为每个具体元素类对应一个访问操作 visit() ，该操作中的参数类型标识了被访问的具体元素。
type IVisitor interface {
	Visit()
}

//具体访问者（ConcreteVisitor）角色：实现抽象访问者角色中声明的各个访问操作，确定访问者访问一个元素时该做什么。
type WeiBoVisitor struct {
}
func (e *WeiBoVisitor) Visit() {
	fmt.Println("Visit WeiBo")
}

//具体访问者（ConcreteVisitor）角色：实现抽象访问者角色中声明的各个访问操作，确定访问者访问一个元素时该做什么。
type IQiYiVisitor struct {
}
func (e *IQiYiVisitor) Visit() {
	fmt.Println("Visit IQiYi")
}

//抽象元素（Element）角色：声明一个包含接受操作 accept() 的接口，被接受的访问者对象作为 accept() 方法的参数。
type IElement interface {
	Accept(visitor IVisitor)
}

//具体元素（ConcreteElement）角色：实现抽象元素角色提供的 accept() 操作，其方法体通常都是 visitor.visit(this) ，另外具体元素中可能还包含本身业务逻辑的相关操作。
type Element struct {
}
func (e *Element) accept(v IVisitor) {
	v.Visit()
}