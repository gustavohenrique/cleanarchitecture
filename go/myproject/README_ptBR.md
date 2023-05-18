# Go para API Rest e gRPC server

## Índice

- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Estrutura de pastas e diretórios](#estrutura-de-pastas-e-diretórios)
- [Como funciona?](#como-funciona)
- [Convenções](#convenções)
- [Novo projeto](#novo-projeto)
- [Como contribuir?](#como-contribuir)


## Arquitetura

Arquitetura de software é uma estrutura que suporta os casos de uso do projeto
com o objetivo de facilitar a manutenção do código.
A arquitetura criada aqui foi inspirada em Arquitetura Hexagonal, Clean Architecture 
e Domain-Driven Design (DDD). As seguintes premissas foram adotadas:

- Independência de Frameworks
- Fácil de testar
- Estrutura de dados flexível
- Injeção de dependência
- Encapsulamento de bibliotecas externas
- Agilidade na mudança de banco de dados

As decisões para alcançar as premissas foram:

- Models representam estrutura de dados do domínio da aplicação e regras de negócio
- Wires são estruturas de dados que representam o input e output da aplicação e converte elas para models
- Repositories armazenam e recuperam informações de sistemas externos (banco de dados, APIs remotas, filesystem...)
- Usecases contém as regras de negócio e podem orquestrar a utilização de múltiplos repositories
- Ports contém as interfaces da camada domain
- Infra é a camada para código que não pertencem a lógica de negócio
- Controllers orquestram o fluxo da aplicação desde a entrada de dados até a saída deles
- Suporte a múltiplos databases
- Suporte a API Rest, gRPC, gRPC-Web e Websockets


## Estrutura de pastas e diretórios

```
.
├── cmd/                               # The place where main.go lives
│   ├── webapp/
│   │   └── main.go
├── src/
│   ├── adapters/
│   │   ├── controllers/               # Functions used by servers to handle requests
│   │   └── repositories/              # Save and fetch data from datastores
│   ├── infrastructure/                # Wrap libs and frameworks to deal with external world
│   │   ├── servers/                   # Every code to listen for incoming requests
│   │   │   ├── servers.go             # Container for servers
│   │   │   ├── natsserver/
│   │   │   ├── httpserver/
│   │   │   ├── grpcwebserver/
│   │   │   └── grpcserver/
│   │   ├── clients/                   # Every code to deal with outcome requests
│   │   │   └── natsclient/
│   │   ├── datastores/                # Databases and external storage
│   ├── domain/                        # Where the application and business logic lives
│   │   ├── models/                    # Data structure that represents business and its rules
│   │   ├── usecases/                  # Business roles
│   │   └── ports/                     # Interfaces of the domain
│   ├── wire/                          # Map input/output data structure to/from models
│   ├── components/                    # Utilities that could be used by any layers
│   │   ├── configurator/              # A singleton for the entire application config
│   │   └── logger/
├── assets/                            
│   ├── static/                        # Static files will be embed and served via /static endpoint
│   │   └── proto/                     # Protocol buffer definitions source files
│   └── web/                           # HTML files will be embed and renderized by HTTP server
├── init/                              # System init and process manager/supervisor configs
├── configs/                           # Examples of config.yaml
├── sdk/                               # SDK for clients
├── scripts/                           # Shell scripts for dev/devops needs
├── migrations/                        # History of database changes and initial data
├── .makefiles/                        # Automates some development tasks
├── .editorconfig                      # Editor config
└── .gitignore                         # GIT ignore paths
```

## Como funciona?

### Tecnologias

#### Echo Framework

Echo é uma framework para desenvolvimento web que facilita a criação de APIs
Rest e renderização de templates HTML.  Suas principais características são:

- Router otimizado com zero alocação de memória dinâmica Organização de APIs em grupos
- Suporte a HTTP/2 Fácil de adicionar middlewares Data binding para o payload de request HTTP
- Renderização de templates

#### gRPC (Google Remote Procedure Call)

O gRPC é um protocolo de comunicação criado pelo Google no qual serviços
distribuídos podem se comunicar sem necessidade de entendimento da camada de
rede.  Os dados são transmitidos pela rede em formato binário, diferente do
HTTP que utiliza texto puro.  

Esse protocolo utiliza o modelo client/server. O desenvolvedor especifica em um
arquivo texto, utilizando uma sintaxe chamada Protobuf, quais serviços serão
expostos no server.  O arquivo deve ser disponibilizado para que outros
desenvolvedores possam utilizar o compilador `protoc` para ler essas definições
e criar um **stub** para a linguagem de programação utilizada pelo client.  Por
fim, o client só precisa importar o stub como uma biblioteca de terceiros e
invocar as funções dos serviços expostos pelo server.

#### gRPC Web

O gRPC Web é uma implementação Javascript de gRPC client para browsers.  Ele é
totalmente compatível com gRPC servers rodando em Go. Para outras linguagens é
necessário adicionar o proxy Envoy para traduzir a requisição enviada do
browser antes de encaminha-la para o server.

Basta importar os stubs gerados para Javascript e chamar as respectivas funções
de acordo com o serviço exposto. O gRPC Web converte os dados para binário e
envia usando HTTP POST. O server em Go roda um servidor HTTP que intercepta as
requisições desse tipo e trata como requisições gRPC.

Exemplo de um arquivo com definições de serviços escrito em Protobuf:

```proto
syntax = "proto3";

package private;
option go_package = "proto";

service AccountService {
    rpc Create(User) returns (User) {}
    rpc AuthenticateByEmailAndPassword(User) returns (Account) {}
}

message User {
    string email = 1;
    string password = 2;
}

message Account {
    string token = 1;
}
```

#### Nats

Nats é um sistema de menssageria que atua como pub/sub.

### Arquitetura

Fluxo dos dados:

1. Controller recebe o input de dados. Deve existir controllers diferentes para cada tipo, HTTP, gRPC e gRPC Web
2. Controller utiliza o wire para converter a estrutura de dados de/para models
3. Controller se comunica com usecases ou repositories
4. Usecases possuem regras de negócio ou da aplicação e podem utilizar repositories
5. Repository armazena ou recupera os dados a partir de um database ou API externa

## Novo projeto

### Primeiros passos

Instale as dependências com `make setup` e siga as instruções para a instalação manual das ferramentas para gRPC.

```sh
make setup     # Instala as dependências
make sqlite    # Cria um arquivo database.db utilizando migrations/sqlite/schema.sql
make postgres  # Cria um container Docker rodando Postgres utilizando migrations/postgres/schema.sql
make proto     # Roda o compilador protoc para gerar os stubs a partir do assets/static/main.proto
make lint      # Roda o goimports para formatar o código Go
make docs      # Cria doc swagger
make mocks     # Cria mocks para os testes
make test      # Executa os testes unitários
make start     # Roda a aplicação em localhost
```

## Como contribuir?

1. Abra uma issue para discutir sua proposta de mudança e se ela será aceita
2. Faça um fork desse repositório se sua mudança for aceita
3. Crie um feature branch: `git checkout -b my-new-feature`
4. Faça um commit e push das suas mudanças: `git commit -am 'Add some feature' && git push origin my-new-feature`
5. Crie um pull request
