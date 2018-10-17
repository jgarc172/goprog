// Package animals provides a shell to enter animal commands
// which are intepreted to their internal representation, and the command is invoked
package main

import (
	"fmt"
)

var animals map[string]Animal

func init() {
	animals = make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}
}

func main() {
	var name, method string

	fmt.Println("Usage: animal method | exit")
	for name != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&name, &method)
		if n < 2 {
			continue
		}
		if animal, ok := animals[name]; ok {
			fmt.Println(animal.invoke(method))
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func (a Animal) invoke(method string) string {
	switch method {
	case "eat":
		return a.Eat()
	case "move":
		return a.Move()
	case "speak":
		return a.Speak()
	default:
		return "unknown method"
	}
}
