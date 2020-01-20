package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	var pwd string
	// Read the input from stdin
	_, err := fmt.Scan(&pwd)
	if err != nil {
		panic(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hash))
}
