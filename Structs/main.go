package main

import (
	"fmt";
	// "reflect"
)

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct{
	email string
	zip int32 
}

func main(){
	fb := person{
		firstName: "Fabio",
		lastName: "Bregasi",
		contact : contactInfo{
			email: "fabio.bregasi@yahoo.it",
			zip: 28100,
		},
	}
	fb.printPerson()
	//pointer to a person --> & = address in memory of the struct
	// fbPointer := &fb
	// fbPointer.updateName("Daniela")
	// Go does not need the reference to the direct address in memory:
	//even the variable name is enough to get the pointer through the function
	fb.updateName("Daniela")
	fb.printPerson()
	var t2 person
	t2.printPerson()
	t2.firstName = "Daniela"
	t2.lastName = "Perdeica"
	t2.contact = contactInfo{email: "danielaperdeica@gmail.com", zip: 13900}
	t2.printPerson()
}


func (p person) printPerson() {
	fmt.Printf("%+v\n", p)
}

func (pointerToAPerson *person) updateName(newName string) {
	(*pointerToAPerson).firstName = newName
}