/**
当一个类内部具备两种或多种变化维度时，使用桥接模式可以解耦这些变化的维度，使高层代码架构稳定。

场景。
	当一个类存在两个独立变化的维度，且这两个维度都需要进行扩展时。
	当一个系统不希望使用继承或因为多层次继承导致系统类的个数急剧增加时。
	当一个系统需要在构件的抽象化角色和具体化角色之间增加更多的灵活性时。

	桥接模式的一个常见使用场景就是替换继承。我们知道，继承拥有很多优点，比如，抽象、封装、多态等，父类封装共性，子类实现特性。继承可以很好的实现代码复用（封装）的功能，但这也是继承的一大缺点。

	因为父类拥有的方法，子类也会继承得到，无论子类需不需要，这说明继承具备强侵入性（父类代码侵入子类），同时会导致子类臃肿。因此，在设计模式中，有一个原则为优先使用组合/聚合，而不是继承
 */
package Bridge

import "fmt"

//实现化（Implementor）角色：定义实现化角色的接口，供扩展抽象化角色调用。
type Draw interface {
	DrawCircle(radius, x, y int)
}

//具体实现化（Concrete Implementor）角色：给出实现化角色接口的具体实现。
type RedCircle struct {
}
func (e *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("red radius, x, y", radius, x, y)
}

//具体实现化（Concrete Implementor）角色：给出实现化角色接口的具体实现。
type YellowCircle struct {
}
func (e *YellowCircle) DrawCircle(radius, x, y int) {
	fmt.Println("yellow radius, x, y", radius, x, y)
}

//抽象化（Abstraction）角色：定义抽象类，并包含一个对实现化对象的引用。使用组合代替继承，实现了桥接
type Shape struct {
	draw Draw
}
func (e *Shape) Shape(d Draw) {
	e.draw = d
}

//扩展抽象化（Refined Abstraction）角色：是抽象化角色的子类，实现父类中的业务方法，并通过组合关系调用实现化角色中的业务方法。
type Circle struct {
	shape Shape
	x int
	y int
	radius int
}
func (e *Circle) Constructor(x, y, radius int, draw Draw) {
	e.x = x
	e.y = y
	e.radius = radius
	e.shape.Shape(draw)
}
func (e *Circle) Draw() {
	//e.shape 桥接对象，将圆和绘制类组合起来，这样方便扩展，当需要维护其他的圆，则只需要新增具体的圆即可
	e.shape.draw.DrawCircle(e.radius, e.x, e.y)
}