package main

// Logger is an interface which will be used for an argument of a function.
type Logger interface {
	Printf(string, ...interface{})
}

// FooController is a struct which will be returned by function.
type FooController struct {
	Logger Logger
}

// NewFooController is a function for an example, "Accept interfaces, return structs".
// Also, this style of a function take on a role of constructor for struct.
func NewFooController(logger Logger) *FooController {
	return &FooController{
		Logger: logger,
	}
}
