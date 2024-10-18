package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"studentmanagement/utility"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}

func getNumberInput(prompt string) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)
		if err == nil {
			return number
		}
		fmt.Println("Not a valid number!")
	}
}

func getOptInput(opts []rune, prompt string) rune {
	for {
		fmt.Print(prompt)
		input, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		_, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		for _, option := range opts {
			if option == input {
				return input
			}
		}
	}
}

func loop() {
	options := []rune{'a', 'b', 'c', 'd'}
	option := getOptInput(options, "Choose a option(a: To add a student, b: To update a student, c: To list all students, d: Delete student)")

	switch option {
	case 'a':
		name := getInput("Name: ")
		age := getNumberInput("Age: ")
		email := getInput("Email: ")
		student := utility.Student{Name: name, Email: email, Age: int8(age)}
		utility.AddStudent(student)
	case 'c':
		students := utility.GetAllStudent()
		fmt.Println(students)
	case 'd':
		email := getInput("Email: ")
		utility.DeleteStudent(email)
	default:
		break
	}
}

func main() {
	loop()
}
