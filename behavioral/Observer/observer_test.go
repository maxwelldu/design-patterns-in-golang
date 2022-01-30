package Observer

import (
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	// 创建一个具体的主题
	n := eventSubject{ Observers: sync.Map{} }

	//  创建两个观察者
	obs1 := eventObserver{ID: 1, Time: time.Now()}
	obs2 := eventObserver{ID: 2, Time: time.Now()}

	// 将两个观察者添加到主题中
	n.AddListener(obs1)
	n.AddListener(obs2)

	for x := range Fib(10) {
		// 主题通知，会通知之前添加的2个观察者
		n.Notify(Event{Data: x})
	}
}
