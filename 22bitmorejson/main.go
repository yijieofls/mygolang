package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON!!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"full-stack", "js"}},
		{"PYTHON Bootcamp", 299, "LearnCodeOnline.in", "abc123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"price": 199,
		"website": "LearnCodeOnline.in",
		"tags": ["full-stack","js"]
	}
	`)

	var lcoCourses courses
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("not valid JSON")
	}

	// create without struct, interface make sure value type can be general

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourses)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
