package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

type ValidationError struct {
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s", ve.Message)
}

func NewUser(name string, age int, email string) (*User, error) {
	if name == "" {
		return nil, &ValidationError{"Name cannot be empty"}
	}
	if age <= 0 || age > 150 {
		return nil, &ValidationError{"Age must be between 1 and 150"}
	}
	if !strings.Contains(email, "@") {
		return nil, &ValidationError{"Email must contain '@'"}
	}
	return &User{Name: name, Age: age, Email: email}, nil
}

func (user *User) IsAdult() bool {
	return user.Age >= 18
}

func main() {
	name := "John Doe"
	age := 25
	email := "john.doe@gmail.com"
	user, err := NewUser(name, age, email)
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	fmt.Println("User created: ", user.Name)
	fmt.Println("Is adult: ", user.IsAdult())
}