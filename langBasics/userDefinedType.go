package main

import "fmt"

func main() {
	var newUser user //Initialized with null values for each type of struct
	fmt.Println("newUser values", newUser)

	anulav := user{ //Declaring the User with properties. No order needs to be defined.
		name:  "Anulav Sharan",
		email: "anulav@email.com",
		exp:   4,
	}
	fmt.Println(anulav)
	anulav2 := user{"Anulav Sharan", "anulav@email.com", 4} //decalring the User without properties
	fmt.Println(anulav2)

	admin1 := admin{anulav2, 5}
	fmt.Println(admin1)

	var dur Duration
	dur = Duration(1000)
	fmt.Println(dur)

}

type user struct {
	name  string
	email string
	exp   int
}

type admin struct {
	person user
	level  int
}

type Duration int64 //New Type declared on int64
