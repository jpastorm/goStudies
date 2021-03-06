package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	var persons []Person
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	var fileName string
	fmt.Println("Enter a file Name")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		per := Person{firstName: names[0]}
		per.firstName = names[0]
		per.lastName = names[1]
		persons =
			append(persons, per)
		//fmt.Print(persons)
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, each_person := range persons {

		//	structValue, _ := json.Marshal(each_ln)
		fmt.Printf(each_person.firstName + " " + each_person.lastName)
		fmt.Println("")

	}
}
