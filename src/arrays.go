package main

import "fmt"

func main() {
	names := make([]string, 4)
	//names := []string{}
	//names = append(names, "Keith")
	//names = append(names, "Keeanu", "Kang", "Kang")

	names[0] = "Keith"
	names[1] = "Keeanu"
	names[2] = "Kang"
	names[3] = "Kaas"
	//names[4] = "Kato"

	fmt.Println(names)
	fmt.Println("names[2] is nill", names[2] == "")

}
