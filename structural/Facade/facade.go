/**
场景：
	对分层结构系统构建时，使用外观模式定义子系统中每层的入口点可以简化子系统之间的依赖关系。
	当一个复杂系统的子系统很多时，外观模式可以为系统设计一个简单的接口供外界访问。
	当客户端与多个子系统之间存在很大的联系时，引入外观模式可将它们分离，从而提高子系统的独立性和可移植性。
 */
package Facade

import "fmt"

//子系统（Sub System）角色：实现系统的部分功能，客户可以通过外观角色访问它。
type CarModel struct {
}
func NewCarModel() *CarModel {
	return &CarModel{}
}
func (e *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}

//子系统（Sub System）角色：实现系统的部分功能，客户可以通过外观角色访问它。
type CarEngine struct {}
func NewCarEngine() *CarEngine {
	return &CarEngine{}
}
func (e *CarEngine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}

//子系统（Sub System）角色：实现系统的部分功能，客户可以通过外观角色访问它。
type CarBody struct {}
func NewCarBody() *CarBody {
	return &CarBody{}
}
func (e *CarBody) SetCarBody() {
	fmt.Println("CarBody Set Body")
}

//外观（Facade）角色：为多个子系统对外提供一个共同的接口。
type CarFacade struct {
	model *CarModel
	engine *CarEngine
	body *CarBody
}
func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  NewCarModel(),
		engine: NewCarEngine(),
		body:   NewCarBody(),
	}
}
func (e *CarFacade) CreateCompleteCar() {
	e.model.SetModel()
	e.engine.SetEngine()
	e.body.SetCarBody()
}