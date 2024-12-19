import os
import pytest
import logging
from fastapi.testclient import TestClient
from main import app, get_db
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from models import Base, Review  # Ensure Review model is imported

# Logging configuration
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

os.makedirs("db", exist_ok=True)

SQLALCHEMY_DATABASE_URL = "sqlite:///./db/test.db"
engine = create_engine(SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False})
TestingSessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def override_get_db():
    try:
        db = TestingSessionLocal()
        yield db
    finally:
        db.close()

app.dependency_overrides[get_db] = override_get_db

client = TestClient(app)

@pytest.fixture(scope="session", autouse=True)
def setup_database():
    Base.metadata.create_all(bind=engine)
    yield

@pytest.fixture(scope="function", autouse=True)
def insert_reviews():
    # Insert the same set of reviews before each test
    db = TestingSessionLocal()
    db.query(Review).delete()
    reviews = [
        {"user_id": 1, "content": "Great event!", "rating": 5, "event_id": 1},
        {"user_id": 2, "content": "Not bad", "rating": 3, "event_id": 1},
        {"user_id": 3, "content": "Could be better", "rating": 2, "event_id": 1},
        {"user_id": 4, "content": "Loved it!", "rating": 5, "event_id": 2},
        {"user_id": 5, "content": "It was okay", "rating": 3, "event_id": 2},
        {"user_id": 6, "content": "Terrible experience", "rating": 1, "event_id": 3},
        {"user_id": 7, "content": "Pretty good", "rating": 4, "event_id": 3},
        {"user_id": 8, "content": "Not worth it", "rating": 2, "event_id": 4},
        {"user_id": 9, "content": "Fantastic!", "rating": 5, "event_id": 4}
    ]
    for review in reviews:
        db.add(Review(**review))
    db.commit()
    db.close()

@pytest.fixture(scope="module")
def test_client():
    yield client

def test_create_review(test_client):
    logger.info("Starting test_create_review")
    response = test_client.post("/reviews/", json={
        "user_id": 4,
        "content": "Amazing event!",
        "rating": 5,
        "event_id": 1
    })
    assert response.status_code == 200
    assert response.json()["content"] == "Amazing event!"

def test_get_review(test_client):
    logger.info("Starting test_get_review")
    response = test_client.get("/reviews/1")
    assert response.status_code == 200
    assert response.json()["content"] == "Great event!"

def test_update_review(test_client):
    logger.info("Starting test_update_review")
    response = test_client.put("/reviews/1", json={
        "content": "Updated content",
        "rating": 4
    })
    assert response.status_code == 200
    assert response.json()["content"] == "Updated content"

def test_delete_review(test_client):
    logger.info("Starting test_delete_review")
    response = test_client.delete("/reviews/1")
    assert response.status_code == 200
    assert response.json()["detail"] == "Review deleted successfully"

def test_get_reviews_by_event_id(test_client):
    logger.info("Starting test_get_reviews_by_event_id")
    response = test_client.get("/reviews/event/1")
    assert response.status_code == 200
    reviews = response.json()
    assert isinstance(reviews, list)
    assert len(reviews) == 3

    reviews.sort(key=lambda x: x["user_id"])

    assert reviews[0]["content"] == "Great event!"
    assert reviews[1]["content"] == "Not bad"
    assert reviews[2]["content"] == "Could be better"
    logger.info("Finished test_get_reviews_by_event_id")

def test_get_all_reviews(test_client):
    logger.info("Starting test_get_all_reviews")
    response = test_client.get("/reviews/")
    assert response.status_code == 200
    reviews = response.json()
    assert isinstance(reviews, list)
    assert len(reviews) == 9

    reviews.sort(key=lambda x: x["user_id"])

    assert reviews[0]["content"] == "Great event!"
    assert reviews[1]["content"] == "Not bad"
    assert reviews[2]["content"] == "Could be better"
    assert reviews[3]["content"] == "Loved it!"
    assert reviews[4]["content"] == "It was okay"
    assert reviews[5]["content"] == "Terrible experience"
    assert reviews[6]["content"] == "Pretty good"
    assert reviews[7]["content"] == "Not worth it"
    assert reviews[8]["content"] == "Fantastic!"
    logger.info("Finished test_get_all_reviews")