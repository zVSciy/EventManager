from fastapi import FastAPI, HTTPException, Depends
from model.models import Event
from model.database import SessionLocal
from schemas import EventCancel, EventCreate, EventUpdate, TicketsUpdate
from sqlalchemy.orm import  Session
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],          
    allow_credentials=True,         
    allow_methods=["*"],            
    allow_headers=["*"],            
)


# Dependency to get DB session
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# API Endpoints
@app.get("/event/")
def get_all_events(db: Session = Depends(get_db)):
    events = db.query(Event).all()
    return events

@app.get("/event/{event_id}")
def get_event(event_id: int, db: Session = Depends(get_db)):
    event = db.query(Event).filter(Event.ID == event_id).first()
    if event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    return event

@app.get("/event/available_tickets/{event_id}")
def get_event_tickets(event_id: int, db: Session = Depends(get_db)):
    event = db.query(Event).filter(Event.ID == event_id).first()
    if event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    return {
        "ID": event.ID,
        "available_normal_tickets": event.available_normal_tickets,
        "available_vip_tickets": event.available_vip_tickets
    }

@app.post("/event/")
def create_event(event: EventCreate, db: Session = Depends(get_db)):
    new_event = Event(
        name=event.name,
        location=event.location,
        organisator=event.organisator,
        startdate=event.startdate,
        available_normal_tickets=event.available_normal_tickets,
        available_vip_tickets=event.available_vip_tickets
    )
    db.add(new_event)
    db.commit()
    db.refresh(new_event)
    return {"code": 200, "response": "Event was created successfully", "eventID": new_event.ID}

@app.put("/event/{event_id}/")
def update_event(event_id: int, event: EventUpdate, db: Session = Depends(get_db)):
    db_event = db.query(Event).filter(Event.ID == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    
    for key, value in event.dict(exclude_unset=True).items():
        setattr(db_event, key, value)
    
    db.commit()
    db.refresh(db_event)
    return {"code": 200, "response": "Event was updated successfully", "eventID": db_event.ID}

@app.put("/event/cancel/{event_id}")
def cancel_event(event_id: int, event: EventCancel, db: Session = Depends(get_db)):
    db_event = db.query(Event).filter(Event.ID == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    
    db_event.canceled = event.canceled
    db.commit()
    db.refresh(db_event)
    return {"code": 200, "response": "Event was canceled successfully", "eventID": db_event.ID}

@app.put("/event/updateTicket/{event_id}")
def update_event_tickets(event_id: int, tickets: TicketsUpdate, db: Session = Depends(get_db)):
    db_event = db.query(Event).filter(Event.ID == event_id).first()
    if db_event is None:
        raise HTTPException(status_code=404, detail="Event not found")
    if tickets.available_normal_tickets == 1:
        db_event.available_normal_tickets += 1
        print(db_event.available_normal_tickets)
    elif tickets.available_normal_tickets ==0 and (db_event.available_normal_tickets-1)>=0:
        db_event.available_normal_tickets -= 1
    elif tickets.available_normal_tickets ==2:
        print(db_event.available_normal_tickets)
    else:
        return {"status": 400, "response": "Not enough tickets available"}
    if tickets.available_vip_tickets == 1:
        db_event.available_vip_tickets += 1
        print(db_event.available_vip_tickets)
    elif tickets.available_vip_tickets ==0 and (db_event.available_vip_tickets-1)>=0:
        db_event.available_vip_tickets -= 1
    elif tickets.available_vip_tickets ==2:
        print(db_event.available_vip_tickets)
    else:
        return {"status": 400, "response": "Not enough tickets available"}
    
    db.commit()
    db.refresh(db_event)
    return {"status": 200, "response": "Available tickets were updated successfully", "eventID": db_event.ID}
