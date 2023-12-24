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
	fmt.Println(tasks.IncB(1))
	fmt.Println(tasks.HelloUser(1997))
	tasks.Cycle(1, 10)
	tasks.GoTo()
	tasks.FizzBuzz()
	//tasks.Ucaz()
	//tasks.Stroke()
	tasks.Slace(5)
}
