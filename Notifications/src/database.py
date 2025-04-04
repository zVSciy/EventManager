from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base
from datetime import datetime
import os



MYSQL_DATABASE = os.getenv("NOTIF_MYSQL_DATABASE", "notif")
MYSQL_USER = os.getenv("NOTIF_MYSQL_USER", "NotificationAPI")
MYSQL_PASSWORD = os.getenv("NOTIF_MYSQL_PASSWORD", "IamNotificationPassword")
MYSQL_HOST = "notification_db"
MYSQL_PORT = os.getenv("NOTIF_MYSQL_DATABASE_PORT", 3306)

DATABASE_URL = f"mysql+pymysql://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}:{MYSQL_PORT}/{MYSQL_DATABASE}"

engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

Base.metadata.create_all(bind=engine)