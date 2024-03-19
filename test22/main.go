package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(9223372036854775807)
	y := big.NewInt(2223371036834773801)

	addition := new(big.Int).Add(x, y)
	substraction := new(big.Int).Sub(x, y)
	mulitplication := new(big.Int).Mul(x, y)
	divsion := new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt(y))

	fmt.Printf("x = %d\ny = %d\n", x, y)
	fmt.Printf("x + y = %d\n", addition)
	fmt.Printf("x - y = %d\n", substraction)
	fmt.Printf("x * y = %d\n", mulitplication)
	fmt.Printf("x / y = %f\n", divsion)
}
