// Pacote main é o ponto de entrada principal para a execução do programa.
package main

import (
	"appCLI/app" // Importa o pacote app, que contém a lógica da aplicação de linha de comando.
	"log"
	"os"
)

// Função main é a função principal que inicia a execução do programa.
func main() {
	// Chama a função Gerar do pacote app para obter a aplicação de linha de comando.
	aplicacao := app.Gerar()

	// Executa a aplicação de linha de comando, passando os argumentos da linha de comando.
	// Qualquer erro durante a execução é tratado e impresso no log.
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
