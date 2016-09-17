package main

import (
	"fmt"
	"sync")

const (
	goRutinas = 20
	capacidad = 4)

var wg sync.WaitGroup

var recursos = make(chan string, capacidad)

func main(){

	wg.Add(goRutinas)

	for gr :=1;gr<=goRutinas;gr++ {
		go func(gr int) {
			worker(gr)
			wg.Done()
		}(gr)
	}


	for s := 'A'; s < 'A' + capacidad; s++ {
		recursos <- string(s)
	}
	wg.Wait()
	}
func worker(worker int){

	valor := <-recursos

	fmt.Printf("Worker:%d : %s\n", worker,valor)

	recursos <- valor
}


