from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
import models

# Database configuration
DATABASE_URL = "mysql+pymysql://root@localhost:3306/EventManagement"

# Initialize database connection
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

models.Base.metadata.create_all(bind=engine)