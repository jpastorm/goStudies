// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type animal struct {
// 	food, locomotion, sound string
// }

// func (c *animal) eat() {
// 	fmt.Println(c.food)
// }

// func (c *animal) move() {
// 	fmt.Println(c.locomotion)
// }

// func (c *animal) speak() {
// 	fmt.Println(c.sound)
// }

// func main() {
// 	var cow, bird, snake animal
// 	var output animal
// 	cow.food = "grass"
// 	cow.locomotion = "walk"
// 	cow.sound = "moo"

// 	bird.food = "worms"
// 	bird.locomotion = "fly"
// 	bird.sound = "peep"

// 	snake.food = "mice"
// 	snake.locomotion = "slither"
// 	snake.sound = "hsss"

// 	var input string
// 	fmt.Println("Enter the name ")
// 	fmt.Println(">")
// 	in := bufio.NewReader(os.Stdin)

// 	input, err := in.ReadString('\n')
// 	if err != nil {
// 		return
// 	}
// 	inputs := strings.Split(input, " ")
// 	switch inputs[0] {
// 	case "cow":
// 		output = cow

// 	case "bird":
// 		output = bird
// 	case "snake":
// 		output = snake
// 	}
// 	inputs[1] = strings.TrimSuffix(inputs[1], "\n")
// 	switch inputs[1] {
// 	case "eat":
// 		output.eat()
// 	case "move":
// 		output.move()
// 	case "speak":
// 		output.speak()
// 	}
// }

// // package main

// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"os"
// // 	"strings"
// // )

// // type Animal struct {
// // 	food       string
// // 	locomotion string
// // 	noise      string
// // }

// // func (a Animal) Eat() {
// // 	fmt.Println(a.food)
// // }

// // func (a Animal) Move() {
// // 	fmt.Println(a.locomotion)
// // }

// // func (a Animal) Speak() {
// // 	fmt.Println(a.noise)
// // }

// // func main() {
// // 	cow := Animal{"grass", "walk", "moo"}
// // 	bird := Animal{"worms", "fly", "peep"}
// // 	snake := Animal{"mice", "slither", "hsss"}

// // 	fmt.Println("Enter your request for information about the folowing animals:")
// // 	fmt.Println("cow, bird, snake.")
// // 	fmt.Println("You can request information about what they eat, how they move, and what they speak.")
// // 	fmt.Println("They can be accessed with the following keywords:")
// // 	fmt.Println("eat, move, speak.")
// // 	fmt.Println("Enter your request in the format: animal request")
// // 	fmt.Println("For example:")
// // 	fmt.Println("cow speak")
// // 	fmt.Println("press ctl-c to exit")
// // 	fmt.Println("====================================================================================")

// // 	for {
// // 		fmt.Print("> ")

// // 		in := bufio.NewReader(os.Stdin)
// // 		line, _ := in.ReadString('\n')
// // 		inputs := strings.Fields(line)
// // 		var animal Animal

// // 		if inputs[0] == "cow" {
// // 			animal = cow
// // 		} else if inputs[0] == "bird" {
// // 			animal = bird
// // 		} else if inputs[0] == "snake" {
// // 			animal = snake
// // 		}

// // 		if inputs[1] == "eat" {
// // 			animal.Eat()
// // 		} else if inputs[1] == "move" {
// // 			animal.Move()
// // 		} else if inputs[1] == "speak" {
// // 			animal.Speak()
// // 		}
// // 	}

// // }
