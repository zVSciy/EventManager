from sqlalchemy import create_engine, URL
from sqlalchemy.orm import sessionmaker
import os

#import env variables
db_user = os.environ.get('MYSQL_USER', 'test_user')
db_password = os.environ.get('MYSQL_PASSWORD', 'test_pw')
db_host = os.environ.get('MYSQL_DATABASE_HOST', 'db')
db_port = os.environ.get('MYSQL_DATABASE_PORT', 3306)
db_database = os.environ.get('MYSQL_DATABASE', 'api_db')

SQLALCHEMY_DB_URL = f"mysql+pymysql://{db_user}:{db_password}@{db_host}:{db_port}/{db_database}"

engine = create_engine(SQLALCHEMY_DB_URL, echo=True)
DBSession = sessionmaker(engine, autoflush=False)

def get_db(): # a generator function that yields a generator of db
    db = DBSession()
    try:
        yield db
    finally:
        db.close()
