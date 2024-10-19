from pydantic import BaseModel

class TicketInput(BaseModel):
    price: int = ''
    row: str = ''
    seatNumber: int = ''
    vip: bool = False
    userid: int = ''
    eventid: int = ''