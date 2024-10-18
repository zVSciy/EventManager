from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, init_db, engine, Base
from models import Review
from pydantic import BaseModel

app = FastAPI()

Base.metadata.create_all(bind=engine)

class UserCreate(BaseModel):
    name: str
    email: str

class ReviewCreate(BaseModel):
    user_id: int
    content: str
    rating: int

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Initialisiere die Datenbank synchron beim Start der Anwendung
init_db()

@app.get("/reviews/{review_id}")
def get_review(review_id: int, db: Session = Depends(get_db)):
    review = db.query(Review).filter(Review.id == review_id).first()
    if review is None:
        raise HTTPException(status_code=404, detail="Review not found")
    return review

@app.post("/reviews/")
def create_review(reviewId:int, user_id:int, content:str, rating:int , db: Session = Depends(get_db)):
    review = Review(id=reviewId, user_id=user_id, content=content, rating=rating)
    db.add(review)
    db.commit()
    db.refresh(review)
    return review


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)