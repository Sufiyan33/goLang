package main

import "fmt"

type Employee struct {
	name    string
	salary  float64
	domain  string
	city    string
	country string
}

// method of a receiver type employee
func (e Employee) printDetails() {
	fmt.Println("Employee name is: ", e.name)
	fmt.Println("Employee salary is: ", e.salary)
	fmt.Println("Employee domain is: ", e.domain)
	fmt.Println("Employee city is: ", e.city)
	fmt.Println("Employee country is: ", e.country)
}
