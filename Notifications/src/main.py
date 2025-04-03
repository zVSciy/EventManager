from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, engine, Base
from models import Notifications
from datetime import datetime
from inputModels import BaseNotification, BaseNotificationOptional
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI()

origins = [
    "*"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


# Create the database tables
Base.metadata.create_all(bind=engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/notifications/{notification_id}")
def read_notification(notification_id: int, db: Session = Depends(get_db)):
    notification = db.query(Notifications).filter(Notifications.id == notification_id).first()
    if(notification == None):
        raise HTTPException(status_code=404, detail="Notification not found")
    return notification

@app.get("/notifications/")
def read_all_notifactions(skip: int = 0, limit: int = 10, db: Session = Depends(get_db)):
    notifications = db.query(Notifications).offset(skip).limit(limit).all()
    if(notifications == None or len(notifications) == 0):
        raise HTTPException(status_code=404, detail="No notifications found")
    return notifications

@app.get("/notifications/user/{user_id}")
def read_all_notifactions(user_id: str, db: Session = Depends(get_db)):
    notifications = db.query(Notifications).filter(Notifications.userId == user_id).all()
    if(notifications == None or len(notifications) == 0):
        raise HTTPException(status_code=404, detail="No notifications found")
    return notifications


@app.post("/notifications/")
def create_notification(input_notification: BaseNotification, db: Session = Depends(get_db)):

    new_notification = Notifications(**input_notification.model_dump())
    db.add(new_notification)
    db.commit()
    db.refresh(new_notification)
    return new_notification

@app.put("/notifications/{notification_id}")
def update_notification(notification_id: int, input_notification: BaseNotificationOptional, db: Session = Depends(get_db)):
    # Evaluate inputs
    notification = db.query(Notifications).filter(Notifications.id == notification_id).first()
    if not notification:
        raise HTTPException(status_code=404, detail="Notification not found")

    # Unpack the input_notification
    for key, value in input_notification.dict(exclude_unset=True).items():
        setattr(notification, key, value)

    db.commit()
    db.refresh(notification)
    return notification

@app.delete("/notifications/{notification_id}")
def delete_notification(notification_id: int, db: Session = Depends(get_db)):
    notification = db.query(Notifications).filter(Notifications.id == notification_id).first()
    if not notification:
        raise HTTPException(status_code=404, detail="Notification not found")
    
    db.delete(notification)
    db.commit()
    return {"detail": "Notification deleted"}