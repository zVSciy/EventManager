FROM python:3.10

COPY src/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . code
WORKDIR /code

EXPOSE 8083

ENTRYPOINT [ "python", "src/main.py" ]
CMD ["uvicorn", "run", "--host", "0.0.0.0", "--port=8000", "--reload", "--log-config", "/code/src/logging_config.yaml"]