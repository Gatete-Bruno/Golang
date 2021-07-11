package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Keith": "02/06/1990",
		"Kang":  "01/01/1957",
		"Kaas":  "11/04/1987",
	}

	delete(birthdays, "Keith")

	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Kang"] = 62
	ages["Kaas"] = 78

	fmt.Println(ages, ages["Kevin"])

}
