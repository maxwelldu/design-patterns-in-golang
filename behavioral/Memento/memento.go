/*
场景:
	需要保存与恢复数据的场景，如玩游戏时的中间结果的存档功能。
	需要提供一个可回滚操作的场景，如 Word、记事本、Photoshop，Eclipse 等软件在编辑时按 Ctrl+Z 组合键，还有数据库中事务操作。
 */
package Memento

//备忘录（Memento）角色：负责存储发起人的内部状态，在需要的时候提供这些内部状态给发起人。
type Memento struct {
	state int
}
func NewMemento(value int) *Memento {
	return &Memento{value}
}

//发起人（Originator）角色：记录当前时刻的内部状态信息，提供创建备忘录和恢复备忘录数据的功能，实现其他业务功能，它可以访问备忘录里的所有信息。
type Number struct {
	value int
}
func NewNumber(value int) *Number {
	return &Number{value: value}
}
func (e *Number) Double() {
	e.value *= 2
}
func (e *Number) Half() {
	e.value /= 2
}
func (e *Number) Value() int {
	return e.value
}
func (e *Number) CreateMemento() *Memento {
	// 记录当前的一个状态
	return NewMemento(e.value)
}
func (e *Number) ReinstateMemento(memento *Memento) {
	// 从备忘录中将状态取出来
	e.value = memento.state
}
