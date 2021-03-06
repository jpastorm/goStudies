package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	eat()
	move()
	speak()
}

type cow struct {
	name       string
	food       string `default:"grass"`
	locomotion string `default:"walk"`
	sound      string `default:"moo"`
}

type bird struct {
	name       string
	food       string `default:"worms"`
	locomotion string `default:"fly"`
	sound      string `default:"peep"`
}

type snake struct {
	name       string
	food       string `default:"mice"`
	locomotion string `default:"slither"`
	sound      string `default:"hsss"`
}

func (c cow) eat() {
	fmt.Println(c.food)
}

func (c cow) move() {
	fmt.Println(c.locomotion)
}

func (c cow) speak() {
	fmt.Println(c.sound)
}
func (c bird) eat() {
	fmt.Println(c.food)
}

func (c bird) move() {
	fmt.Println(c.locomotion)
}

func (c bird) speak() {
	fmt.Println(c.sound)
}
func (c snake) eat() {
	fmt.Println(c.food)
}

func (c snake) move() {
	fmt.Println(c.locomotion)
}

func (c snake) speak() {
	fmt.Println(c.sound)
}
func main() {
	var cows []cow
	var birds []bird
	var snakes []snake

	for {
		fmt.Println("\n Enter anyone of the following command \n\n  newanimal - for creating new animal (eg: newanimal cow cowname) \n  query - for printing animal info (eg: query animalName eat)  \n  Exit - for exiting")
		fmt.Print("\n> ")
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		if err != nil {
			return
		}
		inputs := strings.Split(input, " ")
		if len(inputs) > 1 {
			inputs[2] = strings.TrimSuffix(inputs[2], "\n")
		} else {
			inputs[0] = strings.TrimSuffix(inputs[0], "\n")
		}
		switch inputs[0] {
		case "Exit":
			return
		case "newanimal":
			fmt.Println("")
			switch inputs[2] {
			case "cow":
				var newCow cow = cow{food: "grass", locomotion: "walk", sound: "moo"}
				newCow.name = inputs[1]
				cows = append(cows, newCow)
				fmt.Printf("%v", cows[0].name)
			case "bird":
				var newBird bird = bird{food: "worms", locomotion: "fly", sound: "peep"}
				newBird.name = inputs[1]
				birds = append(birds, newBird)
				fmt.Printf("%v", birds[0].name)
			case "snake":
				var newSnake snake = snake{food: "mice", locomotion: "slither", sound: "hsss"}
				newSnake.name = inputs[1]
				snakes = append(snakes, newSnake)
				fmt.Printf("%v", snakes[0].name)

			}
		case "query":
			fmt.Println("")
			var newAnimal animal
			for _, value := range cows {
				if value.name == inputs[1] {
					newAnimal = value
				}
			}
			for _, value := range birds {
				if value.name == inputs[1] {
					newAnimal = value
				}
			}
			for _, value := range snakes {
				if value.name == inputs[1] {
					newAnimal = value
				}
			}

			switch inputs[2] {
			case "eat":
				newAnimal.eat()
			case "move":
				newAnimal.move()
			case "snake":
				newAnimal.speak()

			}

		}
	}

}
