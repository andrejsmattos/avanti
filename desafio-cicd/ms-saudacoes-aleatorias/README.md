# API de Saudações Aleatórias

Este é um simples microserviço RESTful construído em Go que fornece saudações aleatórias e permite o cadastro de novas saudações.

## ✨ Funcionalidades

  * Obter uma saudação aleatória do banco de dados.
  * Cadastrar uma nova saudação.
  * Utiliza o framework Gin para o roteamento e gerenciamento das requisições HTTP.
  * Usa GORM como ORM para interagir com o banco de dados.
  * Utiliza SQLite como banco de dados, que é criado e populado automaticamente na primeira execução.
  * O ambiente de desenvolvimento é gerenciado pelo Devbox.

## 🛠️ Tecnologias Utilizadas

  * **Go (Golang)**: Linguagem de programação principal.
  * **Gin**: Framework web para Go.
  * **GORM**: ORM para Go.
  * **SQLite**: Banco de dados SQL embarcado.
  * **Devbox**: Ferramenta para criar ambientes de desenvolvimento isolados.

## 🚀 Como Executar o Projeto

### Pré-requisitos

Antes de começar, você precisa ter o [Devbox](https://www.google.com/search?q=https://www.jetify.com/devbox/docs/installing-devbox/) instalado em sua máquina.

### Passos

1.  **Clone o repositório:**

    ```bash
    git clone <URL_DO_SEU_REPOSITORIO>
    cd ms-saudacoes-aleatorias
    ```

2.  **Inicie o ambiente Devbox:**
    O Devbox instalará automaticamente o Go na versão especificada no arquivo `devbox.json`.

    ```bash
    devbox shell
    ```

3.  **Execute a aplicação:**
    Este comando irá iniciar o servidor na porta `8080`.

    ```bash
    go run main.go
    ```

Ao iniciar, a aplicação criará um arquivo de banco de dados chamado `greetings.db` e o populará com uma lista inicial de saudações.

## 📖 API Endpoints

A API possui o prefixo `/api`.

### Obter uma Saudação Aleatória

Retorna uma saudação aleatória do banco de dados.

  * **Método:** `GET`
  * **Endpoint:** `/api/saudacoes/aleatorio`
  * **Resposta de Sucesso (200 OK):**
    ```json
    {
      "saudação": "Que a Força esteja com você"
    }
    ```
  * **Exemplo com cURL:**
    ```bash
    curl http://localhost:8080/api/saudacoes/aleatorio
    ```

### Cadastrar uma Nova Saudação

Adiciona uma nova saudação ao banco de dados.

  * **Método:** `POST`
  * **Endpoint:** `/api/saudacoes`
  * **Corpo da Requisição (JSON):**
    O campo `text` é obrigatório.
    ```json
    {
      "text": "Sua nova saudação aqui"
    }
    ```
  * **Resposta de Sucesso (201 Created):**
    ```json
    {
      "data": {
        "ID": 10,
        "CreatedAt": "2024-05-18T16:05:23.038166-03:00",
        "UpdatedAt": "2024-05-18T16:05:23.038166-03:00",
        "DeletedAt": null,
        "Text": "Sua nova saudação aqui"
      }
    }
    ```
  * **Exemplo com cURL:**
    ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"text":"Live long and prosper"}' \
      http://localhost:8080/api/saudacoes
    ```