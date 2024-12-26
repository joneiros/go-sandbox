package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fbCountStr := os.Args[1]
	fbCount, err := strconv.Atoi(fbCountStr)
	if err != nil {
		fmt.Println("Error Reached")
		panic(err)
	}
	fmt.Println("You want to FizzBuzz to " + fbCountStr + ".")

	for i := 1; i <= fbCount; i++ {
		strf := fizz(i)
		strb := buzz(i)

		if strf != "" || strb != "" {
			fmt.Println(strf + strb + "!")
		} else {
			fmt.Println(strconv.Itoa(i))
		}
	}

}

func fizz(fizzable int) string {
	if fizzable%3 == 0 {
		return "Fizz"
	}

	return ""
}

func buzz(buzzable int) string {
	if buzzable%5 == 0 {
		return "Buzz"
	}

	return ""
}
