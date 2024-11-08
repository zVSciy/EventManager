import logging
from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, init_db, engine, Base
from models import Review
from pydantic import BaseModel

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = FastAPI()

# Initialize the database
init_db()

class ReviewCreate(BaseModel):
    user_id: int
    content: str
    rating: int
    event_id: int

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/reviews/{review_id}")
def get_review(review_id: int, db: Session = Depends(get_db)):
    logger.info(f"Fetching review with ID: {review_id}")
    review = db.query(Review).filter(Review.id == review_id).first()
    if review is None:
        logger.error(f"Review with ID {review_id} not found")
        raise HTTPException(status_code=404, detail="Review not found")
    return review

@app.post("/reviews/")
def create_review(review: ReviewCreate, db: Session = Depends(get_db)):
    logger.info(f"Creating review for user ID: {review.user_id}")
    db_review = Review(user_id=review.user_id, content=review.content, rating=review.rating, event_id=review.event_id)
    db.add(db_review)
    db.commit()
    db.refresh(db_review)
    logger.info(f"Review created successfully with ID: {db_review.id}")
    return db_review

@app.get("/reviews/{event_id}")
def get_reviews(event_id:int, db: Session = Depends(get_db)):
    logger.info(f"Fetching review with ID: {event_id}")
    reviews = db.query(Review).filter(Review.event_id == event_id).all()
    if reviews is None:
        raise HTTPException(status_code=404, detail="Reviews not found")
    return reviews

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)