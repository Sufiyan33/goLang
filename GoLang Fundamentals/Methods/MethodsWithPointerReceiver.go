package main

type Student struct {
	name    string
	roll_no int
	branch  string
	city    string
	country string
}

func (s *Student) showDetails(branch, name string) {
	(*s).branch = branch
	(*s).name = name
}
