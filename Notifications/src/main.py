from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from database import SessionLocal, engine, Base
from models import Notifications
from datetime import datetime

app = FastAPI()

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
    return notification

@app.get("/notifications/")
def read_all_notifactions(skip: int = 0, limit: int = 10, db: Session = Depends(get_db)):
    notifications = db.query(Notifications).offset(skip).limit(limit).all()
    return notifications

@app.get("/notifications/user/{user_id}")
def read_all_notifactions(user_id: int, db: Session = Depends(get_db)):
    notifications = db.query(Notifications).filter(Notifications.userId == user_id).all()
    return notifications


@app.post("/notifications/")
def create_notification(description: str, status: str,timestamp: int, eventId: int, paymentId: int, ticketId: int, userId:int , db: Session = Depends(get_db)):
    new_notification = Notifications(description=description, status=status, timestamp=timestamp,eventId=eventId, paymentId=paymentId, ticketId=ticketId, userId=userId)
    db.add(new_notification)
    db.commit()
    db.refresh(new_notification)
    return new_notification

@app.put("/notifications/{notification_id}")
def update_notification(notification_id:int, description: str, status: str,timestamp: int, eventId: int, paymentId: int, ticketId: int, userId:int, db: Session = Depends(get_db)):
    notification = db.query(Notifications).filter(Notifications.id == notification_id).first()
    notification.description = description
    notification.status = status
    notification.timestamp = timestamp
    notification.eventId = eventId
    notification.paymentId = paymentId
    notification.ticketId = ticketId
    db.commit()
    db.refresh(notification)
    return notification