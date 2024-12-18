from pydantic import BaseModel, EmailStr

class UserInput(BaseModel):
    email: EmailStr
    password: str
    first_name: str
    last_name: str
    role: str = "user"
