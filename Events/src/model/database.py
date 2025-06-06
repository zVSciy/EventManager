from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from dotenv import load_dotenv
import os

load_dotenv()
user = os.environ.get("EVENTS_MYSQL_USER")
hostname = os.environ.get("EVENTS_MYSQL_HOST")
port = os.environ.get("EVENTS_MYSQL_DATABASE_PORT")
password = os.environ.get("EVENTS_MYSQL_ROOT_PASSWORD")

# Database configuration
DATABASE_URL = f"mysql+pymysql://root:{password}@{hostname}:{port}/EventManagement"
print(DATABASE_URL)

# Initialize database connection
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
