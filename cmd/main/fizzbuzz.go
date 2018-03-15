package main

import "fmt"

func main()  {

	for i := 1; i <= 100; i++ {
		var three, five = false, false
		if i % 3 == 0 {
			fmt.Print("Fizz")
			three = true
		}
		if i % 5 == 0 {
			fmt.Print("Buzz")
			five = true
		}
		if !three && !five {
			fmt.Print(i)
		}
		fmt.Println()
	}

}
