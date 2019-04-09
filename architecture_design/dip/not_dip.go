package main

// sqlHandler is a struct for handling sql.
type sqlHandler struct{}

// Execute is a function for executing sql.
func (sqlHandler *sqlHandler) Execute() {
	// do something...
}

// FooRepository is a struct depending on details.
type FooRepository struct {
	sqlHandler sqlHandler
}

// Find is a method depending on details.
func (ur *FooRepository) Find() {
	// do something
	ur.sqlHandler.Execute()
}
