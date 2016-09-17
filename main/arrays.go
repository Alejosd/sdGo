package main

import "fmt"

func main(){

	var nombres[5] string

	amigos := [5]string{"alejo","cris","juan","pipe","pedro"}

	nombres = amigos

	fmt.Println(nombres)

	var numeros []int

	for i:=0; i< 10 ;i++ {
		numeros = append(numeros,i*10)
	}

	fmt.Println(numeros)

	for _, numero:= range numeros{

		fmt.Println(numero)
	}

	frutas := []string{"manzana","naranja","pera","sandía","aguacate"}

	for i, fruta:= range frutas {
		fmt.Printf("Index: %d Fruta: %s\n",i,fruta)
	}

	slice := frutas [0:1]

	for i, fruta:=range slice {
		fmt.Printf("Index: %d Fruta: %s\n",i,fruta)
	}

}
