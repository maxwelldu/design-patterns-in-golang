package Singleton

// 模拟一个client
func IncrementAge1() {
	p := GetInstance()
	p.IncrementAge()
}
