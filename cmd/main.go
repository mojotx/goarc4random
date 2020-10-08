package main

import "fmt"
import arc4 "github.com/mojotx/goarc4random/pkg/arc4random"

func main() {
	const (
		iterations = 100
		upperBound = 10
	)

	fmt.Println("Howdy!")
	fmt.Printf("A random number is %v\n\n", arc4.Arc4Random())
	fmt.Println("Generating ten numbers between 1-10")

	count := make([]uint, upperBound+1, upperBound+1)

	for i := 0; i < iterations; i++ {
		n := arc4.Arc4RandomUniform(upperBound) + 1
		save := count[n]
		count[n]++
		fmt.Printf("count[%d] was %d, now %d\n", n, save, count[n])
		fmt.Printf("%d: %d\n", i, n)
	}

	fmt.Println("\n\x1B[36mDumping...\x1B[0m")
	fmt.Println()

	sum := uint(0)
	for i, e := range count {
		if e > 0 {
			sum += e
			fmt.Printf("%d: %d\n", i, e)
		}
	}
	fmt.Printf("Sum = %d\n", sum)
}
