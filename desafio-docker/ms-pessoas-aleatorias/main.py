# main.py

import random
from fastapi import FastAPI, Depends, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy.orm import Session
from sqlalchemy import func # Importamos 'func' para usar funções SQL

# Importamos nossos módulos criados
import models
import schemas
from database import SessionLocal, engine

# Cria as tabelas no banco de dados (só executa se elas não existirem)
models.Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="API de Pessoas Aleatórias",
    description="Uma API simples para cadastrar e sortear pessoas.",
    version="1.0.0"
)

# --- NOVA SEÇÃO: Configuração do CORS ---

# Lista de origens que têm permissão para fazer requisições
# Para desenvolvimento, "*" permite qualquer origem.
# Para produção, use uma lista explícita, ex: ["https://seu-dominio.com"]
origins = ["*"] 

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True, # Permite cookies (se você usar autenticação)
    allow_methods=["*"],    # Permite todos os métodos (GET, POST, etc.)
    allow_headers=["*"],    # Permite todos os cabeçalhos
)

# --- Lógica para Carga Inicial de Dados ---

NOMES_PREDEFINIDOS = [
    "Alice", "Bruno", "Carla", "Daniel", "Eva", "Fábio",
    "Gabriela", "Heitor", "Íris", "João", "Larissa", "Marcos"
]

@app.on_event("startup")
def popular_banco_de_dados():
    """
    Função executada na inicialização da API para popular o banco de dados.
    Verifica se os nomes pré-definidos já existem antes de inseri-los.
    """
    print("INFO:     Verificando e populando o banco de dados com dados iniciais...")
    db = SessionLocal()
    try:
        # Pega todos os nomes que já existem no banco para evitar duplicatas
        nomes_existentes = {nome[0] for nome in db.query(models.Pessoa.nome).all()}

        pessoas_para_adicionar = []
        for nome in NOMES_PREDEFINIDOS:
            if nome not in nomes_existentes:
                pessoas_para_adicionar.append(models.Pessoa(nome=nome))
        
        if pessoas_para_adicionar:
            db.add_all(pessoas_para_adicionar)
            db.commit()
            print(f"INFO:     {len(pessoas_para_adicionar)} novas pessoas adicionadas ao banco de dados.")
        else:
            print("INFO:     O banco de dados já está populado com os dados iniciais.")

    finally:
        db.close()

# --- Dependência para obter a sessão do banco de dados ---
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# --- Endpoints da API ---

@app.post("/pessoas/", response_model=schemas.Pessoa, status_code=201, tags=["Pessoas"])
def criar_pessoa(pessoa: schemas.PessoaCreate, db: Session = Depends(get_db)):
    """
    Cadastra uma nova pessoa no banco de dados.
    - **nome**: O nome da pessoa a ser cadastrada.
    """
    # Verifica se a pessoa já existe
    db_pessoa = db.query(models.Pessoa).filter(models.Pessoa.nome == pessoa.nome).first()
    if db_pessoa:
        raise HTTPException(status_code=400, detail="Pessoa com este nome já cadastrada.")

    nova_pessoa = models.Pessoa(nome=pessoa.nome)
    db.add(nova_pessoa)
    db.commit()
    db.refresh(nova_pessoa)
    return nova_pessoa

@app.get("/pessoas/aleatoria/", response_model=schemas.Pessoa, tags=["Pessoas"])
def obter_pessoa_aleatoria(db: Session = Depends(get_db)):
    """
    Retorna uma pessoa aleatória cadastrada no banco de dados.
    """
    # SQLAlchemy oferece a função 'func.random()' ou 'func.rand()' que é traduzida
    # para a função de aleatoriedade nativa do banco de dados (RANDOM() no SQLite).
    # Esta abordagem é muito mais eficiente do que carregar todos os registros para a memória.
    pessoa_aleatoria = db.query(models.Pessoa).order_by(func.random()).first()

    if not pessoa_aleatoria:
        raise HTTPException(status_code=404, detail="Nenhuma pessoa cadastrada no banco de dados.")

    return pessoa_aleatoria

@app.get("/", include_in_schema=False)
def root():
    return {"message": "Bem-vindo à API de Pessoas Aleatórias! Acesse /docs para ver a documentação."}
