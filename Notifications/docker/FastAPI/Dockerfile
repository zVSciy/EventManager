FROM python:3.10
ENV PYTHONUNBUFFERED 1
WORKDIR /app
COPY /src /app/
# COPY ./ssl /app/ssl
RUN pip install -r requirements.txt
CMD uvicorn main:app --reload