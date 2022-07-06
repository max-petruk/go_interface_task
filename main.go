package main

import (
	"fmt"
)

func main() {
	sliceOfAnimals := [] animal {
		cow{name: "Cow", weight:136},
		cat{name: "Cat", weight: 7},
		dog{name: "Dog", weight:15},
	}
	totalFeedIntake(sliceOfAnimals)
}


type animal interface {
	getName() string
	feedIntake() float32
	getWeight() float32
}

type cow struct {
	weight float32
	name string
}

type cat struct {
	weight float32
	name string
}

type dog struct {
	weight float32
	name string
}

// returning the name of an animal

func (c cow) getName() string {
	return c.name
}

func (c cat) getName() string {
	return c.name
}

func (d dog) getName() string {
	return d.name
}

// returning the feed Intake of an animal

func (c cow) feedIntake() float32 {
	return c.weight * 25.0
}

func (c cat) feedIntake() float32 {
	return c.weight * 7.0
}

func (d dog) feedIntake() float32 {
	return d.weight * 2.0
}

// returning the weight of an animal

func (c cow) getWeight() float32 {
	return c.weight
}

func (c cat) getWeight() float32 {
	return c.weight
}

func (d dog) getWeight() float32 {
	return d.weight
}


func printFeedIntake(a animal) {
	fmt.Printf("Our %v is %v kg and consumes %v kg of feed per month\n", a.getName(), a.getWeight(), a.feedIntake())
}

func totalFeedIntake(animals []animal) {
	var total float32
	for _, a := range animals {
		printFeedIntake(a)
		total = total + a.feedIntake()
	}
	fmt.Printf("The total amound of feed for all the animals is %v kg\n", total)
}


