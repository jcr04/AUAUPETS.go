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
![Screenshot_1](https://github.com/jcr04/AUAUPETS.go/assets/70778525/014cdaea-7e69-4fb6-81f3-6f3a424be0d0)

As rotas da API são definidas em `internal/infrastructure/api/router.go`. Utilizamos a biblioteca `gorilla/mux` para definir as rotas.

A API oferece os seguintes endpoints:

Animals:
- /animals (POST): Criar animal 
* - ![Screenshot_3](https://github.com/jcr04/AUAUPETS.go/assets/70778525/609a1ba1-0bc2-4c82-9300-4d7239d09a7b)
/animals (GET): Listar Todos os Animais
* - ![Screenshot_4](https://github.com/jcr04/AUAUPETS.go/assets/70778525/f044fc3e-b060-41fe-9ec5-ef3a20899c7c)
/animals/{id} (GET): Obter Detalhes de um Animal Específico
* - ![Screenshot_5](https://github.com/jcr04/AUAUPETS.go/assets/70778525/7ec256af-784a-4947-835f-4d7194ae7bc9)
/animals/{id} (PUT): Atualizar Detalhes de um Animal Específico
* - ![Screenshot_5](https://github.com/jcr04/AUAUPETS.go/assets/70778525/52e2a87b-19fe-46e0-aa9d-f516f175274c)
* - ![Screenshot_3](https://github.com/jcr04/AUAUPETS.go/assets/70778525/ccf7a1ca-6754-4b0e-96c4-203a958f4304)
/animals/{id} (DELETE): Deletar um Animal Específico
* - ![Screenshot_4](https://github.com/jcr04/AUAUPETS.go/assets/70778525/12663fa7-37be-4a11-95e9-c9795acc8af0)
* - ![Screenshot_5](https://github.com/jcr04/AUAUPETS.go/assets/70778525/3e6bea1c-4331-4bfa-bd05-d4222a31afe8)

Reservations:
- /reservations (POST) Criar Reserva
* - ![Screenshot_6](https://github.com/jcr04/AUAUPETS.go/assets/70778525/02cfff17-d745-44d1-b1d5-c66c78e21b77)
/reservations (GET): Listar todas as Reservas
* - ![Screenshot_11](https://github.com/jcr04/AUAUPETS.go/assets/70778525/da887630-74f2-4304-b872-996226e372c6)

Hosting: 
- /hosting (POST): Criar Hospedagem
* - ![Screenshot_7](https://github.com/jcr04/AUAUPETS.go/assets/70778525/a88274f3-4f71-4420-98c9-bf9c35fdc2c4)
- /hosting (GET): Listar todas as Hospedagem
* - ![Screenshot_2](https://github.com/jcr04/AUAUPETS.go/assets/70778525/a7e126cb-bc2c-478e-b766-cb6bfd99c438)


## Executando a API

Para executar a API, siga os passos abaixo:

1. Navegue até o diretório raiz do projeto.
2. Execute o comando `go run cmd/server/main.go`.

A API estará disponível no endereço `http://localhost:8080`.
