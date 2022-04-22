package main

import "fmt"

type user struct {
	name  string
	email string
}

/*The receiver for notify is declared as a value of type user. When you declare a method
using a value receiver, the method will always be operating against a copy of the value
used to make the method call */
func (u user) changeEmail(email string) {
	u.email = email
}

func (u *user) changeEmailPtr(email string) {
	u.email = email
}

/*
	The parameter between the keyword func and the function name is called a receiver and binds the function to
	the specified type. When a function has a receiver, that function is called a method.
*/

func main() {
	newUser := user{"Anulav", "anulav@email.com"}
	newUser.changeEmail("anulav@gmail.com") // Acts on the copy of the value only.
	fmt.Println("Without ptr", newUser)

	newUser.changeEmailPtr("anulav@gmail.com")
	fmt.Println("With ptr", newUser) //Acts on the object ptr passed.

	//Pointers of type user can also be used to call its methods.
	newUser2 := &user{"Bruce Wayne", "batman@gotham.com"}
	newUser2.changeEmail("joker@gotham.com") //Internally Go compiler dereferences the pointer so that the method call can be in compliance with the value receiver. Like internally it does, (*newUser2).changeEmail("joker@gotham.com")
	fmt.Println("Wihtout ptr", newUser2)

	newUser2.changeEmailPtr("joker@gotham.com")
	fmt.Println("With ptr", newUser2)

}

/*
	Which type of receiver to choose value or pointer for a method?
	For this ask, Does adding or removing something from a value of this type need to create
	a new value or mutate the existing one? If the answer is create a new value, then use
	value receivers for your methods. If the answer is mutate the value, then use pointer
	receivers.

	Although the built-in types such as strings, numberic and boolean should only be passed as a
	value.

*/
