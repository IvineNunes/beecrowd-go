package bee

//como fiz:
// import "fmt"

// func Atv1001() {
// 	var A int
// 	var B int

// 	fmt.Print("")
// 	fmt.Scan(&A)
// 	fmt.Print("")
// 	fmt.Scan(&B)

// 	var X = A + B

// 	fmt.Println("X = \n", X)
// }

//Beecrowd:

import "fmt"

func Atv1001() {
	var a, b int
	var x int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	x = a + b

	fmt.Printf("X = %d\n", x)

}
