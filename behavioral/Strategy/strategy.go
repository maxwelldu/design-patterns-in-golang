/*
策略模式在很多地方用到，如 Java SE 中的容器布局管理就是一个典型的实例，Java SE 中的每个容器都存在多种布局供用户选择。在程序设计中，通常在以下几种情况中使用策略模式较多。
场景：
	一个系统需要动态地在几种算法中选择一种时，可将每个算法封装到策略类中。
	一个类定义了多种行为，并且这些行为在这个类的操作中以多个条件语句的形式出现，可将每个条件分支移入它们各自的策略类中以代替这些条件语句。
	系统中各算法彼此完全独立，且要求对客户隐藏具体算法的实现细节时。
	系统要求使用算法的客户不应该知道其操作的数据时，可使用策略模式来隐藏与算法相关的数据结构。
	多个类只区别在表现行为不同，可以使用策略模式，在运行时动态选择具体要执行的行为。
 */
package Strategy

import "fmt"

//抽象策略（Strategy）类：定义了一个公共接口，各种不同的算法以不同的方式实现这个接口，环境角色使用这个接口调用不同的算法，一般使用接口或抽象类实现。
type Strategy interface {
	Execute()
}

//具体策略（Concrete Strategy）类：实现了抽象策略定义的接口，提供具体的算法实现。
type strategyA struct {}
func NewStrategyA() Strategy {
	return &strategyA{}
}
func (e *strategyA) Execute() {
	fmt.Println("A plan Executed.")
}

//具体策略（Concrete Strategy）类：实现了抽象策略定义的接口，提供具体的算法实现。
type strategyB struct {}
func NewStrategyB() Strategy {
	return &strategyB{}
}
func (e *strategyB) Execute() {
	fmt.Println("B plan Executed.")
}

//环境（Context）类：持有一个策略类的引用，最终给客户端调用。
type Context struct {
	strategy Strategy
}
func NewContext() *Context {
	return &Context{}
}
func (e *Context) SetStrategy(strategy Strategy) {
	e.strategy = strategy
}
func (e *Context) Execute() {
	e.strategy.Execute()
}