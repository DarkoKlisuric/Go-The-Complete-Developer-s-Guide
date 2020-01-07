package main

import "fmt"

type contactInfo struct {
	email string
	yearOfBirth int
}
type person struct {
	firstName string
	lastName string
	contactInfo
}

func main()  {
	/*
	// 1# way to define struct
	 darko := person{"Darko", "Klisurić"}
	// 2# way to define struct
	 klisuric := person{firstName: "Darko", lastName: "Klisurić"}

	var darkoKlisuric person
	darkoKlisuric.lastName = "Klisurić"
	darkoKlisuric.firstName = "Darko"

	fmt.Printf("%+v" , darkoKlisuric)
	fmt.Println(darkoKlisuric)
	fmt.Println(klisuric, darko)
	*/

	// embedded struct
	marko := person{
		firstName: "Marko",
		lastName:  "Markić",
		contactInfo: contactInfo{
			email: "marko@marko.com",
			yearOfBirth: 1995,
		},
	}
	markoPointer := &marko
	markoPointer.updateName("Markan")
	marko.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
