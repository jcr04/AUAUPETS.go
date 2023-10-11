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

Animals:
/animals (POST): Criar animal 
* - ![Screenshot_22](https://github.com/jcr04/AUAUPETS.go/assets/70778525/a9f02d16-d423-4aa7-b37f-ade47ffcbd24)
* - ![Screenshot_23](https://github.com/jcr04/AUAUPETS.go/assets/70778525/c94846c0-049d-4103-ae21-fc0f27642338)
/animals (GET): Listar Todos os Animais
* - ![Screenshot_1](https://github.com/jcr04/AUAUPETS.go/assets/70778525/ed183851-fcd8-42b4-9313-10f37a3975e4)
/animals/{id} (GET): Obter Detalhes de um Animal Específico
* - ![Screenshot_6](https://github.com/jcr04/AUAUPETS.go/assets/70778525/e0424da2-2a5b-4834-9193-2c4076fa96ea)
/animals/{id} (PUT): Atualizar Detalhes de um Animal Específico
* - ![Screenshot_2](https://github.com/jcr04/AUAUPETS.go/assets/70778525/a51d015f-8161-4935-b049-3529da01e530)
* - ![Screenshot_3](https://github.com/jcr04/AUAUPETS.go/assets/70778525/ccf7a1ca-6754-4b0e-96c4-203a958f4304)
/animals/{id} (DELETE): Deletar um Animal Específico
* - ![Screenshot_4](https://github.com/jcr04/AUAUPETS.go/assets/70778525/12663fa7-37be-4a11-95e9-c9795acc8af0)
* - ![Screenshot_5](https://github.com/jcr04/AUAUPETS.go/assets/70778525/3e6bea1c-4331-4bfa-bd05-d4222a31afe8)

## Executando a API

Para executar a API, siga os passos abaixo:

1. Navegue até o diretório raiz do projeto.
2. Execute o comando `go run cmd/server/main.go`.

A API estará disponível no endereço `http://localhost:8080`.
