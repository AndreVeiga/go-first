package core

import (
	"integration"
	"os"
	"repository"
)

func ExecutaCore(comando int) []string {
	var result []string
	switch comando {
	case 1:
		iniciarMonitoracao()
	case 2:
		result = exibirLogs()
	case 0:
		os.Exit(0)
	default:
		os.Exit(-1)
	}

	return result
}

func iniciarMonitoracao() {
	sites := repository.LeSites()

	for _, site := range sites {
		status, _ := integration.TestaSite(site)
		repository.RegistraLog(site, status)
	}
}

func exibirLogs() []string {
	return repository.LeLogs()
}
