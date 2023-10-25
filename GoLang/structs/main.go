package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type employee struct {
	firstName string
	lastName  string
	contactDetails
}

type contactDetails struct {
	email   string
	pinCode int
}

func main() {
	//3 Ways to declaring variable
	// way : 1
	//sufi := person{"Sufiyan", "Ahmad"}
	//fmt.Println(sufi)

	// way : 2
	name := person{firstName: "Sufiyan", lastName: "Ahmad"}
	fmt.Println(name)

	//way : 3
	var suf person
	suf.firstName = "Sufiyan"
	suf.lastName = "Ahmad"
	fmt.Println(suf)
	fmt.Printf("%+v", suf)

	//Now print contactInfo struct using person struct
	jim := person{
		firstName: "Amit",
		lastName:  "kumar",
		contact: contactInfo{
			email:   "amit@gmai.com",
			pinCode: 560037,
		},
	}
	//Adding pointer here to update value
	jimPointer := &jim
	jimPointer.UpdateName("Sanjay")
	jim.print()

	//You can also write structs and assign value like below (check Employee struct)
	empName := employee{
		firstName: "Sabir",
		lastName:  "Ansari",
		contactDetails: contactDetails{
			email:   "sabir@gmail.com",
			pinCode: 560037,
		},
	}
	fmt.Printf("%+v", empName)
}

//Update some fields of person
//Here made changes according to pointer
func (pointerToPerson *person) UpdateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
