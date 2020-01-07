package main

import "fmt"

type contactInfo struct {
	email string
	yearOfBirth int
}
type person struct {
	firstName string
	lastName string
	contact contactInfo
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
		contact:   contactInfo{
			email: "marko@marko.com",
			yearOfBirth: 1995,
		},
	}

	fmt.Printf("%+v" , marko)


}
