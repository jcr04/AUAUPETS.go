# AUAUPETS

AUAUPETS é uma API de gestão de animais e reservas construída em Go. É uma solução perfeita para gerenciar pets e suas estadias.

## Leia Code documentation
[Codedoc](https://github.com/jcr04/AUAUPETS.go/blob/main/codedoc.md)

## 🚀 Começando


Essas instruções fornecerão uma cópia do projeto em execução na sua máquina local para desenvolvimento e testes.

### 📋 Pré-requisitos

- Go versão 1.x
- PostgresSQL

### 🔧 Instalação

1. Clone o repositório para a sua máquina local:
```bash
git clone https://github.com/jcr04/AUAUPETS.go.git
```
2. Navegue até o diretório do projeto:
    cd AUAUPETS.go
3. Instale todas as dependências necessárias:
    go mod tidy
4. Configure o banco de dados Postgres e atualize o arquivo de configuração de banco de dados conforme necessário.
    1. execute essas qeurys:
        1. 
        ```sql
        CREATE DATABASE Pets
        ```
        em seguida coloque
        
        ```sql
        CREATE TABLE Pets (
        ID SERIAL PRIMARY KEY,
        Name VARCHAR(255),
        Breed VARCHAR(255),
        Age INT,
        CheckIn DATE,
        CheckOut DATE
         );
        ```
    2. Execute o servidor:
        go run cmd/server/main.go
