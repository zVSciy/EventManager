from pydantic import BaseModel

class UserInput(BaseModel):
    email: str = ''
    password: str = ''
    first_name: str = ''
    last_name: str = ''
    role: str = ''