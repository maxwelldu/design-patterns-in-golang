package State

import "testing"

func TestState(t *testing.T) {
	machine := NewMachine() //初始状态是OFF
	machine.Off() //OFF.Off
	machine.On() //OFF.On 状态变为 ON
	machine.On() //ON.On
	machine.Off() //ON.Off 状态变为 OFF
}
