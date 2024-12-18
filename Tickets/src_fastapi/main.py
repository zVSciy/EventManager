from fastapi import FastAPI, HTTPException, Depends, status
from model.database import DBSession, get_db
from model import models
from schemas import TicketInput
from sqlalchemy.orm.exc import UnmappedInstanceError
from sqlalchemy.orm import Session

app = FastAPI()

@app.get("/tickets")
def read_tickets(event_id: int = None, db: Session = Depends(get_db)):
    if event_id:
        tickets = db.query(models.Ticket).filter(models.Ticket.event_id == event_id).all()
    else:
        tickets = db.query(models.Ticket).all()
    return tickets

@app.post("/tickets")
def add_ticket(ticket: TicketInput, db: Session = Depends(get_db)):
    try:
        new_ticket = models.Ticket(
            price=ticket.price,
            row=ticket.row,
            seat_number=ticket.seat_number,
            vip=ticket.vip,
            user_id=ticket.user_id,
            event_id=ticket.event_id,
        )
        db.add(new_ticket)
        db.commit()
        db.refresh(new_ticket)
        return new_ticket

    except Exception:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Server Error",
            "msg": "Unexpected error occured during ticket creation - are all inputs correct?"
        })
    

@app.put("/tickets/{ticket_id}")
def edit_ticket(ticket_id: int, updated_ticket: TicketInput, db: Session = Depends(get_db)):
    ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
    if not ticket:
        raise HTTPException(status_code=404, detail={
            "status": "Error 404 - Not Found",
            "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
        })

    ticket.price = updated_ticket.price
    ticket.row = updated_ticket.row
    ticket.seat_number = updated_ticket.seat_number
    ticket.vip = updated_ticket.vip
    ticket.user_id = updated_ticket.user_id
    ticket.event_id = updated_ticket.event_id

    db.commit()
    db.refresh(ticket)
    return ticket

@app.delete("/tickets/{ticket_id}")
def delete_ticket(ticket_id: int, db: Session = Depends(get_db)):
    try:
        ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
        db.delete(ticket)
        db.commit()
    except UnmappedInstanceError:
        raise HTTPException(status_code=404, detail={
            "status": "Error 404 - Not Found",
            "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
        })
    return {
        "status": "200",
        "msg": "Ticket cancelled successfully."
    }