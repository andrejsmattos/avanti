# API de Pessoas Aleatórias

Uma API simples desenvolvida em FastAPI para cadastrar e sortear pessoas. O projeto utiliza SQLAlchemy para interação com um banco de dados SQLite.

## Descrição

Esta API oferece funcionalidades para adicionar novas pessoas a um banco de dados e obter uma pessoa aleatória dentre as cadastradas. Na inicialização, o banco de dados é populado com uma lista predefinida de nomes para garantir que haja dados disponíveis para teste.

## Funcionalidades

  * Cadastrar novas pessoas.
  * Sortear uma pessoa aleatória do banco de dados.
  * População inicial do banco de dados com nomes pré-definidos.
  * Documentação da API gerada automaticamente via Swagger UI.
  * Suporte a CORS para permitir requisições de diferentes origens.

## Tecnologias Utilizadas

  * **Python**: Linguagem de programação principal.
  * **FastAPI**: Framework web para a construção da API.
  * **SQLAlchemy**: ORM para manipulação do banco de dados.
  * **Uvicorn**: Servidor ASGI para rodar a aplicação FastAPI.
  * **SQLite**: Banco de dados relacional utilizado no projeto.

## Pré-requisitos

  * Python 3.x

## Instalação e Execução

1.  **Clone o repositório:**

    ```bash
    git clone <url-do-repositorio>
    cd <diretorio-do-repositorio>
    ```

2.  **Crie e ative um ambiente virtual:**

    ```bash
    python -m venv .venv
    source .venv/bin/activate  # Para Linux/macOS
    # ou
    .venv\Scripts\activate  # Para Windows
    ```

3.  **Instale as dependências:**

    ```bash
    pip install -r requirements.txt
    ```

4.  **Inicie a aplicação:**

    ```bash
    uvicorn main:app --reload
    ```

    O servidor estará disponível em `http://127.0.0.1:8000`.

## Endpoints da API

A API possui os seguintes endpoints:

### Pessoas

  * **`POST /pessoas/`**

      * **Descrição:** Cadastra uma nova pessoa no banco de dados.
      * **Status Code:** `201 CREATED`
      * **Corpo da Requisição:**
        ```json
        {
          "nome": "string"
        }
        ```
      * **Resposta:**
        ```json
        {
          "id": 0,
          "nome": "string"
        }
        ```
      * **Observação:** Retorna um erro `400 Bad Request` se o nome já existir.

  * **`GET /pessoas/aleatoria/`**

      * **Descrição:** Retorna uma pessoa aleatória cadastrada no banco de dados.
      * **Resposta:**
        ```json
        {
          "id": 0,
          "nome": "string"
        }
        ```
      * **Observação:** Retorna um erro `404 Not Found` se não houver pessoas cadastradas.

### Raiz

  * **`GET /`**
      * **Descrição:** Endpoint raiz que exibe uma mensagem de boas-vindas e direciona para a documentação.
      * **Resposta:**
        ```json
        {
          "message": "Bem-vindo à API de Pessoas Aleatórias! Acesse /docs para ver a documentação."
        }
        ```

## Documentação Interativa

Acesse `http://127.0.0.1:8000/docs` no seu navegador para visualizar e interagir com a documentação da API gerada pelo Swagger UI.

## Banco de Dados

  * A aplicação utiliza um banco de dados SQLite chamado `pessoas.db`, que é criado automaticamente na raiz do projeto.
  * O schema do banco de dados é definido no arquivo `models.py` e inclui uma tabela `pessoas` com os campos `id` e `nome`.
  * Na inicialização da API, o banco de dados é populado com uma lista de nomes pré-definidos se ainda não estiverem presentes. Os nomes são: "Alice", "Bruno", "Carla", "Daniel", "Eva", "Fábio", "Gabriela", "Heitor", "Íris", "João", "Larissa", e "Marcos".

## Estrutura do Projeto

```
.
├── .gitignore
├── database.py       # Configuração da conexão com o banco de dados
├── devbox.json       # Configuração do ambiente de desenvolvimento Devbox
├── main.py           # Lógica principal da API, endpoints e inicialização
├── models.py         # Modelos ORM do SQLAlchemy
├── requirements.txt  # Dependências Python
└── schemas.py        # Schemas Pydantic para validação de dados
```