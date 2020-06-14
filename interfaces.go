package main

import "fmt"

type Person interface {
	getAge() int
}

type Parent struct {
	name string
	born int
}

type Son struct {
	name string
	born int
}

func (son Son) getAge() int {
	return 2020 - son.born
}

func (parent Parent) getAge() int {
	return 2020 - parent.born
}

func printAge(p Person) {
	fmt.Println(p.getAge())
}

func main() {
	father := Parent{"Pai", 1950}
	son := Son{"Filho", 1990}

	fmt.Println(father)
	fmt.Println(son)

	printAge(father)
	printAge(son)

}
