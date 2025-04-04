from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal, engine, Base
from models import Notifications
from datetime import datetime
from inputModels import BaseNotification, BaseNotificationOptional
from fastapi.middleware.cors import CORSMiddleware
import os
import requests  # Import requests to make HTTP calls to the tickets service
import logging  # Import logging for detailed error logging

app = FastAPI()

origins = [
    "*"
]

os.environ.NODE_TLS_REJECT_UNAUTHORIZED = '0';
TICKETS_SERVICE_URL =  "https://tickets_api:8000"

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Configure logging
logging.basicConfig(level=logging.ERROR)
logger = logging.getLogger(__name__)

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
     # Fetch tickets from the tickets service
    try:
        # Disable SSL verification for development purposes
        response = requests.get(f"{TICKETS_SERVICE_URL}/tickets/user/{user_id}", verify=False)
        if response.status_code != 200:
            logger.error(f"Tickets service returned status {response.status_code}: {response.text}")
            raise HTTPException(status_code=500, detail="Failed to fetch tickets from tickets service")

        # Store the JSON response in a variable to avoid consuming the body multiple times
        tickets = response.json()
    except requests.exceptions.RequestException as e:
        logger.error(f"Error connecting to tickets service: {e}")
        raise HTTPException(status_code=500, detail="Error connecting to tickets service")
    except Exception as e:
        logger.error(f"Unexpected error: {e}")
        raise HTTPException(status_code=500, detail="Internal Server Error")
    
    # Sync notifications with tickets
    for ticket in tickets:
        # Map the ticket model to the notification model
        existing_notification = db.query(Notifications).filter(Notifications.ticketId == ticket["id"]).first()
        if not existing_notification:
            # Create a new notification for the ticket
            new_notification = Notifications(
                userId=ticket["user_id"],
                ticketId=ticket["id"],
                eventId=ticket["event_id"],
                description=f"Ticket for event {ticket['event_id']} is available.",
                status="unread",
                timestamp=int(0),
                paymentId=0  # Assuming paymentId is not available in the ticket model
            )
            db.add(new_notification)
    
    db.commit()

    # Return all notifications for the user
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