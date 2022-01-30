/*
当系统的某项操作具备命令语义，且命令实现不稳定（变化）时，可以通过命令模式解耦请求与实现。使用抽象命令接口使请求方的代码架构稳定，封装接收方具体命令的实现细节。接收方与抽象命令呈现弱耦合（内部方法无需一致），具备良好的扩展性。

场景:
	请求调用者需要与请求接收者解耦时，命令模式可以使调用者和接收者不直接交互。
	系统随机请求命令或经常增加、删除命令时，命令模式可以方便地实现这些功能。
	当系统需要执行一组操作时，命令模式可以定义宏命令来实现该功能。
	当系统需要支持命令的撤销（Undo）操作和恢复（Redo）操作时，可以将命令对象存储起来，采用备忘录模式来实现。
 */
//案例来源：https://www.cnblogs.com/amunote/p/15362468.html
package Command

import "fmt"

//抽象命令类（Command）角色：声明执行命令的接口，拥有执行命令的抽象方法 execute()。
type Command interface {
	Execute()
}

//具体命令类（Concrete Command）角色：是抽象命令类的具体实现类，它拥有接收者对象，并通过调用接收者的功能来完成命令要执行的操作。
type onCommand struct {
	device Device
}
func (e *onCommand) Execute() {
	e.device.on()
}

//具体命令类（Concrete Command）角色：是抽象命令类的具体实现类，它拥有接收者对象，并通过调用接收者的功能来完成命令要执行的操作。
type offCommand struct {
	device Device
}
func (e *offCommand) Execute() {
	e.device.off()
}


//Receiver interface
type Device interface {
	on()
	off()
}

//实现者/接收者（Receiver）角色：执行命令功能的相关操作，是具体命令对象业务的真正实现者。
type Tv struct {
	isRunning bool
}
func (e *Tv) on() {
	e.isRunning = true
	fmt.Println("Turning tv on")
}
func (e *Tv) off() {
	e.isRunning = false
	fmt.Println("Turning tv off")
}

//调用者/请求者（Invoker）角色：是请求的发送者，它通常拥有很多的命令对象，并通过访问命令对象来执行相关请求，它不直接访问接收者。
type button struct {
	command Command
}
func (e *button) press() {
	e.command.Execute()
}