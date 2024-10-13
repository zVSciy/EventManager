from pydantic import BaseModel
from typing import List, Optional



# Pydantic models (for validation)
class EventCreate(BaseModel):
    name: str
    location: Optional[str] = None
    organisator: Optional[str] = None
    startdate: Optional[str] = None
    available_normal_tickets: Optional[int] = 0
    available_vip_tickets: Optional[int] = 0

class EventUpdate(BaseModel):
    name: Optional[str] = None
    location: Optional[str] = None
    organisator: Optional[str] = None
    startdate: Optional[str] = None
    available_normal_tickets: Optional[int] = None
    available_vip_tickets: Optional[int] = None

class EventCancel(BaseModel):
    canceled: bool

class TicketsUpdate(BaseModel):
    available_normal_tickets: Optional[int] = None
    available_vip_tickets: Optional[int] = None