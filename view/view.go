package view

import (
	"fmt"
)

func ExibeIntroducao() {
	var nome string

	fmt.Println("Olá, digite seu nome: ")
	fmt.Scan(&nome)

	limpaTela()

	fmt.Println("Seja bem-vindo, ", nome)
}

func limpaTela() {
	fmt.Print("\033[H\033[2J")
}

func ExibirMenu() int {
	fmt.Println("Escolha uma opção: ")
	fmt.Println("")
	fmt.Println("|========================|")
	fmt.Println("| 1- Começar monitoração |")
	fmt.Println("| 2- Exibir logs         |")
	fmt.Println("| 0- Sair programa       |")
	fmt.Println("|========================|")

	var comando int
	fmt.Scan(&comando)

	return comando
}

func Imprime(arquivos []string) {
	for _, log := range arquivos {
		fmt.Println(log)
	}
}
