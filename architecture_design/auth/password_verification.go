package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hash(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func verify(hash, s string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
}

func main() {
	hash, err := hash("password")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	if err = verify(hash, "password"); err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("verified")
	}
}
