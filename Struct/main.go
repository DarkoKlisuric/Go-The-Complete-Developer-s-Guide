package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main()  {
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

}
