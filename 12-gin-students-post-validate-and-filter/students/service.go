package students

import (
	"errors"
	"sort"
)

func GetAllStudents(filter Filter) []Student {
	readyStudents := []Student{}

	for _, elem := range Students {

		// фильтр по городу
		if filter.City != nil && elem.City != *filter.City {
			continue
		}

		// фильтр по минимальному возрасту
		if filter.MinAge != nil && elem.Age < *filter.MinAge {
			continue
		}

		// фильтр по максимальному возрасту
		if filter.MaxAge != nil && elem.Age > *filter.MaxAge {
			continue
		}

		readyStudents = append(readyStudents, elem)
	}

	// сортировка
	if filter.Sort != nil {

		if *filter.Sort == "age" {
			sort.Slice(readyStudents, func(i, j int) bool {
				return readyStudents[i].Age < readyStudents[j].Age
			})
		}

		if *filter.Sort == "full_name" {
			sort.Slice(readyStudents, func(i, j int) bool {
				return readyStudents[i].FirstName < readyStudents[j].FirstName
			})
		}
	}

	// offset
	start := 0

	if filter.Offset != nil {
		start = *filter.Offset
	}

	if start > len(readyStudents) {
		return []Student{}
	}

	// limit
	end := len(readyStudents)

	if filter.Limit != nil {
		end = start + *filter.Limit

		if end > len(readyStudents) {
			end = len(readyStudents)
		}
	}

	return readyStudents[start:end]
}

func AddStudent(student Student) {
	Students = append(Students, student)
}

func ValidateStudent(student Student) error {

	if student.FirstName == "" {
		return errors.New("поле first_name обязательно")
	}

	if student.LastName == "" {
		return errors.New("поле last_name обязательно")
	}

	if student.City == "" {
		return errors.New("поле city обязательно")
	}

	if student.Age < 16 || student.Age > 100 {
		return errors.New("age должен быть в диапазоне 16..100")
	}

	return nil
}