package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
