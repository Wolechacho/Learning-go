package main

import (
	"fmt"
	"sort"
)

func main() {
	// xi := []int{10, 19, 59, 2, 3, 5, 6, 7, 8}
	// xs := []string{"Zara", "Magi", "James", "John", "Mike", "Ojo"}

	// sort.Ints(xi)
	// fmt.Println(xi)

	// sort.Strings(xs)
	// fmt.Println(xs)

	user1 := user{
		Name: "Steven",
		Age:  30,
		Sayings: []string{
			"You are awesome",
			"Great are great",
		},
	}
	user2 := user{
		Name: "Kedu",
		Age:  40,
		Sayings: []string{
			"This is him",
			"He is an awesome teacher",
		},
	}

	user3 := user{
		Name: "Eke",
		Age:  500,
		Sayings: []string{
			"Joyous man",
			"Incredible",
		},
	}

	users := []user{user1, user2, user3}
	//fmt.Println(users)
	//sort.Sort(ByAge(users))
	//fmt.Println(users)
	sort.Sort(ByName(users))
	fmt.Println(users)


}
