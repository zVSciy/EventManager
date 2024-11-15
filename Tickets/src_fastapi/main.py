from fastapi import FastAPI, HTTPException, Depends, status
from model.database import DBSession, get_db
from model import models
from schemas import TicketInput
from sqlalchemy.orm.exc import UnmappedInstanceError
from sqlalchemy.orm import Session

app = FastAPI()

@app.get("/tickets")
def read_tickets(event_id: None, db: Session = Depends(get_db)):
    if event_id:
        tickets = db.query(models.Ticket).filter(models.Ticket.event_id == event_id).all()
    else:
        tickets = db.query(models.Ticket).all()
    return tickets

@app.post("/tickets")
def add_ticket(ticket:TicketInput, db: Session = Depends(get_db)):
    try:
        if len(ticket.price) == 0:
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": "Price is empty, but it must be provided."
            })
        
        if len(ticket.vip) == 0:
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": "VIP tag is empty, but it must be provided."
            })

        if len(ticket.userid) == 0:
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": "UserID is empty, but it must be provided."
            })

        if len(ticket.eventid) == 0:
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": "EventID is empty, but it must be provided."
            })

        new_ticket = models.Ticket(
            price=ticket.price,
            row=ticket.row,
            seatNumber=ticket.seatNumber,
            vip=ticket.vip,
            userid=ticket.userid,
            eventid=ticket.eventid,
        )
        db.add(new_ticket)
        db.commit()
        db.refresh(new_ticket)

    except:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Server Error",
            "msg": "Unexpected error occured during ticket creation - are all inputs correct?"
        })
    
    finally:
        db.close()
    return new_ticket

@app.put("/tickets/payment/{ticket-id}")
def payment_ticket(ticket_id: int, db: Session = Depends(get_db)):
    ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
    if not ticket:
        raise HTTPException(status_code=404, detail={
            "status": "Error 404 - Not Found",
            "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
        })
    ticket.paid = True
    db.commit()
    db.refresh(note)
    return ticket

@app.delete("/tickets/cancel/{ticket-id}")
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