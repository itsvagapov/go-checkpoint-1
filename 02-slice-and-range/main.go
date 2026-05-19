package main

import "fmt"

func main() {
	names := []string{
		"Иба",
		"Мохьмад",
		"Ислам",
		"Эла",
		"Халид",
	}

	for _, name := range names {
		fmt.Println(name)
	}
}