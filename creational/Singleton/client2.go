package Singleton

// 模拟一个client
func IncrementAge2() {
	p := GetInstance()
	p.IncrementAge()
}
