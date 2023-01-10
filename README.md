# API de processamento de dados em Go

Este projeto desenvolve uma API em Go que permite o processamento de dados e o envio dos resultados para um tópico do pub/sub.

## Instalação

Para instalar o projeto em sua máquina, siga os seguintes passos:

1. Clone o repositório para sua máquina:
~~~
git clone https://github.com/PauloMarvin/Golang-API-CVM
~~~

2. Instale as dependências do projeto:
~~~
go get -d -v ./...
~~~

3. Compile o projeto:
~~~
go build
~~~

4. Execute o projeto:
~~~
go run .\main.go
~~~

## Uso

A API expõe um endpoint `/process` que aceita um payload de entrada e retorna os resultados do processamento. Por exemplo:

~~~
curl -X POST -d '{"dados_entrada": "valor"}' http://localhost:8000/data
~~~

O resultado do processamento será enviado para o tópico do pub/sub especificado nas configurações.

## Dependências

- Go versão 1.15 ou superior
- Pacote pubsub do Google Cloud

## Licença

Este projeto está disponível sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
