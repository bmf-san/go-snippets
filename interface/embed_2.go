// see: https://medium.com/eureka-engineering/golang-embedded-ac43201cf772
package main

import (
	"errors"
	"os/user"
	"testing"
)

// define struct
type User struct {
	Name   string
	Gender string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetGender() string {
	return u.Gender
}

// define an interface
type UserInterface interface {
	GetName() string
	GetGender() string
}

func GetNameAndGender(u UserInterface) string {
	return u.GetName() + u.GetGender()
}

// writes test codes
// mock struct
type mockUser struct {
	user.User // satisfy interface. duck typing.
	Name      string
}

// override a method if you want to replace method.
func (u *mockUser) GetName() string {
	return "mockname"
}

func TestGetName(t *testing.T) {
	u := mockUser{
		Name: "name",
	}

	u.GetName() // mockname
}

// type error is also an interface
func Exec() error { // error is an interface
	return errors.New("error")
}

type error interface {
	Error() string
}

// Also you can embed interface.
// In this case, if you access  access a method that is not defined by mockUserByInteraface, it will panic.
type mockUserByInteraface struct {
	UserInterface // satisfy interface
	Name          string
}

// override a method
func (u *mockUserByInteraface) GetName() string {
	return "mockname"
}

func TestGetNameByInterface(t *testing.T) {
	u := mockUserByInteraface{
		Name: "name",
	}

	u.GetName() // mockname
}
