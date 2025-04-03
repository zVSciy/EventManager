FROM python:3
ENV PYTHONUNBUFFERED 1
WORKDIR /app
COPY /src_fastapi /app
COPY /ssl /app/ssl
RUN pip install -r requirements.txt