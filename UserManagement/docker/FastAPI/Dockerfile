FROM python:3.10
WORKDIR /app
COPY ./src /app
RUN pip install -r /app/requirements.txt
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

# FROM python:3.10
# ENV PYTHONUNBUFFERED 1
# WORKDIR /app
# COPY /src /app/
# COPY ./ssl /app/ssl
# RUN pip install -r requirements.txt
# CMD uvicorn main:app --reload

# FROM python:3.10
# WORKDIR /app
# COPY ./src /app
# COPY ./ssl /app/ssl
# RUN pip install -r requirements.txt
# CMD uvicorn main:app --host 0.0.0.0 --port 8000 --ssl-keyfile /app/ssl/certificate.key --ssl-certfile /app/ssl/certificate.crt