from fastapi import FastAPI, HTTPException, Depends, status
from model.database import DBSession, get_db
from model import models
from schemas import TicketInput
from sqlalchemy.orm.exc import UnmappedInstanceError
from sqlalchemy.orm import Session

app = FastAPI()

@app.get("/tickets")
def read_tickets(event_id: None, db: Session = Depends(get_db))
    if user_id:
        tickets = db.query(models.Ticket).filter(models.Ticket.event_id == event_id).all()
    else:
        tickets = db.query(models.Ticket).all()
    return tickets

@app.post("/tickets")
def add_ticket(ticket:TicketInput db: Session = Depends(get_db)):
    pass