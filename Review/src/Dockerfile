FROM python:3.10
ENV PYTHONUNBUFFERED 1
WORKDIR /app
COPY /src /app/
# COPY ./ssl /app/ssl
RUN pip install -r requirements.txt
ENTRYPOINT [ "python", "src/main.py" ]
CMD ["uvicorn", "run", "--host", "0.0.0.0", "--reload"]