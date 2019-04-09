package main

import (
	"os"
	"testing"
)

func setup() {
	println("do something")
}

func tearDown() {
	println("do something")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestFoo(t *testing.T) {
	// setup will be called
	println("do something")
	// tearDown will be called
}
