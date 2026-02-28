A seguir, apresento um guia passo a passo detalhado para construir um `Dockerfile` otimizado para seu projeto, criar a imagem Docker e enviá-la para o Docker Hub, seguindo as melhores práticas.

### Pré-requisitos

1. **Docker instalado:** Certifique-se de que o Docker Desktop ou o Docker Engine esteja instalado e em execução em sua máquina.
2. **Conta no Docker Hub:** Você precisará de uma conta no Docker Hub para enviar a imagem. Crie uma, caso ainda não tenha.

-----

### Passo 1: Criando o `Dockerfile`

A melhor prática para criar imagens de contêineres para aplicações compiladas como Go é usar uma **construção multi-stage (multi-stage build)**. Isso resulta em uma imagem final muito menor, contendo apenas o binário executável e os arquivos necessários, sem as ferramentas de compilação ou o código-fonte.

No diretório raiz do seu projeto (`ms-saudacoes-aleatorias`), crie um arquivo chamado `Dockerfile` com o seguinte conteúdo:

```dockerfile
# --- Estágio de Build ---
# Use uma imagem do Go baseada em Alpine. A versão pode ser ajustada.
FROM golang:1.24-alpine AS builder

# Instale as ferramentas de compilação C (gcc, etc.)
# 'build-base' é um meta-pacote que inclui o necessário em Alpine.
RUN apk add --no-cache build-base gcc

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos de gerenciamento de dependências primeiro para aproveitar o cache do Docker
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código-fonte da sua aplicação
COPY . .

# Compile a aplicação com CGO habilitado.
# -a: Força a reconstrução de pacotes que estão desatualizados.
# -installsuffix cgo: Evita conflitos entre pacotes CGO e não-CGO.
# O resultado é um binário estaticamente vinculado, que não precisa da libsqlite3.so na imagem final.
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# --- Estágio Final ---
# Define a imagem final, que será muito menor.
# A imagem 'alpine' é uma distribuição Linux mínima, ideal para produção.
FROM alpine:latest

# Define o diretório de trabalho na imagem final.
WORKDIR /app

# Copia o binário compilado do estágio 'builder' para a imagem final.
COPY --from=builder /app/main .

# Copia o banco de dados inicial, se necessário.
# Como o banco de dados 'greetings.db' é criado na primeira execução, não precisamos copiá-lo.
# A aplicação irá criá-lo quando o contêiner iniciar.

# Expõe a porta que a aplicação usa, conforme definido no main.go.
EXPOSE 8080

# Define o comando que será executado quando o contêiner for iniciado.
# Este é o ponto de entrada da nossa aplicação.
CMD ["./main"]
```

### Passo 2: Construindo a Imagem Docker

Agora que o `Dockerfile` está criado, vamos usá-lo para construir a imagem da nossa aplicação.

1.  **Abra o terminal** na raiz do seu projeto.

2.  **Execute o comando de build:**

      * Substitua `seu-usuario-dockerhub` pelo seu nome de usuário no Docker Hub.
      * `ms-saudacoes-aleatorias` é o nome que daremos à imagem.
      * `:1.0` é a *tag*, representando a versão da imagem.

    <!-- end list -->

    ```bash
    docker build -t seu-usuario-dockerhub/ms-saudacoes-aleatorias:1.0 .
    ```

    O `.` no final do comando indica que o contexto de build é o diretório atual, onde se encontra o `Dockerfile`.

    Você verá o Docker executar cada passo definido no `Dockerfile`. Graças à construção multi-stage, a imagem final será significativamente menor do que o estágio de compilação.

### Passo 3: Executando o Contêiner Localmente

Antes de enviar a imagem para o Docker Hub, é uma boa prática testá-la localmente para garantir que tudo está funcionando como esperado.

1.  **Execute o contêiner a partir da imagem que você acabou de criar:**

    ```bash
    docker run -p 8080:8080 --name saudacoes-app seu-usuario-dockerhub/ms-saudacoes-aleatorias:1.0
    ```

      * `-p 8080:8080`: Mapeia a porta 8080 do seu host para a porta 8080 do contêiner.
      * `--name saudacoes-app`: Dá um nome amigável ao seu contêiner para facilitar o gerenciamento.

2.  **Teste a aplicação:**

    Abra um novo terminal e use o `curl` para testar os endpoints, conforme descrito no `README.md`:

      * **Obter uma saudação aleatória:**

        ```bash
        curl http://localhost:8080/api/saudacoes/aleatorio
        ```

        Você deverá receber uma resposta como: `{"saudação":"Olá"}`.

      * **Cadastrar uma nova saudação:**

        ```bash
        curl -X POST \
          -H "Content-Type: application/json" \
          -d '{"text":"Docker é incrível"}' \
          http://localhost:8080/api/saudacoes
        ```

3.  **Pare e remova o contêiner** de teste após a verificação:

    ```bash
    docker stop saudacoes-app
    docker rm saudacoes-app
    ```

### Passo 4: Enviando a Imagem para o Docker Hub

Com a imagem construída e testada, o último passo é enviá-la para o Docker Hub para que possa ser compartilhada e usada em outros ambientes.

1.  **Faça login no Docker Hub** através do terminal:

    ```bash
    docker login
    ```

    Você será solicitado a fornecer seu nome de usuário e senha do Docker Hub.

2.  **Envie a imagem:**

    Use o comando `docker push` com o mesmo nome e tag que você usou para construir a imagem.

    ```bash
    docker push seu-usuario-dockerhub/ms-saudacoes-aleatorias:1.0
    ```

    O Docker fará o upload das camadas da sua imagem para o repositório no Docker Hub.

### Conclusão

Pronto\! Você acaba de construir um `Dockerfile` otimizado, criar uma imagem Docker leve para sua aplicação Go e publicá-la no Docker Hub. Agora, qualquer pessoa com acesso ao Docker pode baixar e executar sua aplicação com um simples comando `docker run`.