package main

import "fmt"

func main(){

	departamentos:= make(map[string]int)

	departamentos["Devs"] = 25
	departamentos["QA"] = 10
	departamentos["Ops"] = 10

	fmt.Println(departamentos["Devs"])
	fmt.Println(departamentos["QA"])
	fmt.Println(departamentos["Ops"])

	for key, value := range departamentos{
		fmt.Printf("Dept: %s Personas: %d\n",key,value)
	}


}
