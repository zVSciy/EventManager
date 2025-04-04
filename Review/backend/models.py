from sqlalchemy import Column, Integer, String, CheckConstraint
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class Review(Base):
    __tablename__ = 'reviews'
    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(String, nullable=False)
    content = Column(String, nullable=False)
    rating = Column(Integer, nullable=False)
    event_id = Column(Integer, nullable=False)
    
    __table_args__ = (
        CheckConstraint('rating >= 1 AND rating <= 5', name='rating_check'),
    )