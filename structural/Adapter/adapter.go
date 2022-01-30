/**
场景:
	以前开发的系统存在满足新系统功能需求的类，但其接口同新系统的接口不一致。
	使用第三方提供的组件，但组件接口定义和自己要求的接口定义不同。
 */
package Adapter

import "fmt"

//目标（Target）接口：当前系统业务所期待的接口，它可以是抽象类或接口。
type Target interface {
	Execute()
}

//适配者（Adaptee）类：它是被访问和适配的现存组件库中的组件接口。
type Adaptee struct {
}
func (e *Adaptee) SpecificExecute() {
	fmt.Println("充电")
}

//适配器（Adapter）类：它是一个转换器，通过继承或引用适配者的对象，把适配者接口转换成目标接口，让客户按目标接口的格式访问适配者。
type Adapter struct {
	*Adaptee
}
func (e *Adapter) Execute() {
	e.SpecificExecute()
}
