from pydantic import BaseModel

class TicketInput(BaseModel):
    price: int = ''
    row: str = ''
    seat_number: int = ''
    vip: bool = False
    user_id: int = ''
    event_id: int = ''