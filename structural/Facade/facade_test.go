package Facade

import "testing"

//客户（Client）角色：通过一个外观角色访问各个子系统的功能。
func TestCarFacade_CreateCompleteCar(t *testing.T) {
	f := NewCarFacade()
	f.CreateCompleteCar()
}
