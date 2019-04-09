package main

// SQLHandler is an interface for handling sql.
type SQLHandler interface {
	Execute()
}

// sqlHandler is a struct which will be returned by function.
type sqlHandler struct{}

// NewSQLHandler is a function for an example of DIP.
// This function depend on abstruction(interface).
// This pattern is an idiom of constructor in golang.
// You can do DI(Dependency Injection) by using nested struct.
func NewSQLHandler() SQLHandler {
	// do something ...

	// sqlHandler struct implments SQLHandler interface.
	return &sqlHandler{}
}

// Execute is a function for executing sql.
// A sqlHanlder struct implments a SQLHandler interface by defining Execute().
func (s *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on an interface.
type FooRepository struct {
	SQLHandler SQLHandler
}

// Find is a method of FooRepository depending on an interface.
func (ur *FooRepository) Find() {
	// do something
	ur.SQLHandler.Execute()
}
