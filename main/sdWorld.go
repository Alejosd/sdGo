package main

import "fmt"
import "math"

func main() {

	var helloWorld string = "helloworld"
	fmt.Println("Rock And Roll "+helloWorld)

	var a,b float64 = 8.5,25.0
	fmt.Println(a/b)

	var t = true
	var f = false

	fmt.Println(t&&f)
	fmt.Println(t ||f)
	fmt.Println(!f)

	//declaracion variable
	var entero int
	fmt.Println(entero)

	//declaracion corta

	z:= "short"
	fmt.Println(z)

	//constantes

	const n = 500
	const d = 3e20 / n

	fmt.Println(n)

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}