import unittest
from fastapi.testclient import TestClient
from main import app  # Importiere deine FastAPI-App

client = TestClient(app)

class TestEventAPI(unittest.TestCase):

    def test_get_all_events(self):
        response = client.get("/event/")
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), list)

    def test_get_event(self):
        event_id = 1  # Beispiel-ID
        response = client.get(f"/event/{event_id}")
        if response.status_code == 404:
            self.assertEqual(response.json(), {"detail": "Event not found"})
        else:
            self.assertEqual(response.status_code, 200)
            self.assertIn("ID", response.json())

    def test_get_event_tickets(self):
        event_id = 1  # Beispiel-ID
        response = client.get(f"/event/available_tickets/{event_id}")
        if response.status_code == 404:
            self.assertEqual(response.json(), {"detail": "Event not found"})
        else:
            self.assertEqual(response.status_code, 200)
            self.assertIn("available_normal_tickets", response.json())
            self.assertIn("available_vip_tickets", response.json())

    def test_create_event(self):
        event_data = {
            "name": "Concert",
            "location": "Stadium",
            "organisator": "Music Inc",
            "startdate": "2024-12-31T20:00:00",
            "available_normal_tickets": 100,
            "available_vip_tickets": 50
        }
        response = client.post("/event/", json=event_data)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["code"], 200)
        self.assertEqual(response.json()["response"], "Event was created successfully")

    def test_update_event(self):
        event_id = 1  # Beispiel-ID
        updated_event_data = {
            "name": "Updated Concert Name",
            "location": "New Stadium"
        }
        response = client.put(f"/event/{event_id}/", json=updated_event_data)
        if response.status_code == 404:
            self.assertEqual(response.json(), {"detail": "Event not found"})
        else:
            self.assertEqual(response.status_code, 200)
            self.assertEqual(response.json()["response"], "Event was updated successfully")

    def test_cancel_event(self):
        event_id = 1  # Beispiel-ID
        cancel_data = {"canceled": True}
        response = client.put(f"/event/cancel/{event_id}", json=cancel_data)
        if response.status_code == 404:
            self.assertEqual(response.json(), {"detail": "Event not found"})
        else:
            self.assertEqual(response.status_code, 200)
            self.assertEqual(response.json()["response"], "Event was canceled successfully")

    def test_update_event_tickets(self):
        event_id = 1  # Beispiel-ID
        ticket_update_data = {
            "available_normal_tickets": 80,
            "available_vip_tickets": 40
        }
        response = client.put(f"/event/updateTicket/{event_id}", json=ticket_update_data)
        if response.status_code == 404:
            self.assertEqual(response.json(), {"detail": "Event not found"})
        else:
            self.assertEqual(response.status_code, 200)
            self.assertEqual(response.json()["response"], "Available tickets were updated successfully")

if __name__ == "__main__":
    unittest.main()