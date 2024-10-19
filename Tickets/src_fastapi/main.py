from fastapi import FastAPI, HTTPException, Depends, status
from model.database import DBSession, get_db
from model import models
from schemas import TicketInput
from sqlalchemy.orm.exc import UnmappedInstanceError
from sqlalchemy.orm import Session

app = FastAPI()

# @app.get("/tickets")
# ...
