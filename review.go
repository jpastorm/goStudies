package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
	return
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
	return
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
	return
}

func main() {

	var inputType, inputAnimal, inputInfo, inputValue string

	var common AnimalInterface

	am := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Println("Enter selection type followed by the name of the animal followed by your requested information (for eg: query cow eat )")
		fmt.Print(">")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputValue = scanner.Text()

		inputType = strings.Split(inputValue, " ")[0]
		inputAnimal = strings.Split(inputValue, " ")[1]
		inputInfo = strings.Split(inputValue, " ")[2]

		if inputType == "query" {
			common = am[inputAnimal]
			if inputInfo == "eat" {
				common.Eat()
			} else if inputInfo == "move" {
				common.Move()
			} else if inputInfo == "speak" {
				common.Speak()
			} else {
				fmt.Println("Invalid data input")
			}
		} else if inputType == "newanimal" {
			am[inputAnimal] = am[inputInfo]
			fmt.Println("Created It!")
		} else {
			fmt.Println("Invalid Data")
		}

	}

}
