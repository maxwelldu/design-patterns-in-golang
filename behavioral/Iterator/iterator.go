/**
场景，迭代器模式通常在以下几种情况使用。
	当需要为聚合对象提供多种遍历方式时。
	当需要为遍历不同的聚合结构提供一个统一的接口时。
	当访问一个聚合对象的内容而无须暴露其内部细节的表示时。

由于聚合与迭代器的关系非常密切，所以大多数语言在实现聚合类时都提供了迭代器类，因此大数情况下使用语言中已有的聚合类的迭代器就已经够了。
 */
package Iterator

//抽象迭代器（Iterator）角色：定义访问和遍历聚合元素的接口，通常包含 hasNext()、first()、next() 等方法。
type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

//具体迭代器（Concretelterator）角色：实现抽象迭代器接口中所定义的方法，完成对聚合对象的遍历，记录遍历的当前位置。
type ArrayIterator struct {
	array []interface{}
	index *int
}

func (e *ArrayIterator) Index() *int {
	return e.index
}
func (e *ArrayIterator) Value() interface{} {
	return e.array[*e.index]
}
func (e *ArrayIterator) HasNext() bool {
	return *e.index+1 <= len(e.array)
}
func (e *ArrayIterator) Next() {
	if e.HasNext() {
		*e.index++
	}
}