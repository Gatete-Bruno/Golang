package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 77

	if ages["Kevin"] < 18 {
		fmt.Println("Kevin can not Vote")

	} else if ages["Kevin"] < 67 {

		fmt.Println("Kevin is not yet retired")

	} else {
		fmt.Println("Enjoy your retirement Kev")
	}
}
