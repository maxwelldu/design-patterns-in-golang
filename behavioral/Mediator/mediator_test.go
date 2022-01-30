package Mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := NewConcreteMediator()
	mediator.Ted.Talk()
}
