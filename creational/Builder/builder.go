/**
建造者模式唯一区别于工厂模式的是针对复杂对象的创建。也就是说，如果创建简单对象，通常都是使用工厂模式进行创建，而如果创建复杂对象，就可以考虑使用建造者模式。
当需要创建的产品具备复杂创建过程时，可以抽取出共性创建过程，然后交由具体实现类自定义创建流程，使得同样的创建行为可以生产出不同的产品，分离了创建与表示，使创建产品的灵活性大大增加。

场景：
	相同的方法，不同的执行顺序，产生不同的结果。
	多个部件或零件，都可以装配到一个对象中，但是产生的结果又不相同。
	产品类非常复杂，或者产品类中不同的调用顺序产生不同的作用。
	初始化一个对象特别复杂，参数多，而且很多参数都具有默认值。
 */
package Builder

// 行为和具体的实现是分离的
//抽象建造者（Builder）：它是一个包含创建产品各个子部件的抽象方法的接口，通常还包含一个返回复杂产品的方法 getResult()。
type Builder interface {
	Build()
	BuildHead()
}

//指挥者（Director）：它调用建造者对象中的部件构造与装配方法完成复杂对象的创建，在指挥者中不涉及具体产品的信息。
type Director struct {
	builder Builder
}
func (e *Director) Construct() {
	e.builder.Build()
	e.builder.BuildHead()
}
func NewDirector(b Builder) Director {
	return Director{b}
}

//具体建造者(Concrete Builder）：实现 Builder 接口，完成复杂产品的各个部件的具体创建方法。
type ConcreteBuilder struct {
	built bool
	head string
}
func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{built: false}
}
func (e *ConcreteBuilder) BuildHead() {
	e.head = "head"
}
func (e *ConcreteBuilder) Build() {
	e.built = true
}

//产品角色（Product）：它是包含多个组成部件的复杂对象，由具体建造者来创建其各个零部件。
type Product struct {
	Built bool
	Head string
}
func (e *ConcreteBuilder) GetResult() Product {
	return Product{e.built, e.head}
}