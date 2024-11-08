from sqlalchemy import Column, Integer, String, Boolean
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True)
    email = Column(String(50), unique=True)
    hashed_password = Column(String(60))
    first_name = Column(String(20))
    last_name = Column(String(50))
    role = Column(String(5), default='user') # admin, user