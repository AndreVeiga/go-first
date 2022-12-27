package repository

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const arquivosLogs = "logs.txt"
const arquivoSites = "sites.txt"

func LeSites() []string {
	return leArquivo(arquivoSites)
}

func LeLogs() []string {
	return leArquivo(arquivosLogs)
}

func leArquivo(nomeArquivo string) []string {
	var arquivos []string

	arquivo, err := os.Open(nomeArquivo)

	if err == nil {
		leitor := bufio.NewReader(arquivo)

		for {
			linha, err := leitor.ReadString('\n')

			linha = strings.TrimSpace(linha)

			arquivos = append(arquivos, linha)

			if err == io.EOF {
				break
			}
		}

		arquivo.Close()
	} else {
		fmt.Println("Ocorreu um erro:", err)
	}

	return arquivos
}

func RegistraLog(site string, status int) {
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
