/*
场景：
	当一个对象的行为取决于它的状态，并且它必须在运行时根据状态改变它的行为时，就可以考虑使用状态模式。
	一个操作中含有庞大的分支结构，并且这些分支决定于对象的状态时。
 */
package State

import "fmt"

//抽象状态（State）角色：定义一个接口，用以封装环境对象中的特定状态所对应的行为，可以有一个或多个行为。
type State interface {
	On(m *Machine)
	Off(m *Machine)
}

//具体状态（Concrete State）角色：实现抽象状态所对应的行为，并且在需要的情况下进行状态切换。
type ON struct {}
func NewON() State {
	return &ON{}
}
func (e *ON) On(m *Machine) {
	fmt.Println("设置已经开启...")
}
func (e *ON) Off(m *Machine) {
	fmt.Println("从On的状态到Off...")
	// 状态变更
	m.setCurrent(NewOFF())
}

//具体状态（Concrete State）角色：实现抽象状态所对应的行为，并且在需要的情况下进行状态切换。
type OFF struct {}
func NewOFF() State {
	return &OFF{}
}
func (e *OFF) On(m *Machine) {
	fmt.Println("从Off状态到on...")
	// 状态变更
	m.setCurrent(NewON())
}
func (e *OFF) Off(m *Machine) {
	fmt.Println("已经关闭")
}

//环境类（Context）角色：也称为上下文，它定义了客户端需要的接口，内部维护一个当前状态，并负责具体状态的切换。
type Machine struct {
	current State
}
func NewMachine() *Machine {
	// 初始是关闭的状态
	return &Machine{current: NewOFF()}
}
func (e *Machine) setCurrent(s State) {
	e.current = s
}
func (e *Machine) On() {
	e.current.On(e)
}
func (e *Machine) Off() {
	e.current.Off(e)
}
