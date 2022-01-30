package TemplateMethod
import "testing"

func Test(t *testing.T) {
	worker := NewWorker(&Coder{})
	worker.Daily()
}