package main


import "fmt"

func main()  {
	var name string = "Elton1"
	var version float64 = 1.2

	fmt.Println("Olá, digite seu nome: ", name)
	fmt.Println("A versão desse software é ", version)
	fmt.Println("Pressione ENTER para começar")
	var nada int
	fmt.Scanf("%d", &nada)

	exibirMenu()
}

func exibirMenu() int {
	fmt.Print("\033[H\033[2J")
	fmt.Println("|========================|")
	fmt.Println("| 1- Começar monitoração |")
	fmt.Println("| 2- Exibir logs         |")
	fmt.Println("| 0- Sair programa       |")
	fmt.Println("|========================|")
	var comando int
	fmt.Scan(&comando)

	return comando;
}