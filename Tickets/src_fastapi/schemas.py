from pydantic import BaseModel
from typing import Optional

class TicketInput(BaseModel):
    price: int = ''
    row: Optional[str] = None
    seat_number: Optional[int] = None
    vip: bool = False
    user_id: int = ''
    event_id: int = ''