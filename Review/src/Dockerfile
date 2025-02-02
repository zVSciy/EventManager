FROM python:3.10

COPY src/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . code
WORKDIR /code/src

EXPOSE 8083
CMD uvicorn main:app --reload