/**
场景：
	当需要给一个现有类添加附加职责，而又不能采用生成子类的方法进行扩充时。例如，该类被隐藏或者该类是终极类或者采用继承方式会产生大量的子类。
	当需要通过对现有的一组基本功能进行排列组合而产生非常多的功能时，采用继承关系很难实现，而采用装饰器模式却很好实现。
	当对象的功能要求可以动态地添加，也可以再动态地撤销时。
 */
package Decorator

import "math"
import "log"
import "time"

type piFunc func(int) float64

// 不同语言实现设计模式可能不太一样；主要是在Java里面会抽象的比较厉害，因为 Java是纯面向对象的；Golang可以传递函数就简化很多
// 这里装饰一个Pi函数，为Pi添加一个记录耗时的功能
func WrapLogger(fun piFunc, logger *log.Logger) piFunc {
	return func(i int) float64 {
		fn := func(n int) (result float64){
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n , result)
			}(time.Now())
			return fun(n)
		}
		return fn(i)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go func(ch chan float64, i float64) {
			ch <- 4 * math.Pow(-1, i) / (2 * i + 1)
		}(ch, float64(i))
	}

	res := 0.0
	for i := 0; i < n; i++ {
		res += <-ch
	}
	return res
}
