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
	students := []Student {
		{
		FirstName: "Муслим",
		LastName: "Анзоров",
		Age: 21,
		City: "Лос-Анджелес",
		Course: "JS-frontend",
		IsMale: true,
		HeightMeters: 1.81,
		WeightKg: 82.65,
		ITExperience: false,
		HasCar: true,
	},
	{
		FirstName: "Ахьмад",
		LastName: "Сулейманов",
		Age: 21,
		City: "Грозный",
		Course: "PY-backend",
		IsMale: true,
		HeightMeters: 1.75,
		WeightKg: 82.65,
		ITExperience: false,
		HasCar: false,
		},
		{
		FirstName: "Арби",
		LastName: "Батукаев",
		Age: 26,
		City: "Грозный",
		Course: "PY-backend",
		IsMale: true,
		HeightMeters: 1.87,
		WeightKg: 90.10,
		ITExperience: true,
		HasCar: false,	
		},
	}

	for _, s := range students {
		fmt.Printf("%s %s - %s\n", s.FirstName, s.LastName, s.Status())
	}
}