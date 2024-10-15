from sqlalchemy import Column, Integer, String, Boolean, DateTime
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

# Event model (SQLAlchemy)
class Event(Base):
    __tablename__ = "Events"
    ID = Column(Integer, primary_key=True, index=True)
    name = Column(String(255), nullable=False)
    location = Column(String(255), nullable=True)
    organisator = Column(String(255), nullable=True)
    startdate = Column(DateTime, nullable=True)
    available_normal_tickets = Column(Integer, nullable=True)
    available_vip_tickets = Column(Integer, nullable=True)
    canceled = Column(Boolean, default=False)
