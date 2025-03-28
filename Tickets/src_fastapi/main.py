from fastapi import FastAPI, HTTPException, Depends, status
from model.database import DBSession, get_db
from model import models
from schemas import TicketInput
from sqlalchemy.orm.exc import UnmappedInstanceError
from sqlalchemy.orm import Session

app = FastAPI()

@app.get("/tickets")
def read_tickets(event_id: int = None, db: Session = Depends(get_db)):
    try:
        if type(event_id) == int:
            tickets = db.query(models.Ticket).filter(models.Ticket.event_id == event_id).all()
        else:
            tickets = db.query(models.Ticket).all()
        return tickets

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })

@app.get("/tickets/user/{user_id}")
def read_tickets(user_id: str, event_id: int = None, db: Session = Depends(get_db)):
    try:
        if type(event_id) == int:
            tickets = db.query(models.Ticket).filter(models.Ticket.user_id == user_id).filter(models.Ticket.event_id == event_id).all()
        else:
            tickets = db.query(models.Ticket).filter(models.Ticket.user_id == user_id).all()
        return tickets

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })

@app.get("/tickets/{ticket_id}")
def read_tickets(ticket_id: int = None, db: Session = Depends(get_db)):
    try:
        tickets = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
        return tickets

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })
        
@app.post("/tickets")
def add_ticket(ticket: TicketInput, db: Session = Depends(get_db)):
    try: 
        price_max_length = models.Ticket.price.property.columns[0].type.display_width
        row_max_length = models.Ticket.row.property.columns[0].type.length
        seat_number_max_length = models.Ticket.seat_number.property.columns[0].type.display_width

        if len(str(ticket.price)) > price_max_length: #Check if price is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `price` is too long. This attribute mustn't be longer than {price_max_length} numbers."
            })

        if ticket.row and len(ticket.row) > row_max_length: #Check if row is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `row` is too long. This attribute mustn't be longer than {row_max_length} character(s)."
            })

        if ticket.seat_number and len(str(ticket.seat_number)) > seat_number_max_length: #Check if seat_number is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `seat_number` is too long. This attribute mustn't be longer than {seat_number_max_length} numbers."
            })

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

    except HTTPException as http_ex:
        raise http_ex

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })

@app.put("/tickets/{ticket_id}")
def edit_ticket(ticket_id: int, updated_ticket: TicketInput, db: Session = Depends(get_db)):
    try:
        ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
        if not ticket:
            raise HTTPException(status_code=404, detail={
                "status": "Error 404 - Not Found",
                "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
            })

        price_max_length = models.Ticket.price.property.columns[0].type.display_width
        row_max_length = models.Ticket.row.property.columns[0].type.length
        seat_number_max_length = models.Ticket.seat_number.property.columns[0].type.display_width
        
        if len(str(updated_ticket.price)) > price_max_length: #Check if price is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `price` is too long. This attribute mustn't be longer than {price_max_length} numbers."
            })

        if len(str(updated_ticket.row)) > row_max_length: #Check if row is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `row` is too long. This attribute mustn't be longer than {row_max_length} character(s)."
            })

        if len(str(updated_ticket.seat_number)) > seat_number_max_length: #Check if seat_number is too long (max length in MySQL database)
            raise HTTPException(status_code=400, detail={
                "status": "Error 400 - Bad Request",
                "msg": f"Attribute `seat_number` is too long. This attribute mustn't be longer than {seat_number_max_length} numbers."
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

    except HTTPException as http_ex:
        raise http_ex

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })

@app.delete("/tickets/user/{user_id}/ticket/{ticket_id}")
def delete_ticket(user_id: str, ticket_id: int, db: Session = Depends(get_db)):
    try:
        ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).filter(models.Ticket.user_id == user_id).first()
        if not ticket:
            raise HTTPException(status_code=403, detail={
                "status": "Error 403 - Blocked",
                "msg": f"Ticket with `id`: `{ticket_id}` is not assigned to user with `id`: `{user_id}`."
            })
        
        db.delete(ticket)
        db.commit()

        return {
            "status": "200",
            "msg": "Ticket cancelled successfully."
        }
        
    except UnmappedInstanceError:
        raise HTTPException(status_code=404, detail={
            "status": "Error 404 - Not Found",
            "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
        })

@app.delete("/tickets/{ticket_id}")
def delete_ticket(ticket_id: int, db: Session = Depends(get_db)):
    try:
        ticket = db.query(models.Ticket).filter(models.Ticket.id == ticket_id).first()
        db.delete(ticket)
        db.commit()

        return {
            "status": "200",
            "msg": "Ticket cancelled successfully."
        }
        
    except UnmappedInstanceError:
        raise HTTPException(status_code=404, detail={
            "status": "Error 404 - Not Found",
            "msg": f"Ticket with `id`: `{ticket_id}` doesn't exist."
        })

    except Exception as ex:
        raise HTTPException(status_code=500, detail={
            "status": "Error 500 - Internal Server Error",
            "msg": str(ex)
        })
    