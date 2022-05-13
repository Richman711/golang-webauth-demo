package main

import (
	"log"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "1234567890"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}
	log.Println(hashedPass)

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")
}

func hashPassword(password string)([]byte, error) {
	log.Println(password)
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}