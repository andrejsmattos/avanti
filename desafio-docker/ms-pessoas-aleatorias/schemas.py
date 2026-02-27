# schemas.py

from pydantic import BaseModel

# Schema base para uma pessoa, contendo apenas o nome.
# Usado para criar uma nova pessoa.
class PessoaCreate(BaseModel):
    nome: str

# Schema completo para uma pessoa, incluindo o ID.
# Usado para retornar dados da API.
class Pessoa(BaseModel):
    id: int
    nome: str

    # Configuração para permitir que o Pydantic leia dados de um modelo ORM.
    class Config:
        from_attributes = True
