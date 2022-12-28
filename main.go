package main

import (
	"core"
	"view"
)

func main() {
	view.ExibeIntroducao()

	core.ExecutaCore(1)

	for {
		comando := view.ExibirMenu()

		view.Imprime(core.ExecutaCore(comando))
	}
}
