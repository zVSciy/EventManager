import os
from jose import jwt
from dotenv import load_dotenv
from datetime import datetime, timedelta

def create_access_token(data: dict):
    to_encode = data.copy()
    expire = datetime.utcnow() + timedelta(minutes=15)
    to_encode.update({"exp": expire})
    return jwt.encode(to_encode, os.environ.get("AUTH_SECRET_KEY"), algorithm=os.environ.get("AUTH_ALGORITHM"))