from model import models
from schemas import UserInput
from sqlalchemy.orm import Session
from model.database import get_db
from security.hashing import pwd_context
from sqlalchemy.orm.exc import UnmappedInstanceError
from fastapi.security import OAuth2PasswordRequestForm
from fastapi import FastAPI, HTTPException, Depends, status
from security.authentication import create_access_token, verify_token

# Start FastAPI-Service: uvicorn main:app --reload
app = FastAPI()


@app.post("/register")
async def register_user(user_data: UserInput, db: Session = Depends(get_db)):
    hashed_password = pwd_context.hash(user_data.password)
    user = models.User(email=user_data.email, hashed_password=hashed_password)
    db.add(user)
    db.commit()
    db.refresh(user)
    return {"message": "User registered successfully", "user": user.email}


@app.post("/login")
async def login_for_access_token(form_data: OAuth2PasswordRequestForm = Depends(), 
                                 db: Session = Depends(get_db)):
    user = db.query(models.User).filter(models.User.email == form_data.username).first()

    if user and pwd_context.verify(form_data.password, user.hashed_password):
        access_token = create_access_token(data={"sub": user.email})
        return {"access_token": access_token, "token_type": "bearer"}

    raise HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Invalid credentials",
        headers={"WWW-Authenticate": "Bearer"},
    )