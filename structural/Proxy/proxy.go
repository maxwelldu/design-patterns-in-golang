/**
当无法或不想直接引用某个对象或访问某个对象存在困难时，可以通过代理对象来间接访问。使用代理模式主要有两个目的：一是保护目标对象，二是增强目标对象。

前面分析了代理模式的结构与特点，现在来分析以下的应用场景。
	远程代理，这种方式通常是为了隐藏目标对象存在于不同地址空间的事实，方便客户端访问。例如，用户申请某些网盘空间时，会在用户的文件系统中建立一个虚拟的硬盘，用户访问虚拟硬盘时实际访问的是网盘空间。
	虚拟代理，这种方式通常用于要创建的目标对象开销很大时。例如，下载一幅很大的图像需要很长时间，因某种计算比较复杂而短时间无法完成，这时可以先用小比例的虚拟代理替换真实的对象，消除用户对服务器慢的感觉。
	安全代理，这种方式通常用于控制不同种类客户对真实对象的访问权限。
	智能指引，主要用于调用目标对象时，代理附加一些额外的处理功能。例如，增加计算真实对象的引用次数的功能，这样当该对象没有被引用时，就可以自动释放它。
	延迟加载，指为了提高系统的性能，延迟对目标的加载。例如，Hibernate 中就存在属性的延迟加载和关联表的延时加载。

 */
package Proxy
import (
	"fmt"
	"strconv"
)
//代理（Proxy）类：提供了与真实主题相同的接口，其内部含有对真实主题的引用，它可以访问、控制或扩展真实主题的功能。
type ITask interface {
	RentHouse(desc string, price int)
}

//真实主题（Real Subject）类：实现了抽象主题中的具体业务，是代理对象所代表的真实对象，是最终要引用的对象。
type Task struct { }
func (e *Task) RentHouse(desc string, price int) {
	fmt.Println(fmt.Sprintf("租房地址%s，价钱%s，中介费%s", desc, strconv.Itoa(price), strconv.Itoa(price)))
}

//抽象主题（Subject）类：通过接口或抽象类声明真实主题和代理对象实现的业务方法。
type AgentTask struct {
	task *Task
}
func NewAgentTask() *AgentTask {
	return &AgentTask{task: &Task{}}
}
func (e *AgentTask) RentHouse(desc string, price int) {
	e.task.RentHouse(desc, price)
}