FROM python:3

COPY frontend/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . code
WORKDIR /code


ENTRYPOINT ["python", "frontend/app.py"]
CMD ["flask", "run", "--host=0.0.0.0", "--port=5000"]

