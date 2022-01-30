/**
当系统中多处需要同一组信息时，可以把这些信息封装到一个对象中，然后对该对象进行缓存，这样，一个对象就可以提供给多出需要使用的地方，避免大量同一对象的多次创建，降低大量内存空间的消耗。

享元模式其实是工厂方法模式的一个改进机制，享元模式同样要求创建一个或一组对象，并且就是通过工厂方法模式生成对象的，只不过享元模式为工厂方法模式增加了缓存这一功能。

前面分析了享元模式的结构与特点，下面分析它适用的应用场景。享元模式是通过减少内存中对象的数量来节省内存空间的，所以以下几种情形适合采用享元模式。
	系统中存在大量相同或相似的对象，这些对象耗费大量的内存资源。
	大部分的对象可以按照内部状态进行分组，且可将不同部分外部化，这样每一个组只需保存一个内部状态。
	由于享元模式需要额外维护一个保存享元的数据结构，所以应当在有足够多的享元实例时才值得使用享元模式。
 */
package Flyweight

//具体享元（Concrete Flyweight）角色：实现抽象享元角色中所规定的接口。这里简单起见，没有写抽象享元角色
type FlyWeight struct {
	Name string
}
func NewFlyWeight(name string) *FlyWeight {
	return &FlyWeight{Name:name}
}

//享元工厂（Flyweight Factory）角色：负责创建和管理享元角色。当客户对象请求一个享元对象时，享元工厂检査系统中是否存在符合要求的享元对象，如果存在则提供给客户；如果不存在的话，则创建一个新的享元对象。
type FlyWeightFactory struct {
	pool map[string]*FlyWeight
}
func NewFlyWeightFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]*FlyWeight)}
}
func (e *FlyWeightFactory) GetFlyWeight(name string) *FlyWeight {
	weight, ok := e.pool[name]
	if !ok {
		weight = NewFlyWeight(name)
		e.pool[name] = weight
	}
	return weight
}
