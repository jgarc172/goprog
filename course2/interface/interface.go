// Package interface provides a shell
// to prompt the user to create an animal or query for an animal
package main

import (
	"errors"
	"fmt"
)

var animals = make(map[string]Animal)
var usage = "usage: newanimal <name> <animal> | query <name> <action> | exit"

func main() {
	var cmd command
	var name, arg string
	fmt.Println(usage)
	for cmd != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&cmd, &name, &arg)
		if n < 3 {
			if cmd != "exit" {
				fmt.Println("not enough arguments")
				fmt.Println(usage)
			}
			continue
		}
		err := cmd.invoke(name, arg)
		if err != nil {
			fmt.Println(err)
			fmt.Println(usage)
		}
	}
}

type command string

func (c command) invoke(name, arg string) (err error) {
	switch c {
	case "query":
		if animal, ok := animals[name]; ok {
			err = invokeMethod(animal, arg)
		} else {
			err = errors.New(name + " not found")
		}
	case "newanimal":
		err = createAnimal(name, arg)
	default:
		err = errors.New("unknown command")
	}
	return
}

// Animal eats, moves or speaks
type Animal interface {
	Eat()
	Move()
	Speak()
}

func invokeMethod(animal Animal, method string) (err error) {
	switch method {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		err = errors.New(method + " not in actions: eat, move, speak")
	}
	return
}

// baseAnimal is an Animal
type baseAnimal struct {
	food       string
	locomotion string
	noise      string
}

func (b baseAnimal) Eat() {
	fmt.Println(b.food)
}
func (b baseAnimal) Move() {
	fmt.Println(b.locomotion)
}

func (b baseAnimal) Speak() {
	fmt.Println(b.noise)
}

// Cow is an Animal
type Cow struct {
	baseAnimal // struct embedding
}

func newCow() Cow {
	return Cow{baseAnimal{"grass", "walk", "moo"}}
}

// Bird is an Animal
type Bird struct {
	baseAnimal // struct embedding
}

func newBird() Bird {
	return Bird{baseAnimal{"worms", "fly", "peep"}}
}

// Snake is an Animal
type Snake struct {
	baseAnimal // struct embedding
}

func newSnake() Snake {
	return Snake{baseAnimal{"mice", "slither", "hsss"}}
}

func createAnimal(name, kind string) (err error) {
	switch kind {
	case "cow":
		animals[name] = newCow()
	case "bird":
		animals[name] = newBird()
	case "snake":
		animals[name] = newSnake()
	default:
		err = errors.New(kind + " not in animals: cow, bird, snake")
	}
	return
}
