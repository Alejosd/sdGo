package main

import (
	"fmt"
	"runtime"
	"sync")

func init(){
	//permite reservar un procesador logico para que lo use el scheduler
	runtime.GOMAXPROCS(1)
}

func main()  {

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Iniciar Goroutines")
	// Declarar funcion anonima

	go func(){
		for contador := 100;contador>=0;contador--{
			fmt.Printf("[A:%d]\n",contador)
		}
		wg.Done()
	}()

	go func(){
		for contador := 0;contador<100;contador++{
			fmt.Printf("[A:%d]\n",contador)
		}
		wg.Done()
	}()

	fmt.Println("Esperando a Goroutines")

	wg.Wait()

	fmt.Println("finalizar el programa")
}
