/**
场景：
	对象之间相同或相似，即只是个别的几个属性不同的时候。
	创建对象成本较大，例如初始化时间长，占用CPU太多，或者占用网络资源太多等，需要优化资源。
	创建一个对象需要繁琐的数据准备或访问权限等，需要提高性能或者提高安全性。
	系统中大量使用该类对象，且各个调用者都需要给它的属性重新赋值。
 */
package Prototype

//抽象原型类：规定了具体原型对象必须实现的接口。
type Prototype interface {
	Name() string
	Clone() Prototype
}
//具体原型类：实现抽象原型类的 clone() 方法，它是可被复制的对象。
type ConcretePrototype struct {
	name string
}

func (e *ConcretePrototype) Name() string {
	return e.name
}

func (e *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: e.name}
}
