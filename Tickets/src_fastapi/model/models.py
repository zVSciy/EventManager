from sqlalchemy import Column, Integer, String, Boolean
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column

class Base(DeclarativeBase):
    pass

class Ticket(Base):
    __tablename__ = "tickets"
    id: Mapped[int] = mapped_column(Integer(), primary_key=True)
    price: Mapped[int] = mapped_column(Integer(), nullable=False)
    paid: Mapped[bool] = mapped_column(Boolean(), default=False)
    row: Mapped[str] = mapped_column(String(1))
    seatNumber: Mapped[int] = mapped_column(Integer())
    vip: Mapped[bool] = mapped_column(Boolean(), nullable=False)
    userid: Mapped[int] = mapped_column(Integer(), nullable=False)
    eventid: Mapped[int] = mapped_column(Integer(), nullable=False)
