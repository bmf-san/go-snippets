package main

import (
	"fmt"
	"reflect"
	"testing"
)

func testType(t *testing.T) {
	actual := reflect.TypeOf("Actual").String()
	expected := "string"

	if actual != expected {
		fmt.Printf("actual: %v, expected: %v", actual, expected)
	}
}
