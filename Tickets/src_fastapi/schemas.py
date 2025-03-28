from pydantic import BaseModel
from typing import Optional

class TicketInput(BaseModel):
    price: int = ''
    row: Optional[str] = None
    seat_number: Optional[int] = None
    vip: bool = False
    user_id: str = ''
    event_id: int = ''