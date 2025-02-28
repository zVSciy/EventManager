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
            {"user_id": 2, "content": "Not bad", "rating": 3, "event_id": 2},
            {"user_id": 3, "content": "Could be better", "rating": 2, "event_id": 1},
            {"user_id": 4, "content": "Loved it!", "rating": 5, "event_id": 2},
            {"user_id": 5, "content": "It was okay", "rating": 3, "event_id": 2},
            {"user_id": 6, "content": "Terrible experience", "rating": 1, "event_id": 3},
            {"user_id": 7, "content": "Pretty good", "rating": 4, "event_id": 3},
            {"user_id": 8, "content": "Not worth it", "rating": 2, "event_id": 4},
            {"user_id": 9, "content": "Fantastic!", "rating": 5, "event_id": 4}
        ]
        for review in reviews:
            db.add(Review(**review))
        db.commit()
    finally:
        db.close()