# Aplicação de Linha de Comando para Busca de IPs e Nomes de Servidor

Esta é uma aplicação de linha de comando (CLI) escrita em Go (Golang) que permite buscar informações sobre IPs e nomes de servidor na internet. A aplicação fornece dois comandos principais: `ip` para buscar IPs associados a um host e `servidores` para buscar os nomes dos servidores associados a um host.

## Estrutura do Projeto

O projeto está organizado em dois arquivos principais:

1. **main.go**: Este é o ponto de entrada principal para a execução do programa. Ele importa o pacote `app`, que contém a lógica da aplicação de linha de comando. A função `main` chama a função `Gerar` do pacote `app` para obter a aplicação de linha de comando e executa essa aplicação.

2. **app.go**: Este arquivo contém a definição da aplicação de linha de comando. A função `Gerar` retorna uma instância pronta para ser executada do pacote [urfave/cli](https://github.com/urfave/cli), que é utilizado para criar interfaces de linha de comando em Go. A aplicação possui dois comandos principais (`ip` e `servidores`) e usa a biblioteca `net` para realizar buscas de IPs e servidores.

## Como Usar

### Instalação

Certifique-se de ter o Go instalado em seu sistema. Para baixar e instalar, visite [golang.org](https://golang.org/doc/install).

Clone o repositório:

```bash
git clone https://github.com/seu-usuario/nome-do-repositorio.git
cd nome-do-repositorio
```

### Execução

Para executar a aplicação, utilize o seguinte comando:

```bash
go run main.go [comando] [flags]
```

Os comandos disponíveis são `ip` e `servidores`, e ambos aceitam a flag `--host` para especificar o host alvo.

#### Exemplos:

1. Buscar IPs para o host padrão (marcuscarvalho.tech):

```bash
go run main.go ip
```

2. Buscar IPs para um host específico:

```bash
go run main.go ip --host example.com
```

3. Buscar servidores para o host padrão:

```bash
go run main.go servidores
```

4. Buscar servidores para um host específico:

```bash
go run main.go servidores --host example.com
```

## Detalhes de Implementação

A aplicação utiliza a biblioteca `net` para realizar buscas de IPs (`LookupIP`) e servidores (`LookupNS`). As informações são então impressas no console.

### Estrutura de Código

- `main.go`: Ponto de entrada principal e execução do programa.
- `app.go`: Definição da aplicação de linha de comando e lógica dos comandos.

- `buscarIps`: Função de ação para o comando `ip`, que realiza a pesquisa de IPs para o host especificado.
- `buscarServidores`: Função de ação para o comando `servidores`, que realiza a pesquisa de servidores para o host especificado.

### Flags

- `--host`: Permite especificar o host alvo para a busca de IPs e servidores.

## Contribuição

Sinta-se à vontade para contribuir, relatar problemas ou sugerir melhorias. Abra uma issue ou envie um pull request!

**Divirta-se explorando informações na internet com esta simples aplicação de linha de comando!**