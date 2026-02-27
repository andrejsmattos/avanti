# Gerador de Saudações

Este projeto é uma aplicação web de front-end que gera saudações aleatórias. Ele busca dinamicamente uma saudação e o nome de uma pessoa de duas APIs de back-end distintas e os exibe. A interface também permite que os usuários cadastrem novas pessoas e novas saudações, que são enviadas para as respectivas APIs.

## 🚀 Funcionalidades

* **Geração Aleatória**: Exibe uma saudação e um nome de pessoa obtidos aleatoriamente a partir de APIs.
* **Nova Saudação**: Um botão "Gerar Nova Saudação" permite buscar uma nova combinação a qualquer momento.
* **Cadastro de Pessoas**: Um formulário para cadastrar um novo nome de pessoa. Os dados são enviados via requisição `POST` para a API de pessoas.
* **Cadastro de Saudações**: Um formulário para cadastrar uma nova saudação. Os dados são enviados via requisição `POST` para a API de saudações.
* **Feedback ao Usuário**: Exibe mensagens de sucesso ou erro após as tentativas de cadastro. As mensagens desaparecem automaticamente após 4 segundos.
* **Indicador de Carregamento**: Mostra um ícone de "spinner" enquanto os dados estão sendo buscados nas APIs.
* **Design Responsivo**: A interface se adapta a diferentes tamanhos de tela, de dispositivos móveis a desktops.

## 🛠️ Tecnologias Utilizadas

* **HTML5**: Estrutura base da página.
* **Tailwind CSS**: Framework de CSS para estilização rápida e moderna.
* **Alpine.js**: Framework JavaScript minimalista para compor a reatividade e a interatividade da interface.

## 📋 Pré-requisitos (Back-end)

Este projeto é **apenas o front-end** e depende de duas APIs de back-end que devem estar em execução localmente para que a aplicação funcione corretamente.

1.  **API de Pessoas**: Deve estar rodando em `http://localhost:8000`.
2.  **API de Saudações**: Deve estar rodando em `http://localhost:8080`.

### Endpoints da API

A aplicação interage com os seguintes endpoints:

* **Pessoas**:
    * `GET /pessoas/aleatoria`: Retorna um objeto JSON com o nome de uma pessoa aleatória.
        * Exemplo de resposta: `{ "nome": "João" }`
    * `POST /pessoas`: Cadastra uma nova pessoa.
        * Exemplo de corpo da requisição: `{ "nome": "Maria" }`

* **Saudações**:
    * `GET /api/saudacoes/aleatorio`: Retorna um objeto JSON com uma saudação aleatória.
        * Exemplo de resposta: `{ "saudação": "Olá, tudo bem?" }`
    * `POST /api/saudacoes`: Cadastra uma nova saudação.
        * Exemplo de corpo da requisição: `{ "saudação": "E aí, tudo certo?" }`

## ▶️ Como Executar

1.  Certifique-se de que as duas APIs de back-end (Pessoas e Saudações) estão em execução nas portas `8000` e `8080`, respectivamente.
2.  Clone ou baixe este repositório.
3.  Abra o arquivo `index.html` em seu navegador de preferência ou use um servidor web para visualizar a aplicação.

A aplicação irá carregar e buscar automaticamente a primeira saudação.

## 📂 Estrutura do Código

Toda a lógica da aplicação está contida no arquivo `index.html`.

* A **estrutura HTML** define os elementos da página, como o título, os botões e os formulários.
* A **estilização** é feita com classes do **Tailwind CSS** diretamente no HTML.
* A **interatividade** é controlada pelo **Alpine.js** dentro de um bloco `<script>`.
    * A função `greetingApp()` inicializa o estado da aplicação (variáveis como `greeting`, `personName`, `isLoading`, etc.) e os métodos para interagir com as APIs (`fetchData`, `addPerson`, `addGreeting`).
    * `fetchData()`: Utiliza `Promise.all` para buscar dados das duas APIs simultaneamente, melhorando o tempo de carregamento.
    * `addPerson()` e `addGreeting()`: Funções assíncronas que enviam os dados dos formulários para as respectivas APIs usando o método `POST`.
    * `showFeedback()` e `clearFeedback()`: Controlam a exibição de mensagens para o usuário.