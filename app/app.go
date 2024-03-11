// Pacote app é responsável por definir a aplicação de linha de comando.
package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a aplicação de linha de comando pronta para ser executada.
func Gerar() *cli.App {
	// Criar uma nova instância da aplicação CLI.
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	// Definir as flags comuns para os comandos.
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "marcuscarvalho.tech",
			Usage: "Define o host para a busca de IPs e Nomes de Servidor",
		},
	}

	// Definir os comandos disponíveis na aplicação.
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// buscarIps é uma função de ação para o comando "ip".
func buscarIps(c *cli.Context) error {
	// Obter o valor da flag "host" do contexto.
	host := c.String("host")

	// Realizar a pesquisa de IPs para o host especificado.
	ips, erro := net.LookupIP(host)
	if erro != nil {
		return erro
	}

	// Imprimir os IPs encontrados.
	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

// buscarServidores é uma função de ação para o comando "servidores".
func buscarServidores(c *cli.Context) error {
	// Obter o valor da flag "host" do contexto.
	host := c.String("host")

	// Realizar a pesquisa de servidores para o host especificado.
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		return erro
	}

	// Imprimir os nomes dos servidores encontrados.
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

	return nil
}
