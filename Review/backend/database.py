import os
import logging
from sqlalchemy import create_engine, inspect
from sqlalchemy.orm import sessionmaker
from models import Base, Review

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

db_dir = os.path.join(os.path.dirname(__file__), 'db')
os.makedirs(db_dir, exist_ok=True)

db_path = os.path.join(db_dir, 'reviews.sqlite')
SQLALCHEMY_DATABASE_URL = f"sqlite:///{db_path}"

engine = create_engine(SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def init_db():
    logger.info("Initializing the database and creating tables if they don't exist")
    Base.metadata.create_all(bind=engine)
    
    db = SessionLocal()
    try:
        inspector = inspect(engine)
        if not inspector.has_table("reviews"):
            logger.error("The reviews table does not exist.")
            return
        
        db.query(Review).delete()
        reviews = [
            {"user_id": 1, "content": "Great event!", "rating": 5, "event_id": 1},
        ]
        for review in reviews:
            db.add(Review(**review))
        db.commit()
    finally:
        db.close()