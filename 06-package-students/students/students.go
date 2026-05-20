package students

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