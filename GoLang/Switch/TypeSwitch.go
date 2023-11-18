package main

import "fmt"

func typPrint() {
	var value interface{} = "goLang"

	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float type")
	case int:
		fmt.Println("value is of int type")
	case string:
		fmt.Println("value is of string type")
	default:
		fmt.Printf("value is of type %T", q)
	}
}
