FROM python:3.10
ENV PYTHONUNBUFFERED 1

COPY src/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . code
WORKDIR /code


ENTRYPOINT [ "python", "src/main.py" ]
CMD ["uvicorn", "run", "--host", "0.0.0.0", "--reload"]