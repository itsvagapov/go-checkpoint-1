package main

import "fmt"

type Student struct {
	FirstName string
	LastName string
	Age int
	City string
	Course string
	IsMale bool
	HeightMeters float64
	WeightKg float64
	ITExperience bool
	HasCar bool
}

func main() {
	student := Student{
		FirstName: "Ислам",
		LastName: "",
		Age: 32,
		City: "Grozny",
		Course: "Go-backend",
		IsMale: true,
		HeightMeters: 1.80,
		WeightKg: 67.22,
		ITExperience: false,
		HasCar: true,
	}

	fmt.Printf("%#v\n", student)
}