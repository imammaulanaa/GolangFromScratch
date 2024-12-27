package main

import "fmt"

// as parameter
type Filter func(string) string

// anonymous
type Blacklist func(string) bool

func sayHi() {
	fmt.Println("Hi!")
}

// parameter
func sayHiTo(firstName string, lastName string) {
	fmt.Println("Hi", firstName, lastName)
}

// return value
func getHi(name string) string {
	hi := "Hi " + name
	return hi
}

// return multiple values
func getFullName() (string, string) {
	return "Imam", "Maulana"
}

// name return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Imam"
	middleName = "Maulana"
	lastName = "Akbar"

	return firstName, middleName, lastName
}

// variadic
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// as parameter
func sayHiWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hi", filteredName)
}

func spamFilter(name string) string {
	if name == "Imam" {
		return "..."
	} else {
		return name
	}
}

// anonymous
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	sayHi()

	// parameter
	sayHiTo("Imam", "Maulana")

	// return value
	result := getHi("Imam")
	fmt.Println(result)
	fmt.Println(getHi("Akbar"))

	// return multiple values
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	//firstName, _ := getFullName()
	//fmt.Println(firstName)

	// name return value
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	// variadic
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

	// value
	sample := getGoodBye
	second := getGoodBye

	fmt.Println(sample("Imam"))
	fmt.Println(second("Maulana"))

	// as parameter
	sayHiWithFilter("Imam", spamFilter)

	filter := spamFilter
	sayHiWithFilter("Akbar", filter)

	// anonymous
	blacklist := func(name string) bool {
		return name == "akbar"
	}
	registerUser("imam", blacklist)

	registerUser("akbar", func(name string) bool {
		return name == "akbar"
	})

	// recursive
	hasil := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(hasil)
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
