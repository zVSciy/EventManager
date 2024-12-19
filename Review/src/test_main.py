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

# Create an in-memory test database
SQLALCHEMY_DATABASE_URL = "sqlite:///:memory:"
engine = create_engine(SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False})
TestingSessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

# Override the get_db dependency to use the test database
def override_get_db():
    try:
        db = TestingSessionLocal()
        yield db
    finally:
        db.close()

app.dependency_overrides[get_db] = override_get_db

# Create the test client
client = TestClient(app)

@pytest.fixture(scope="session", autouse=True)
def setup_database():
    # Create the test database tables once for the entire session
    Base.metadata.create_all(bind=engine)
    yield
    # Do not drop the test database tables after the session

@pytest.fixture(scope="module")
def test_client():
    yield client

def test_create_review(test_client):
    logger.info("Starting test_create_review")
    response = test_client.post("/reviews/", json={
        "user_id": 1,
        "content": "Great event!",
        "rating": 5,
        "event_id": 1
    })
    assert response.status_code == 200
    assert response.json()["content"] == "Great event!"

def test_get_review(test_client):
    logger.info("Starting test_get_review")
    response = test_client.post("/reviews/", json={
        "user_id": 1,
        "content": "Great event!",
        "rating": 5,
        "event_id": 1
    })
    assert response.status_code == 200
    review_id = response.json()["id"]

    response = test_client.get(f"/reviews/{review_id}")
    assert response.status_code == 200
    assert response.json()["content"] == "Great event!"

def test_update_review(test_client):
    logger.info("Starting test_update_review")
    response = test_client.post("/reviews/", json={
        "user_id": 1,
        "content": "Great event!",
        "rating": 5,
        "event_id": 1
    })
    assert response.status_code == 200
    review_id = response.json()["id"]

    response = test_client.put(f"/reviews/{review_id}", json={
        "content": "Updated content",
        "rating": 4
    })
    assert response.status_code == 200
    assert response.json()["content"] == "Updated content"

def test_delete_review(test_client):
    logger.info("Starting test_delete_review")
    response = test_client.post("/reviews/", json={
        "user_id": 1,
        "content": "Great event!",
        "rating": 5,
        "event_id": 1
    })
    assert response.status_code == 200
    review_id = response.json()["id"]

    response = test_client.delete(f"/reviews/{review_id}")
    assert response.status_code == 200
    assert response.json()["detail"] == "Review deleted successfully"

def test_get_reviews_by_event_id(test_client):
    logger.info("Starting test_get_reviews_by_event_id")
    reviews = [
        {"user_id": 1, "content": "Great event!", "rating": 5, "event_id": 1},
        {"user_id": 2, "content": "Not bad", "rating": 3, "event_id": 1},
        {"user_id": 3, "content": "Could be better", "rating": 2, "event_id": 1}
    ]

    for review in reviews:
        test_client.post("/reviews/", json=review)

    response = test_client.get("/reviews/event/1")
    assert response.status_code == 200
    reviews = response.json()
    assert isinstance(reviews, list)
    assert len(reviews) > 0

    reviews.sort(key=lambda x: x["user_id"])

    assert reviews[0]["content"] == "Great event!"
    assert reviews[1]["content"] == "Not bad"
    assert reviews[2]["content"] == "Could be better"
    logger.info("Finished test_get_reviews_by_event_id")