package main

import "fmt"

func main()  {
	var name1 string = "Elton1"
	var version1 float64 = 1.2
	var age1 int = 32

	var name2 = "Elton2"
	var version2 = 1.2
	var age2 = 32

	name3 := "Elton3"
	version3 := 1.2
	age3 := 32

	imprime(name1, age1, version1)
	imprime(name2, age2, version2)
	imprime(name3, age3, version3)
}

func imprime(name string, age int, version float64) {
	fmt.Println("Hello World in Go, sr.", name, "e sua idade é: ", age)
	fmt.Println("A versão desse software é ", version)
}