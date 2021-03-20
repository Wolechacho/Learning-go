package regex

import (
	"fmt"
	"regexp"
)

func PrintRegex() {
	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

// 	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
// 	fmt.Println(match)

	 r, _ := regexp.Compile("/api/characters/search\\?name=\\w+")
	 fmt.Println(r.MatchString("/api/characters/search?name=walter"))
// 	fmt.Println(r.MatchString("peach"))

// 	fmt.Println(r.FindString("peach punch"))

// 	fmt.Println(r.FindStringIndex("peach punch"))

// 	fmt.Println(r.FindStringSubmatch("peach punch"))

// 	//match all
// 	fmt.Println(r.FindAllString("peach punch pinch", -1))

// 	//match two
// 	fmt.Println(r.FindAllString("peach punch pinch", 2))

// 	//match the byte of a string
// 	fmt.Println(r.Match([]byte("peach")))

// 	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
}
