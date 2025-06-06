package bee

import (
	"fmt"
	"math"
)

func Atv1002() {
	var raio float64
	var area float64
	const pi float64 = 3.14159

	fmt.Scanf("%f", &raio)

	area = pi * math.Pow(raio, 2)

	fmt.Printf("A=%.4f\n", area)

}
