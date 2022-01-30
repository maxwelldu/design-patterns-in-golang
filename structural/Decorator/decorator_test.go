package Decorator

import "testing"
import "log"
import "os"

func TestPi(t *testing.T) {
	f := WrapLogger(Pi, log.New(os.Stdout, "test", 1))
	f(5000)
}