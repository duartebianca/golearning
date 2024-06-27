package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(rand.Float64()))
	fmt.Println("The answer is", math.Pi, "?")
}
