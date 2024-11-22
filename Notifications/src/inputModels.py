from pydantic import BaseModel
from typing import Optional


class BaseNotificationOptional(BaseModel):
    description: Optional[str]
    status: Optional[str]
    timestamp: Optional[int]
    eventId: Optional[int]
    paymentId: Optional[int]
    ticketId: Optional[int]
    userId: Optional[int]


class BaseNotification(BaseModel):
    description: str
    status: str
    timestamp: int
    eventId: int
    paymentId: int
    ticketId: int
    userId: int