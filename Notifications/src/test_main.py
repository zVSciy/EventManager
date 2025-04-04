from fastapi.testclient import TestClient
from main import app  # Import your FastAPI app

client = TestClient(app)

def test_read_all_notifications():
    response = client.get("/notifications")
    if response.status_code == 404:
        return
    assert response.status_code == 200
    assert type(response.json()) == list
    if(len(response.json()) > 0):
        assert type(response.json()[0]) == dict


def test_read_notification():
    response = client.get("/notifications/1")
    print(response)
    if response.status_code == 404:
        return
    assert response.status_code == 200
    assert type(response.json()) == dict

def test_create_notification():
    new_notification = {
    "description": "TESTNOTIFICATION",
    "status": "active",
    "timestamp": 1729235305,
    "eventId": 72,
    "paymentId": 123,
    "ticketId": 123123,
    "userId": 'test1@gmail.com'
    }
    response = client.post("/notifications/", json=new_notification)
    assert response.status_code == 200
    assert response.json()["description"] == new_notification["description"]
    assert response.json()["status"] == new_notification["status"]

def test_read_all_notifications_by_user():
    response = client.get("/notifications/user/test1@gmail.com")
    assert response.status_code == 200
    assert type(response.json()) == list
    if(len(response.json()) > 0):
        assert type(response.json()[0]) == dict



def test_update_notification():
    event_id = client.get("/notifications/user/test1@gmail.com").json()[0]["id"]
    event_id = int(event_id)
    update_data = {
    "description": "string",
    "status": "string",
    "timestamp": 5,
    "eventId": 5,
    "paymentId": 5,
    "ticketId": 5,
    "userId": 'test2@gmail.com'
    }
    response = client.put(f"/notifications/{event_id}", json=update_data)
    # print(response)
    # print(response.json())
    assert response.status_code == 200
    assert response.json()["description"] == update_data["description"]
    assert response.json()["status"] == update_data["status"]
