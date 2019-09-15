package main

import (
	"fmt"
)

func main() {
	// println
	fmt.Println("hello 雷大")
	// array
	names := []string{"it", "home"}
	fmt.Println(names)
	// slice
	game := "it home iron man"
	fmt.Println(game[8:])
	// switch case
	gender := "female"
	switch gender {
	case "female":
		fmt.Println("you are a girl")
	case "male":
		fmt.Println("you are a boy")
	default:
		fmt.Println("you are intersex")
	}

	// key value pairs
	kvs := map[string]string{
		"name":    "json",
		"website": "https://yuanyu90221.github.io",
	}

	for key, value := range kvs {
		fmt.Println(key, value)
	}

	type Post struct {
		ID         int
		Title      string
		Author     string
		Difficulty string
	}

	p := Post{
		ID:         1011111,
		Title:      "下班不講幹話多學點go",
		Author:     "json",
		Difficulty: "Beginner",
	}

	fmt.Println(p)
}
