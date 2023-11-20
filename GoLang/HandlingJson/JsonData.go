package main

import (
	"encoding/json"
	"fmt"
)

/*
	Looks good but the field should be in lowerCase & password should not reflect & it should be nil not null.
	Let's make use of Alias to fix this issues. If slice is not present then not show nil instead of
	this hide that fields.
*/
type course struct {
	Name      string `json:"courseName"`
	Price     int
	Plateform string   `json:"website"`
	Password  string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func EnodeJson() {
	courseDetails := []course{
		{"Java BootCamp", 999, "udemy.com", "abc123", []string{"Back-end", "Java"}},
		{"Go", 1999, "udemy.com", "def123", []string{"Back-end", "go"}},
		{"Java Script", 399, "udemy.com", "sufi1234", []string{"Front-end", "JavaScript"}},
		{"MERN", 999, "udemy.com", "abc123", []string{"Full stack", "MangoDB", "React.js", "Node.js"}},
		{"MangoDB", 999, "udemy.com", "net34ab", nil},
	}
	// package these data into Jdon
	finalJson, err := json.Marshal(courseDetails)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// The output json is not looking so lets make it beautiful
	finalJson1, _ := json.MarshalIndent(courseDetails, "", "\t")
	fmt.Printf("%s\n", finalJson1)
}
