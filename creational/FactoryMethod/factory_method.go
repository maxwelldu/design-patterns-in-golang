/**
场景：
	客户只知道创建产品的工厂名，而不知道具体的产品名。如 TCL 电视工厂、海信电视工厂等。
	创建对象的任务由多个具体子工厂中的某一个完成，而抽象工厂只提供创建产品的接口。
	客户不关心创建产品的细节，只关心产品的品牌
 */
package FactoryMethod

import "fmt"

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

//抽象工厂（Abstract Factory）：提供了创建产品的接口，调用者通过它访问具体工厂的工厂方法 newProduct() 来创建产品。
type AbstractFactory interface {
	NewRestaurant()
}

//具体工厂（ConcreteFactory）：主要是实现抽象工厂中的抽象方法，完成具体产品的创建。
type KFCFactory struct {
}
func (e *KFCFactory) NewRestaurant() Restaurant {
	return &KFC{}
}

//具体工厂（ConcreteFactory）：主要是实现抽象工厂中的抽象方法，完成具体产品的创建。
type SubWayFactory struct {
}
func (e *SubWayFactory) NewRestaurant() Restaurant {
	return &Subway{}
}

