// Добрий ранок!
// Домашнє завдання по інтерфейсах:
// кожна тварина в залежності від свого типу споживає різку кількість кілограмів їжі на місяць

// собака - їсть 10кг корму на кожні 5кг власної ваги
// кішка - 7кг на кожен кілограм власної ваги
// корова - 25кг на кожен кілограм власної ваги

// на фермі може бути різна кількість собак, кішок та корів
// кажен тип тварини має сам розраховувати для себе вагу корму
// написати динамічну функцію, яка буде виводити в консоль для кожної тварини на фермі її назву, вагу, та скільки їжі треба для того щоб її прогодувати
// ця функція також моє повертати сумму кг корму на всю ферму для подальшого виводу у консоль
// Важливо: в окремий тред необхідно скинути посилання на пул реквест на github.
// Ревью буде проходити у тому вигляді, в якому це буде вітбуватися на реальній роботі.
// Задачі перевірятимемо до четверга, оскільки далі буде вже нове завдання.
// мій github https://github.com/DGoldiner
// Всі питання по домашньому завданню у цей тред:



package main

import (
	"fmt"
)

func main() {
	sliceOfAnimals := [] animal {cow{name: "Cow", weight:136}, cat{name: "Cat", weight: 7}, dog{name: "Dog", weight:15}}
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
	return d.weight * 10.0
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


func feedIntakeforOneAnimal(a animal) {
	fmt.Printf("Our %v is %v kg and consumes %v kg of feed per month\n", a.getName(), a.getWeight(), a.feedIntake())
}

func totalFeedIntake(animals []animal) {
	var total float32
	for _, a := range animals {
		feedIntakeforOneAnimal(a)
		total = total + a.feedIntake()
	}
	fmt.Printf("The total amound of feed for all the animals is %v kg\n", total)
}


