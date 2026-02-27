A seguir, apresento um guia passo a passo detalhado para construir um `Dockerfile` otimizado para seu projeto, criar a imagem Docker e enviá-la para o Docker Hub, seguindo as melhores práticas.

### Pré-requisitos

1.  **Docker Instalado:** Você precisa ter o Docker instalado e em execução em sua máquina.
2.  **Conta no Docker Hub:** Crie uma conta gratuita no [Docker Hub](https://hub.docker.com/). Você precisará do seu nome de usuário e senha para fazer o login e enviar a imagem.

-----

### Estrutura do Projeto

Para este guia, vamos assumir que o seu arquivo `index.html` está em uma pasta de projeto. Por exemplo:

```
/site-gerador-saudacoes
└── index.html
```

É dentro desta pasta `/site-gerador-saudacoes` que criaremos nosso `Dockerfile`.

-----

### Passo 1: Criando o `Dockerfile`

O `Dockerfile` é um arquivo de texto que contém as instruções para o Docker montar sua imagem. Crie um arquivo chamado `Dockerfile` (sem extensão) na raiz do seu projeto, ao lado do `index.html`, com o seguinte conteúdo:

```dockerfile
# --- Estágio 1: Definir a imagem base ---
# Usamos a imagem oficial do Nginx com a tag 'alpine'.
# 'alpine' resulta em uma imagem muito menor, o que é ótimo para produção.
FROM nginx:alpine

# --- Estágio 2: Copiar os arquivos do projeto ---
# Copia o arquivo 'index.html' da sua máquina local (o contexto do build)
# para o diretório padrão onde o Nginx serve os arquivos HTML.
COPY index.html /usr/share/nginx/html/index.html

# --- Estágio 3: Expor a porta ---
# Informa ao Docker que o contêiner escutará na porta 80 em tempo de execução.
# Esta é a porta padrão do Nginx.
EXPOSE 80
```

#### Entendendo cada linha do Dockerfile:

  * `FROM nginx:alpine`: Esta é a instrução mais importante. Ela define qual imagem usaremos como base. Escolhemos `nginx:alpine`, uma versão leve e segura do popular servidor web Nginx.
  * `COPY index.html /usr/share/nginx/html/index.html`: Esta linha copia o seu arquivo `index.html` para dentro da imagem, no diretório `/usr/share/nginx/html/`. Este é o local padrão onde o Nginx procura por arquivos para servir.
  * `EXPOSE 80`: Documenta que o contêiner irá expor a porta `80` (porta padrão para HTTP). Isso não publica a porta de fato, mas serve como uma documentação para quem for usar a imagem.

-----

### Passo 2: Construindo a Imagem Docker

Agora que o `Dockerfile` está pronto, vamos usá-lo para construir (ou "buildar") a imagem.

1.  **Abra o seu terminal** (Prompt de Comando, PowerShell ou o terminal do seu editor de código).

2.  **Navegue até a pasta do seu projeto**, onde estão o `index.html` e o `Dockerfile`.

3.  **Execute o comando `docker build`**:

    ```bash
    docker build -t seu-usuario-dockerhub/gerador-saudacoes:1.0 .
    ```

#### Analisando o comando:

  * `docker build`: O comando principal para construir uma imagem.
  * `-t`: Abreviação de `--tag`. É como damos um "nome" à nossa imagem, no formato `repositorio:tag`.
      * `seu-usuario-dockerhub`: **Substitua por seu nome de usuário real do Docker Hub.** Isso é essencial para poder enviar a imagem mais tarde.
      * `gerador-saudacoes`: O nome que você quer dar ao seu repositório de imagem (o nome do projeto).
      * `:1.0`: Uma tag de versão. É uma boa prática versionar suas imagens.
  * `.` (ponto no final): Indica que o "contexto" do build é o diretório atual. Ou seja, o Docker procurará o `Dockerfile` nesta pasta e terá acesso aos arquivos dela (como o `index.html`).

Após executar o comando, o Docker fará o download da imagem `nginx:alpine` (se você ainda não a tiver) e executará os passos definidos no seu `Dockerfile`.

-----

### Passo 3: Rodando e Testando a Imagem Localmente

Antes de enviar para o mundo, vamos garantir que a imagem funciona na sua máquina.

1.  **Execute o comando `docker run`**:

    ```bash
    docker run -d -p 8080:80 seu-usuario-dockerhub/gerador-saudacoes:1.0
    ```

#### Analisando o comando:

  * `docker run`: O comando para criar e iniciar um contêiner a partir de uma imagem.
  * `-d`: Roda o contêiner em modo "detached" (em segundo plano), liberando seu terminal.
  * `-p 8080:80`: Mapeia as portas. Esta é uma parte crucial.
      * `8080`: É a porta na sua máquina (o "host").
      * `80`: É a porta dentro do contêiner (que expusemos com `EXPOSE 80`).
      * Isso significa que qualquer requisição que chegar na porta `8080` do seu computador será redirecionada para a porta `80` do contêiner Nginx.
  * `seu-usuario-dockerhub/gerador-saudacoes:1.0`: O nome da imagem que você quer rodar.

<!-- end list -->

2.  **Teste no Navegador:** Abra seu navegador e acesse `http://localhost:8080`. Você deverá ver a sua página "Gerador de Saudações" funcionando\!

> **Observação Importante:** O seu código JavaScript faz chamadas para `localhost:8080` e `localhost:8000`. Quando você acessa a página servida pelo Docker, o "localhost" no contexto do JavaScript se refere à sua própria máquina, não ao contêiner. Portanto, para que a aplicação funcione completamente, as duas APIs de back-end (Saudações e Pessoas) ainda precisam estar rodando na sua máquina local nessas portas.

-----

### Passo 4: Enviando a Imagem para o Docker Hub

Com a imagem construída e testada, é hora de compartilhá-la.

1.  **Faça login no Docker Hub pelo terminal**:

    ```bash
    docker login
    ```

    Você será solicitado a inserir seu nome de usuário e senha do Docker Hub.

2.  **Envie a imagem com o comando `docker push`**:

    ```bash
    docker push seu-usuario-dockerhub/gerador-saudacoes:1.0
    ```

    O Docker irá verificar a imagem que você construiu localmente e enviará suas camadas para o seu repositório no Docker Hub.

3.  **Verifique no site:** Após o push ser concluído, vá ao seu perfil no [Docker Hub](https://hub.docker.com/). Você verá o novo repositório `gerador-saudacoes` com a tag `1.0` disponível publicamente (ou privadamente, dependendo das suas configurações).

Parabéns\! Agora qualquer pessoa com Docker pode baixar e rodar sua aplicação com um simples `docker run seu-usuario-dockerhub/gerador-saudacoes:1.0`.