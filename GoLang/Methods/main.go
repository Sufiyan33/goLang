package main

import "fmt"

func main() {

	// Now initialize struct & it's values
	values := Employee{
		name:    "ABCD",
		salary:  123970.23,
		domain:  "Health Care",
		city:    "Bangalore",
		country: "India",
	}
	values.printDetails()

	fmt.Println("-----------------------------")

	x := data(20)
	y := data(30)
	result := x.addition(y)
	fmt.Println(result)

	fmt.Println("-----------------------------")

	res := Student{
		name:    "Amit",
		roll_no: 123,
		branch:  "CSE",
		city:    "Kanpur",
		country: "India",
	}

	fmt.Println("Student name is (before):", res.name)
	fmt.Println("Student roll no is:", res.roll_no)
	fmt.Println("Student branch is (before):", res.branch)
	fmt.Println("Student city is:", res.city)
	fmt.Println("Student country is:", res.country)

	//creating a pointer
	ptr := &res
	ptr.showDetails("ME", "Sufiyan")
	fmt.Println("After creating pointer values are :")
	fmt.Println("Student name is (after):", res.name)
	fmt.Println("Student roll no is:", res.roll_no)
	fmt.Println("Student branch is (after):", res.branch)
	fmt.Println("Student city is:", res.city)
	fmt.Println("Student country is:", res.country)
}
