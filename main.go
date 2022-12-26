package main

import "fmt"

func main() {
	nome := pegarNome()

	fmt.Println("Seja bem-vindo, ", nome)
	comando := exibirMenu()

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa.")
	} else {
		fmt.Println("Não reconheço esse comando")
	}

}

func pegarNome() string {
	var nome string
	
	fmt.Println("Olá, digite seu nome: ")
	fmt.Scan(&nome)

	limpaTela()

	return nome
}

func limpaTela() {
	fmt.Print("\033[H\033[2J")
}

func exibirMenu() int {
	fmt.Println("Escolha uma opção: ")
	fmt.Println("")
	fmt.Println("|========================|")
	fmt.Println("| 1- Começar monitoração |")
	fmt.Println("| 2- Exibir logs         |")
	fmt.Println("| 0- Sair programa       |")
	fmt.Println("|========================|")
	var comando int
	fmt.Scan(&comando)

	return comando;
}