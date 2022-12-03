package main

import "fmt"

// maxLen defines the maximum length of the Fibo series returned from the function
const maxLen int = 10

func main() {

	greet := greeting("Narinder")

	fmt.Println(greet)

	retArr, err := printFiboSeries(8)
	if err != nil {
		fmt.Println("Error Occured: ", err)
	} else {
		fmt.Println(retArr)
	}

}

func greeting(s string) string {

	greet := fmt.Sprintf("Hello %s", s)
	return greet
}

func printFiboSeries(num int) ([]int, error) {
	var fiboArr [maxLen]int

	if num > maxLen {
		return fiboArr[:], fmt.Errorf("only a maximum of %d numbers can be printed", maxLen)
	} else if num < 1 {
		return fiboArr[:], fmt.Errorf("max. count cannot be zero")
	}

	if num == 1 {
		fiboArr[0] = 1
	}

	var first, second int = 1, 1

	for i := 0; i < num; i++ {
		fiboArr[i] = first
		first = second
		second = fiboArr[i] + first
	}

	return fiboArr[:num], nil
}
