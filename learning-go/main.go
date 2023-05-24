package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	conferenceName := "Go Conference"
// 	const conferenceTickets = 50
// 	var remainingTickets uint = 50
// 	// var bookings []string // slices
// 	bookings := []string{} // slices

// 	fmt.Printf("Welcome to %v booking application\n", conferenceName)
// 	fmt.Printf("We have total of %v tickets and %v is still available\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your ticket from here!!!")

// 	for {
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		fmt.Println("Enter your first name")
// 		fmt.Scan(&firstName)

// 		fmt.Println("Enter your last name")
// 		fmt.Scan(&lastName)

// 		fmt.Println("Enter your email address")
// 		fmt.Scan(&email)

// 		fmt.Println("Enter number of tickets you want")
// 		fmt.Scan(&userTickets)

// 		if userTickets > remainingTickets {
// 			fmt.Printf("We only have %v tickets, so you can't book %v tickets\n", remainingTickets, userTickets)
// 			continue
// 		}

// 		remainingTickets = remainingTickets - userTickets
// 		bookings = append(bookings, firstName+" "+lastName)

// 		fmt.Printf("User %v %v buy %v tickets. A confirmation mail will sent to this %v email address\n", firstName, lastName, userTickets, email)
// 		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// 		firstNames := []string{}
// 		for _, booking := range bookings {
// 			var names = strings.Fields(booking)
// 			firstNames = append(firstNames, names[0])
// 		}
// 		fmt.Printf("The first names of bookings are: %v\n", firstNames)

// 		if remainingTickets == 0 {
// 			// end of program
// 			fmt.Printf("Our conference is booked out, come back next year")
// 			break
// 		}

// 	}

// }

import (
	"fmt"
)

type Address struct {
	street string
	city   string
	state  string
	zip    string
}
type Person struct {
	name    string
	age     int
	address Address
}

type myNumber int

func (num myNumber) square() myNumber {
	if num == 0 {
		return 1
	}
	return num * num
}

func plus(a int, b int) int {
	return a + b
}

// Returning multiple value
func vals() (int, string) {
	return 3, "Another variable"
}

type Animal interface {
	walk() string
	bark() string
}

type Dog struct {
	w string
	b string
}

func (d Dog) walk() string {
	return d.w
}

func (d Dog) bark() string {
	return d.b
}

type User struct {
	Name  string
	Email string
}

func (u User) userDetails() string {
	return fmt.Sprintf("User name is :%s and email is: %s", u.Name, u.Email)
}

func (u *User) changeDetails(newName, newEmail string) {
	u.Name = newName
	u.Email = newEmail
}

// Empty Interface
type MyString string

type Rectangular struct {
	width  float64
	height float64
}

func emptyInterface(i interface{}) {
	fmt.Printf("Value given to emptyInterface function is of type '%T' with value %v\n", i, i)
}

func main() {
	defer fmt.Println("Good bye! with defer")
	fmt.Println("Good bye! without defer")
	var_1, var_2 := 1, "hi"
	//declaring for the first time
	fmt.Println("var1, var2", var_1, var_2)
	var_3, var_2 := 2, "hello"
	//same scope, the variable is reassigned
	fmt.Println(var_3, var_2)
	if var_4, var_2 := 3, "hey"; var_4 > var_1 {
		// variable is declared again in the scope of the if condition
		fmt.Println(var_4, var_2)
	}
	fmt.Println(var_2)
	// fmt.Println(var_4) //uncommenting this should give an compilation error
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	val, str := vals()
	fmt.Println(val)
	fmt.Println(str)

	// Anonimous variable
	_, str2 := vals()
	fmt.Println(len(str2))

	fmt.Println("Channel")

	// Channel
	numberCh := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		numberCh <- i
	}
	for j := 1; j <= 10; j++ {
		fmt.Println(<-numberCh)
	}

	person := &Person{
		name: "Bob",
		age:  35,
		address: Address{
			street: "789 Broadway",
			city:   "New York",
			state:  "NY",
			zip:    "10003",
		},
	}

	fmt.Println(person)

	num := myNumber(25)
	sq := num.square()
	fmt.Printf("The square of %d is %d \n", num, sq)
	fmt.Println("")
	fmt.Println("")

	var a Animal = Dog{"Dog is walking.....!!!", "Dog is barking....!!!"}
	fmt.Println("!.....Dog......!")
	fmt.Println(a.walk())
	fmt.Println(a.bark())

	fmt.Println("")
	fmt.Println("")

	user1 := User{Email: "John Doe", Name: "johndoe@Golinuxcloud.com"}
	fmt.Println(user1.userDetails())
	user1.changeDetails("Mary Doe", "marydoe@Golinuxcloud.com")
	fmt.Println(user1.userDetails())
	fmt.Println("")
	fmt.Println("")

	myString := MyString("My name is MD ABU SALMAN HOSSAIN")
	rectangular := Rectangular{5.2, 7}

	emptyInterface(myString)
	emptyInterface(rectangular)

}
