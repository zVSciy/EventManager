from sqlalchemy import Column, Integer, String, Boolean
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column
from typing import Optional

class Base(DeclarativeBase):
    pass

class Ticket(Base):
    __tablename__ = "tickets"
    id: Mapped[int] = mapped_column(Integer(), primary_key=True)
    price: Mapped[int] = mapped_column(Integer(), nullable=False)
    paid: Mapped[bool] = mapped_column(Boolean(), default=False)
    row: Mapped[str] = mapped_column(String(1), nullable=True)
    seat_number: Mapped[int] = mapped_column(Integer(), nullable=True)
    vip: Mapped[bool] = mapped_column(Boolean(), nullable=False)
    user_id: Mapped[int] = mapped_column(Integer(), nullable=False)
    event_id: Mapped[int] = mapped_column(Integer(), nullable=False)

    def __repr__(self):
        return f'Ticket(id={self.id}, price={self.price}, paid={self.paid}, row={self.row}, seat_number={self.seat_number}, vip={self.vip}, user_id={self.user_id}, event_id={self.event_id})'