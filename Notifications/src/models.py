from sqlalchemy import Column, Integer, String, Text, TIMESTAMP
from sqlalchemy.orm import declarative_base
from database import engine
from datetime import datetime

Base = declarative_base()

class Notifications(Base):
    __tablename__ = "notifications"
    
    id = Column(Integer, primary_key=True, index=True)
    description = Column(Text(1000), nullable=False)
    status = Column(String(20), nullable=False)
    timestamp = Column(Integer, nullable=False)
    eventId = Column(Integer, nullable=False)
    paymentId = Column(Integer, nullable=False)
    ticketId = Column(Integer, nullable=False)
    userId = Column(String(50), nullable=False)

Base.metadata.create_all(bind=engine)