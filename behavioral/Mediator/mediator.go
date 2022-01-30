/*
场景:
	当对象之间存在复杂的网状结构关系而导致依赖关系混乱且难以复用时。
	当想创建一个运行于多个类之间的对象，又不想生成新的子类时。
 */
package Mediator

import "fmt"

//抽象同事类（Colleague）角色：定义同事类的接口，保存中介者对象，提供同事对象交互的抽象方法，实现所有相互影响的同事类的公共功能。
type WildStallion interface {
	SetMediator(mediator Mediator)
}

//具体同事类（Concrete Colleague）角色：是抽象同事类的实现者，当需要与其他同事对象交互时，由中介者对象负责后续的交互。
type Bill struct {
	mediator Mediator
}
func (e *Bill) SetMediator(mediator Mediator) {
	e.mediator = mediator
}
func (e *Bill) Response() {
	fmt.Println("bill what?")
	// 让中介通知,来自bill
	e.mediator.Communicate("bill")
}

//具体同事类（Concrete Colleague）角色：是抽象同事类的实现者，当需要与其他同事对象交互时，由中介者对象负责后续的交互。
type Ted struct {
	mediator Mediator
}
func (e *Ted) Talk() {
	fmt.Println("Ted: Bill?")
	//让中介通知,来自Ted
	e.mediator.Communicate("Ted")
}
func (e *Ted) SetMediator(mediator Mediator) {
	e.mediator = mediator
}
func (e *Ted) Response() {
	fmt.Println("Ted:how much?")
}


//抽象中介者（Mediator）角色：它是中介者的接口，提供了同事对象注册与转发同事对象信息的抽象方法。
type Mediator interface {
	Communicate(who string)
}
//具体中介者（Concrete Mediator）角色：实现中介者接口，定义一个 List 来管理同事对象，协调各个同事角色之间的交互关系，因此它依赖于同事角色。
type ConcreteMediator struct {
	Bill
	Ted
}
func NewConcreteMediator() *ConcreteMediator {
	mediator := &ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}
func (e *ConcreteMediator) Communicate(who string) {
	if who == "Ted" {
		e.Bill.Response()
	} else {
		e.Ted.Response()
	}
}