# /tests/test_api.py
import unittest
from fastapi.testclient import TestClient
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from model.database import get_db
from model.models import Base
from main import app

# Mocking the database with SQLite
SQLALCHEMY_DATABASE_URL = "sqlite:///./test.db"
engine = create_engine(SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False})
TestingSessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base.metadata.create_all(bind=engine)

# Dependency override
def override_get_db():
    db = TestingSessionLocal()
    try:
        yield db
    finally:
        db.close()

app.dependency_overrides[get_db] = override_get_db
client = TestClient(app)

class TestAPI(unittest.TestCase):

    def setUp(self):
        # Setting up test database
        self.db = TestingSessionLocal()
        Base.metadata.create_all(bind=engine)

    def tearDown(self):
        # Clearing test database after each test
        self.db.close()
        Base.metadata.drop_all(bind=engine)
        print("✅ All tests in this TestCase ran successfully!")

    def test_register_user_success(self):
        payload = {
            "email": "testuser@example.com",
            "password": "strongpassword",
            "first_name": "Test",
            "last_name": "User",
            "role": "user"
        }
        response = client.post("/register", json=payload)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json(), {"message": "User registered successfully"})

    def test_register_user_existing_email(self):
        # Register user once
        payload = {
            "email": "testuser@example.com",
            "password": "strongpassword",
            "first_name": "Test",
            "last_name": "User",
            "role": "user"
        }
        client.post("/register", json=payload)
        # Try registering again
        response = client.post("/register", json=payload)
        self.assertEqual(response.status_code, 400)
        self.assertIn("Email already registered", response.json()["detail"])

    def test_login_success(self):
        # Register user first
        payload = {
            "email": "testuser@example.com",
            "password": "strongpassword",
            "first_name": "Test",
            "last_name": "User",
            "role": "user"
        }
        client.post("/register", json=payload)
        # Login with correct credentials
        login_payload = {
            "username": "testuser@example.com",
            "password": "strongpassword"
        }
        response = client.post("/token", data=login_payload)
        self.assertEqual(response.status_code, 200)
        self.assertIn("access_token", response.json())

    def test_login_invalid_credentials(self):
        login_payload = {
            "username": "wronguser@example.com",
            "password": "wrongpassword"
        }
        response = client.post("/token", data=login_payload)
        self.assertEqual(response.status_code, 401)
        self.assertIn("Invalid credentials", response.json()["detail"])

if __name__ == "__main__":
    unittest.main()
