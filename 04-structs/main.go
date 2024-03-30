package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // it will come from the contactInfo struct which is created in L.NO.5
}

func main() {
	/*
		// *** WAY 1 ***
		shayan := person{firstName: "Shayan", lastName: "Saha"}
		fmt.Println(shayan)

		// *** WAY 2 ***
		var irin person

		irin.firstName = "Irin"
		irin.lastName = "Banik"

		fmt.Printf("%+v", irin) // %+v prints the fieldnames and value of a struct in the console
	*/

	shayan := person{
		firstName: "Shayan",
		lastName:  "Saha",
		contact: contactInfo{
			email:   "shayan851997@gmail.com",
			pinCode: 799004,
		},
	}

	shayanPointer := &shayan // address of struct "shayan" in RAM. we write it using '&' (ampersand). without this line also it will work. it is a shortcut with pointers in golang

	shayanPointer.updateEmail("shayansaha.con@gmail.com")
	shayan.print()

}

func (pointerToPerson *person) updateEmail(newEmail string) {
	(*pointerToPerson).contact.email = newEmail
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
