package main

import "fmt"

func main() {
	fmt.Println("map")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages ", languages)
	fmt.Println("JS: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loop in map

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
