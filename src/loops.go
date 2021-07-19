package main

import "fmt"

func main() {

	ages := map[strings]int{}
	ages["Kevin"] = 11
	ages["Kaas"] = 28
	ages["Kang"] = 67
	ages["Matty"] = 18
	ages["Kane"] = 16

	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number!", name))
		case 16:
			fmt.Println(name, "Can Drive now! ")
		case 18:
			fmt.Println(name, "Can vote now! ")
		case 67:
			fmt.Println(name, "Can retire now! ")
		default:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age,", name))

		}
	}

}
