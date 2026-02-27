# database.py

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# URL de conexão com o banco de dados SQLite.
# O arquivo 'pessoas.db' será criado na raiz do projeto.
SQLALCHEMY_DATABASE_URL = "sqlite:///./pessoas.db"

# Cria a 'engine' do SQLAlchemy
engine = create_engine(
    SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
)

# Cria uma classe SessionLocal que será uma sessão do banco de dados
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

# Cria uma classe Base da qual nossos modelos ORM herdarão
Base = declarative_base()
