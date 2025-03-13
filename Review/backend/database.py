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

# Remove the database file if it exists
def clean_db_file():
    if os.path.exists(db_path):
        try:
            logger.info(f"Removing existing database file at {db_path}")
            os.remove(db_path)
        except Exception as e:
            logger.error(f"Error removing database file: {e}")

SQLALCHEMY_DATABASE_URL = f"sqlite:///{db_path}"

engine = create_engine(SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def init_db():
    logger.info("Initializing the database")
    
    # Remove the old database file
    clean_db_file()
    
    # Create all tables
    logger.info("Creating tables")
    Base.metadata.create_all(bind=engine)
    
    # Add sample data
    db = SessionLocal()
    try:
        logger.info("Adding sample data")
        reviews = [
            {"user_id": 1, "content": "Great event!", "rating": 5, "event_id": 9999},
        ]
        for review in reviews:
            db.add(Review(**review))
        db.commit()
        logger.info("Sample data added successfully")
    except Exception as e:
        logger.error(f"Error adding sample data: {e}")
        db.rollback()
    finally:
        db.close()