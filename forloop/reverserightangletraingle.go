package forloop

import "fmt"

func ReverseRightAngleTraingle() {
	var n int
	fmt.Println("Enter number:")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")

		}
		fmt.Println()

	}
}
