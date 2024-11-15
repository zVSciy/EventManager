from fastapi.testclient import TestClient
from main import app  # Import your FastAPI app

client = TestClient(app)

def test_read_all_notifications():
    response = client.get("/notifications")
    assert response.status_code == 200
    assert type(response.json()) == list
    if(len(response.json()) > 0):
        assert type(response.json()[0]) == dict

