A seguir, apresento um guia passo a passo detalhado para construir um `Dockerfile` otimizado para seu projeto, criar a imagem Docker e enviá-la para o Docker Hub, seguindo as melhores práticas.

### Pré-requisitos

1.  **Docker Instalado:** Você precisa ter o Docker instalado e em execução em sua máquina.
2.  **Conta no Docker Hub:** Crie uma conta gratuita no [Docker Hub](https://hub.docker.com/).

-----

## Passo 1: Criar o arquivo `.dockerignore`

Assim como o `.gitignore`, o `.dockerignore` impede que arquivos e diretórios desnecessários sejam enviados para o daemon do Docker durante o processo de build. Isso acelera o build e resulta em uma imagem menor.

Crie um arquivo chamado `.dockerignore` na raiz do seu projeto com o seguinte conteúdo:

```
# Arquivos do Docker
Dockerfile
.dockerignore

# Ambiente virtual e cache do Python
.venv/
__pycache__/
*.pyc
*.pyo
*.pyd
.pytest_cache/
.ruff_cache/

# Banco de dados local
*.db

# Arquivos de configuração de ambiente de desenvolvimento
.devbox/
devbox.lock
```

## Passo 2: Criar o `Dockerfile` com Multi-Stage Build

Usaremos uma abordagem de *multi-stage build* (build em múltiplos estágios). Essa é uma prática recomendada por dois motivos principais:

1.  **Segurança:** Ferramentas de compilação não são incluídas na imagem final.
2.  **Tamanho da Imagem:** A imagem final é significativamente menor porque contém apenas o necessário para executar a aplicação.

Crie um arquivo chamado `Dockerfile` na raiz do projeto com o seguinte conteúdo:

```dockerfile
# --- Estágio 1: Builder ---
# Usamos a imagem Python 3.13 na versão 'slim' como base para o build.
# 'slim' é menor que a imagem padrão.
# Nomeamos este estágio como 'builder'.
FROM python:3.13-slim AS builder

# Define o diretório de trabalho dentro do container.
WORKDIR /app

# Atualiza o pip para a versão mais recente.
RUN pip install --upgrade pip

# Copia apenas o arquivo de dependências para o container.
# Isso aproveita o cache de camadas do Docker. Se o requirements.txt não mudar,
# o Docker reutilizará a camada cacheada da instalação de dependências.
COPY requirements.txt .

# Instala as dependências como 'wheels' em um diretório separado.
# Wheels são um formato pré-compilado que torna a instalação no próximo estágio mais rápida.
RUN pip wheel --no-cache-dir --wheel-dir /app/wheels -r requirements.txt


# --- Estágio 2: Final ---
# Começamos uma nova imagem, do zero, baseada na mesma versão do Python.
# Isso garante que a imagem final seja limpa e contenha apenas o necessário.
FROM python:3.13-slim

# Define o diretório de trabalho.
WORKDIR /app

# Copia as dependências pré-compiladas (wheels) do estágio 'builder'.
COPY --from=builder /app/wheels /app/wheels

# Copia o código-fonte da aplicação para o container.
COPY . .

# Instala as dependências a partir dos wheels locais.
# Isso é mais rápido e não requer acesso à internet.
RUN pip install --no-cache-dir /app/wheels/*

# Expõe a porta 8000, que é a porta padrão onde o Uvicorn irá rodar.
EXPOSE 8000

# Define o comando para iniciar a aplicação quando o container for executado.
# Usa "0.0.0.0" para que a API seja acessível de fora do container.
# Não usamos '--reload' em produção.
# Os arquivos do projeto (main.py, database.py, etc.) estão em /app
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
```

### Explicação Detalhada do `Dockerfile`

  * **`FROM python:3.13-slim as builder`**: Inicia o primeiro estágio. Usamos a imagem oficial do Python 3.13 na variante `slim`, que é otimizada em tamanho. O `as builder` nomeia este estágio.
  * **`WORKDIR /app`**: Define o diretório de trabalho padrão para `/app`. Todos os comandos subsequentes serão executados a partir deste diretório.
  * **`COPY requirements.txt .`**: Copia apenas o `requirements.txt`. Ao separar a cópia das dependências da cópia do código-fonte, aproveitamos o cache do Docker. Se você alterar apenas o código da aplicação, o Docker não precisará reinstalar as dependências.
  * **`RUN pip wheel ...`**: Este comando compila as dependências listadas no `requirements.txt` em arquivos `.whl` (wheels) e os salva no diretório `/app/wheels`.
  * **`FROM python:3.13-slim`**: Inicia o segundo e final estágio, descartando o estágio `builder` e todo o seu conteúdo, exceto o que for explicitamente copiado.
  * **`COPY --from=builder /app/wheels /app/wheels`**: Copia o diretório `/app/wheels` (com as dependências compiladas) do estágio `builder` para o estágio final.
  * **`COPY . .`**: Copia todo o conteúdo do diretório do projeto (respeitando o `.dockerignore`) para o `WORKDIR` (`/app`) no container.
  * **`RUN pip install --no-cache-dir /app/wheels/*`**: Instala as dependências a partir dos arquivos `wheel` locais. É mais rápido e não baixa nada da internet.
  * **`EXPOSE 8000`**: Informa ao Docker que o container escutará na porta 8000. Isso é principalmente para documentação; a publicação da porta é feita no comando `docker run`.
  * **`CMD ["uvicorn", "main:app", ...]`**: Define o comando padrão para executar a aplicação FastAPI usando o servidor Uvicorn.

## Passo 3: Construir a Imagem Docker

Abra seu terminal na raiz do projeto (onde estão o `Dockerfile` e os outros arquivos) e execute o comando a seguir. Substitua `seu-usuario-dockerhub` pelo seu nome de usuário no Docker Hub.

```bash
docker build -t seu-usuario-dockerhub/ms-pessoas-aleatorias:1.0 .
```

  * **`docker build`**: O comando para construir uma imagem.
  * **`-t seu-usuario-dockerhub/ms-pessoas-aleatorias:1.0`**: A flag `-t` (tag) nomeia a imagem. O formato `usuario/nome-da-imagem:versao` é o padrão para o Docker Hub.
  * **`.`**: Indica que o contexto do build (os arquivos a serem enviados para o daemon do Docker) é o diretório atual.

## Passo 4: Executar e Testar o Container Localmente

Após o build ser concluído com sucesso, você pode executar um container a partir da sua nova imagem:

```bash
docker run -d -p 8000:8000 --name api-pessoas seu-usuario-dockerhub/ms-pessoas-aleatorias:1.0
```

  * **`docker run`**: O comando para executar um container.
  * **`-d`**: (detach) Executa o container em segundo plano.
  * **`-p 8000:8000`**: Mapeia a porta 8000 do seu computador (host) para a porta 8000 do container.
  * **`--name api-pessoas`**: Dá um nome amigável ao container para facilitar o gerenciamento.

Agora, você pode testar sua API:

1.  **Documentação Interativa:** Abra seu navegador e acesse `http://localhost:8000/docs`.
2.  **Endpoint de Sorteio:** Acesse `http://localhost:8000/pessoas/aleatoria/` para obter uma pessoa aleatória. O banco de dados será populado na inicialização, conforme definido em `main.py`.

Para parar o container, use: `docker stop api-pessoas`.

## Passo 5: Enviar a Imagem para o Docker Hub

Finalmente, vamos enviar sua imagem para que ela possa ser baixada e usada em qualquer lugar.

1.  **Faça login no Docker Hub pelo terminal:**

    ```bash
    docker login
    ```

    Você precisará inserir seu nome de usuário e senha (ou um token de acesso).

2.  **Envie a imagem:**
    Use o mesmo nome e tag que você definiu no comando `docker build`.

    ```bash
    docker push seu-usuario-dockerhub/ms-pessoas-aleatorias:1.0
    ```

Após o push ser concluído, você poderá ver sua imagem pública (ou privada, dependendo das suas configurações) em seu perfil no site do Docker Hub. Agora, qualquer pessoa com Docker pode executar sua aplicação com um simples `docker run`\!