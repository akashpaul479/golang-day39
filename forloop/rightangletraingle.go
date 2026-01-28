package forloop

import "fmt"

func Rightangletraingle() {
	var n int
	fmt.Println("Enter numbers:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-1; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
