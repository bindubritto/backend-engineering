package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func main() {
	fmt.Println("hello world")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var e float32
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("d", d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	// Switch case

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Println(time.Now())

	// Map
	mp := make(map[string]int, 2)
	mp["a"] = 2
	_, prs := mp["a"]
	fmt.Println("mp", prs)

	// Range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

}
