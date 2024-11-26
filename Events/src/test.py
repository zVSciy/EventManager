import unittest
from fastapi.testclient import TestClient
from main import app  # Importiere deine FastAPI-App

client = TestClient(app)

class TestEventAPI(unittest.TestCase):
    def test_000_create_event(self):
        response = client.get("/event/")
        self.assertEqual(response.status_code, 200)
        existing_events = response.json()

        for event in existing_events:
            if event["name"] == "Updated Concert Name" and event["location"] == "New Stadium" and event["ID"] == 1:
                print("Test-Event existiert bereits, Erstellung wird übersprungen.")
                return  
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
        
    def test_get_all_events(self):
        response = client.get("/event/")
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), list)

    def test_get_event(self):
        event_id = 1  
        response = client.get(f"/event/{event_id}")
        self.assertEqual(response.status_code, 200)
        self.assertIn("ID", response.json())

    def test_get_event_tickets(self):
        event_id = 1  
        response = client.get(f"/event/available_tickets/{event_id}")
        self.assertEqual(response.status_code, 200)
        self.assertIn("available_normal_tickets", response.json())
        self.assertIn("available_vip_tickets", response.json())


    def test_update_event(self):
        event_id = 1
        updated_event_data = {
            "name": "Updated Concert Name",
            "location": "New Stadium"
        }
        response = client.put(f"/event/{event_id}/", json=updated_event_data)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["response"], "Event was updated successfully")

    def test_cancel_event(self):
        event_id = 1
        cancel_data = {"canceled": True}
        response = client.put(f"/event/cancel/{event_id}", json=cancel_data)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["response"], "Event was canceled successfully")

    def test_update_event_tickets(self):
        event_id = 1
        ticket_update_data = {
            "available_normal_tickets": 80,
            "available_vip_tickets": 40
        }
        response = client.put(f"/event/updateTicket/{event_id}", json=ticket_update_data)
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()["response"], "Available tickets were updated successfully")

if __name__ == "__main__":
    unittest.main()
