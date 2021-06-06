package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//cleanString removes leading whitespace and quotes (as often get added when running as a commandline app in windows)
func cleanString(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func main() {

	var pwd string
	// Read the input from stdin
	_, err := fmt.Scan(&pwd)
	if err != nil {
		panic(err)
	}

	pwd = cleanString(pwd)

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hash))
}
