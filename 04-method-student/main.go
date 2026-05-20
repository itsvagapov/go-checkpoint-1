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

func (s Student) Status() string {
	if s.Age > 18 {
		return "Студент взрослый"
	}

	return "Программирование состарит тебя"
}

func main() {
	s := Student {
		FirstName: "Ибрях1им",
		LastName: "Тсуев",
		Age: 17,
		City: "Grozny",
		Course: "Go",
		IsMale: true,
		HeightMeters: 1.82,
		WeightKg: 69.8,
		ITExperience: false,
		HasCar: false,
	}

	fmt.Println(s.Status())
}