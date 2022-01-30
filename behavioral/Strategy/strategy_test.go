package Strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	e := NewContext()

	strategyA := NewStrategyA()

	// 使用策略A
	e.SetStrategy(strategyA)
	e.Execute()

	strategyB := NewStrategyB()

	//使用策略B
	e.SetStrategy(strategyB)
	e.Execute()
}