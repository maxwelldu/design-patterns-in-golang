/**
场景：
	需要频繁创建的一些类，使用单例可以降低系统的内存压力，减少 GC。
	某类只要求生成一个对象的时候，如一个班中的班长、每个人的身份证号等。
	某些类创建实例时占用资源较多，或实例化耗时较长，且经常使用。
	某类需要频繁实例化，而创建的对象又频繁被销毁的时候，如多线程的线程池、网络连接池等。
	频繁访问数据库或文件的对象。
	对于一些控制硬件级别的操作，或者从系统上来讲应当是单一控制逻辑的操作，如果有多个实例，则系统会完全乱套。
	当对象需要被共享的场合。由于单例模式只允许创建一个对象，共享该对象可以节省内存，并加快对象访问速度。如 Web 中的配置对象、数据库的连接池等。
 */
package Singleton

import "sync"

var (
	p *Pet
	once sync.Once
)
//初始化的时候就构造出来
func init() {
	// golang中最重要的是这个
	once.Do(func() {
		p = &Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age int
	mux sync.Mutex
}

//为了线程安全，加个锁
func (e *Pet) SetName(name string) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.name = name
}
// 这个地方读的地方尽量不要加锁，性能会高些，这里是单例
func (e *Pet) GetName() string {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.name
}
func (e *Pet) IncrementAge() {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.age++
}
func (e *Pet) GetAge() int {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.age
}
