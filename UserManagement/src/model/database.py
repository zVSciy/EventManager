import os
from dotenv import load_dotenv
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

# load_dotenv()
user = os.environ.get("MYSQL_USER")
database = os.environ.get("MYSQL_DATABASE")
password = os.environ.get("MYSQL_PASSWORD")
hostname = os.environ.get("MYSQL_DATABASE_HOST")
port = os.environ.get("MYSQL_DATABASE_PORT")

engineStr = f'mysql+pymysql://{user}:{password}@{hostname}:{port}/{database}?charset=utf8mb4'
print(engineStr)

engine = create_engine(engineStr, echo=True)

DBSession = sessionmaker(engine, autoflush=False)

def get_db(): # a generator function that yields a generator of db
    db = DBSession()
    try:
        yield db
    finally:
        db.close()