FROM python:3.10

COPY src/requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . code
WORKDIR /code

ENTRYPOINT [ "python", "src/main.py" ]
CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--reload", "--log-config", "/code/src/logging_config.yaml"]