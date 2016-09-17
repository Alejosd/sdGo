package main

import "fmt"

func main(){

	numeros := []int{2,4,6}
	suma := 0
	for _, num:= range numeros {
		suma += num
	}

	fmt.Println("suma:", suma)

	for i, numero :=range numeros{
		if numero == 3 {
			fmt.Println("index:", i)
		}
	}


	mapa := map[string]string{"a":"auto","b":"bebe"}
	for key, value:=range mapa{
		fmt.Println(key,value)
	}
}


