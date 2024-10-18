from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, init_db, engine, Base
from models import User, Review
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

@app.post("/users/")
def create_user(user: UserCreate, db: Session = Depends(get_db)):
    db_user = User(name=user.name, email=user.email)
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user

@app.post("/reviews/")
def create_review(review: ReviewCreate, db: Session = Depends(get_db)):
    db_user = db.query(User).filter(User.id == review.user_id).first()
    if not db_user:
        raise HTTPException(status_code=404, detail="User not found")
    db_review = Review(user_id=review.user_id, content=review.content, rating=review.rating)
    db.add(db_review)
    db.commit()
    db.refresh(db_review)
    return db_review

@app.get("/reviews/")
def get_reviews(db: Session = Depends(get_db)):
    reviews = db.query(Review).all()
    return reviews

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="127.0.0.1", port=8000)