/**
场景：
	当需要创建的对象是一系列相互关联或相互依赖的产品族时，如电器工厂中的电视机、洗衣机、空调等。
	系统中有多个产品族，但每次只使用其中的某一族产品。如有人只喜欢穿某一个品牌的衣服和鞋。
	系统中提供了产品的类库，且所有产品的接口相同，客户端不依赖产品实例的创建细节和内部结构。

抽象工厂模式的扩展有一定的“开闭原则”倾斜性：
当增加一个新的产品族时只需增加一个新的具体工厂，不需要修改原代码，满足开闭原则。
当产品族中需要增加一个新种类的产品时，则所有的工厂类都需要进行修改，不满足开闭原则。

另一方面，当系统中只存在一个等级结构的产品时，抽象工厂模式将退化到工厂方法模式。
 */
package AbstractFactory

import "fmt"

//抽象产品（Product）：定义了产品的规范，描述了产品的主要特性和功能，抽象工厂模式有多个抽象产品。
type Lunch interface {
	Cook()
}

//具体产品（ConcreteProduct）：实现了抽象产品角色所定义的接口，由具体工厂来创建，它同具体工厂之间是多对一的关系。
type Rise struct {
}
func (e *Rise) Cook() {
	fmt.Println("it is rise.")
}

//具体产品（ConcreteProduct）：实现了抽象产品角色所定义的接口，由具体工厂来创建，它同具体工厂之间是多对一的关系。
type Tomato struct {
}
func (e *Tomato) Cook() {
	fmt.Println("it is tomato.")
}

//抽象工厂（Abstract Factory）：提供了创建产品的接口，它包含多个创建产品的方法 newProduct()，可以创建多个不同等级的产品。
type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

//具体工厂（Concrete Factory）：主要是实现抽象工厂中的多个抽象方法，完成具体产品的创建。
type SimpleLunchFactory struct {
}
func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}
func (e *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}
func (e *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}