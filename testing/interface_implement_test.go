package main

import (
	"testing"
)

func TestInterfaceImplement(t *testing.T) {
	bar, ok := interface{}(bar).(foo)
	if !ok {
		t.Errorf("%v doesn't implement foo.", bar)
	}
}
