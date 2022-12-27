package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5 * time.Second
const arquivoSites = "sites.txt"
const arquivosLogs = "log.txt"

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

func registraLog(site string, status int) {
	arquivo, err := os.OpenFile(arquivosLogs, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err == nil {
		log1 := time.Now().Format("02/01/2006 15:04:05") + " " + site
		log2 := " com status:" + strconv.FormatInt(int64(status), 10)

		arquivo.WriteString(log1 + log2 + "\n")
		
		arquivo.Close()
	} else {
		fmt.Println("Arquivo " + arquivosLogs + "com problemas.")
	}
}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open(arquivoSites)

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
	registraLog(site, res.StatusCode)
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
