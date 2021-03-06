package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
In the main function we have initialised a wait group with a value of 2 to wait for bot the print function go routines to finish executing to see the
output they have produced and in whoch order they are produced.
Since the programs are invoked as goroutines, the execution of the functions will take place in an interleaved manner in the OS thread which is managed
by the go scheduler. The sceheduling takes place non deterministically , so we cannot always correctly say whoch function will get printed
first
While we are  running we see that the order does change in between many runs showing that there is a race between the go rouitnes for resources
*/
func main() {
	wg.Add(2)
	go printHello1()
	go printHello2()
	wg.Wait()

}

/*
We are defining two functions printHello1 and printHello2 to check which executes first
*/
func printHello1() {
	fmt.Println("Helloooo 1")
	wg.Done()
}

func printHello2() {
	fmt.Println("Helloooo 2")
	wg.Done()
}
