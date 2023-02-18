package main

import "fmt"

var voted = make(map[string]bool)

func checkVoter(name string) {
	if voted[name] {
		fmt.Println("Kick them out!")
	} else {
		voted[name] = true
		fmt.Println("Let them vote!")
	}
}

func main() {
	checkVoter("Ivan")
	checkVoter("Michael")
	checkVoter("Ivan")
}
