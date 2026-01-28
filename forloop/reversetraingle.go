package forloop

import "fmt"

func LoopReverse() {
	var n int
	fmt.Println("Enter number:")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
