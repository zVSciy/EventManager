from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base
from datetime import datetime
import os

MYSQL_DATABASE = os.getenv("MYSQL_DATABASE", "notif")
MYSQL_USER = os.getenv("MYSQL_USER", "NotificationAPI")
MYSQL_PASSWORD = os.getenv("MYSQL_PASSWORD", "IamNotificationPassword")
MYSQL_HOST = "mysqlDB"
MYSQL_PORT = os.getenv("MYSQL_DATABASE_PORT", 3306)

DATABASE_URL = f"mysql+pymysql://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}:{MYSQL_PORT}/{MYSQL_DATABASE}"

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

Base.metadata.create_all(bind=engine)