package SimpleFactory

import "fmt"

// 1.简单工厂
//抽象产品（Product）：是简单工厂创建的所有对象的父类，负责描述所有实例共有的公共接口。
type Restaurant interface {
	GetFood()
}

//具体产品（ConcreteProduct）：是简单工厂模式的创建目标。
type KFC struct {
}
func (e *KFC) GetFood() {
	fmt.Println("KFC的汉堡已经生产完毕，就绪。。。")
}

//具体产品（ConcreteProduct）：是简单工厂模式的创建目标。
type Subway struct {
}
func (e *Subway) GetFood() {
	fmt.Println("赛百味的三明治已经生产完毕，就绪。。。")
}

//简单工厂（SimpleFactory）：是简单工厂模式的核心，负责实现创建所有实例的内部逻辑。工厂类的创建产品类的方法可以被外界直接调用，创建所需的产品对象。
func NewRestaurant(name string) Restaurant {
	switch name {
	case "s":
		return &Subway{}
	case "k":
		return &KFC{}
	}
	return nil
}
