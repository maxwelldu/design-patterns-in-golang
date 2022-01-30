/*
在软件系统中，当系统一方行为依赖另一方行为的变动时，可使用观察者模式松耦合联动双方，使得一方的变动可以通知到感兴趣的另一方对象，从而让另一方对象对此做出响应。

场景：
	对象间存在一对多关系，一个对象的状态发生改变会影响其他对象。
	当一个抽象模型有两个方面，其中一个方面依赖于另一方面时，可将这二者封装在独立的对象中以使它们可以各自独立地改变和复用。
	实现类似广播机制的功能，不需要知道具体收听者，只需分发广播，系统中感兴趣的对象会自动接收该广播。
	多层级嵌套使用，形成一种链式触发机制，使得事件具备跨域（跨越两种观察者类型）通知。
 */
package Observer

import (
	"sync"
	"time"
	"fmt"
)

type Event struct {
	Data int
}

//抽象观察者（Observer）角色：它是一个抽象类或接口，它包含了一个更新自己的抽象方法，当接到具体主题的更改通知时被调用。
type Observer interface {
	NotifyCallback(event Event)
}

//具体观察者（Concrete Observer）角色：实现抽象观察者中定义的抽象方法，以便在得到目标的更改通知时更新自身的状态。
type eventObserver struct {
	ID int
	Time time.Time
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Received:%d after %v\n", event.Data, time.Since(e.Time))
}

//抽象主题（Subject）角色：也叫抽象目标类，它提供了一个用于保存观察者对象的聚集类和增加、删除观察者对象的方法，以及通知所有观察者的抽象方法。
type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

//具体主题（Concrete Subject）角色：也叫具体目标类，它实现抽象目标中的通知方法，当具体主题的内部状态发生改变时，通知所有注册过的观察者对象。
type eventSubject struct {
	Observers sync.Map
}
func (e *eventSubject) AddListener(observer Observer) {
	e.Observers.Store(observer, struct {}{})
}
func (e *eventSubject) RemoveListener(observer Observer) {
	e.Observers.Delete(observer)
}
func (e *eventSubject) Notify(event Event)  {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0,1; i<n; i,j = i+j, i {
			out <- i
		}
	}()
	return out
}
