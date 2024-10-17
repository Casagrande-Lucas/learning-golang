package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Defining clear data models with structs and leveraging type safety.

// Best Practice: Use struct tags for serialization and validation.
// Define methods on structs to encapsulate behavior.

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func createUser(username, email string) User {
	return User{
		ID:        rand.Int(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func main() {
	username := "Alice"
	email := "alice@example.com"
	user := createUser(username, email)
	fmt.Println(user)
}
