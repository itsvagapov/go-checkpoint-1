package main

import (
	"fmt"
	"students/students"
)

func main() {
	for _, s := range students.Students {
		fmt.Printf("%s %s - %s\n", s.FirstName, s.LastName, s.Status())
	}
}