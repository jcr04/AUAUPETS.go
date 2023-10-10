# Documentação Técnica: AUAUPETS API

A API AUAUPETS é uma solução de gestão de animais e reservas construída em Go. Esta documentação técnica detalha a estrutura do projeto, os principais componentes e como interagir com a API.

## Índice

- [Estrutura do Projeto](#estrutura-do-projeto)
- [Configuração do Banco de Dados](#configuração-do-banco-de-dados)
- [Modelos de Domínio](#modelos-de-domínio)
- [Repositórios](#repositórios)
- [Manipuladores](#manipuladores)
- [Rotas](#rotas)
- [Executando a API](#executando-a-api)

## Estrutura do Projeto

O projeto segue uma estrutura modular e organizada, facilitando a manutenção e a expansão. Aqui estão os principais diretórios e arquivos do projeto:

- `/internal`:
  - `/application/handler`: Contém os manipuladores de requisições HTTP.
  - `/domain`: Contém os modelos de domínio.
  - `/infrastructure`: Contém código relacionado à infraestrutura, como configuração de banco de dados e roteamento.
- `main.go`: Ponto de entrada da aplicação.

## Configuração do Banco de Dados

A configuração do banco de dados está localizada em `internal/infrastructure/database/db.go`. Utilizamos o PostgreSQL como nosso banco de dados.

## Modelos de Domínio

Os modelos de domínio são definidos em `/internal/domain`. Atualmente, temos dois modelos principais:

- `Animal`: Representa um animal.
- `Reservation`: Representa uma reserva.

## Repositórios

Os repositórios fornecem uma camada de abstração sobre o acesso ao banco de dados. Eles estão localizados em `/internal/infrastructure/repository`.

- `PetRepository`: Provê métodos para interagir com os registros de animais no banco de dados.
- `ReservationRepository`: Provê métodos para interagir com os registros de reservas no banco de dados.

## Manipuladores

Os manipuladores estão localizados em `/internal/application/handler` e contêm a lógica para manipular requisições HTTP e responder a elas.

## Rotas

As rotas da API são definidas em `internal/infrastructure/api/router.go`. Utilizamos a biblioteca `gorilla/mux` para definir as rotas.

A API oferece os seguintes endpoints:

Animnals:
/animals (POST): Criar animnal 
* - ![Screenshot_22](https://github.com/jcr04/AUAUPETS.go/assets/70778525/a9f02d16-d423-4aa7-b37f-ade47ffcbd24)
* - ![Screenshot_23](https://github.com/jcr04/AUAUPETS.go/assets/70778525/c94846c0-049d-4103-ae21-fc0f27642338)

## Executando a API

Para executar a API, siga os passos abaixo:

1. Navegue até o diretório raiz do projeto.
2. Execute o comando `go run cmd/server/main.go`.

A API estará disponível no endereço `http://localhost:8080`.
