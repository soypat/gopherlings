// Problem:

// I AM STILL GOING

package main

import "fmt"

func main() {
	db := userDatabase{
		admins: []string{"clair@example.com"},
	}
	// We'd like to be able to add an admin.
	// Create a method that can add users to the database.
	db.AddAdmin("richie@example.com")
	fmt.Println(db.IsAdmin("richie@example.com"), db.IsAdmin("ned@hack.com"))
}

type userDatabase struct {
	admins []string
}

// Methods are usually found below the type definition in Go.

// IsAdmin returns true if the email is inside the admin database
// and false otherwise.
func (u userDatabase) IsAdmin(email string) bool {
	for _, got := range u.admins {
		if got == email {
			return true
		}
	}
	return false
}
