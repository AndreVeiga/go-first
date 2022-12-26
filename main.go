package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5 * time.Second
const nomeArquivo = "sites.txt"

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

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open(nomeArquivo)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func iniciarMonitoracao() {
	fmt.Println("Monitorando...")

	sites := leArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}

		fmt.Println("")
		time.Sleep(delay)
	}
}

func testaSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

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
