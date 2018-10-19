// Package interface provides a shell
// to prompt the user to create an animal or query for an animal
package main

import (
	"errors"
	"fmt"
)

var animals = make(map[string]Animal)
var usage = "usage: newanimal <name> <arg> | query <name> <arg> | exit"

func main() {
	var cmd command
	var name, arg string
	fmt.Println(usage)
	for cmd != "exit" {
		fmt.Print("> ")
		n, _ := fmt.Scanln(&cmd, &name, &arg)
		if n < 3 {
			if cmd != "exit" {
				fmt.Println("not enough elements")
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
		err = errors.New("unknown method")
	}
	return
}

func createAnimal(name, kind string) (err error) {
	switch kind {
	case "cow":
		animals[name] = Cow{"grass", "walk", "moo"}
	case "bird":
		animals[name] = Bird{"worms", "fly", "peep"}
	case "snake":
		animals[name] = Snake{"mice", "slither", "hsss"}
	default:
		err = errors.New("unknown animal")
	}
	return
}

// Cow is an Animal
type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}
