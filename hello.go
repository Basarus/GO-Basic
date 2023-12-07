package main

import (
	"fmt"

	"./internal/contacts"
	"./tasks"
)

func main() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
	fmt.Println(tasks.IncB(20))
}
