package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		comando := exibirMenu()

		switch comando {
		case 1:
			iniciarMonitoracao()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço esse comando")
			os.Exit(-1)
		}
	}
}

func iniciarMonitoracao() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"

	res, _ := http.Get(site)

	fmt.Println("O site", site, "está com statusCode:", res.StatusCode)
}

func exibeIntroducao() {
	var nome string

	fmt.Println("Olá, digite seu nome: ")
	fmt.Scan(&nome)

	limpaTela()

	fmt.Println("Seja bem-vindo, ", nome)
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

	return comando
}
