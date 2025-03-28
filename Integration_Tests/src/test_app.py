import unittest
import requests

BASE_URL = "http://auth_api:8000"
EVENTS_URL = "http://events_api:8000"
TICKETS_URL = "http://tickets_api:8000"

class TestIntegrationWorkflow(unittest.TestCase):
    def test_user_registration(self):
        # Step 1: Register a user
        user_payload = {
            "email": "testuser@example.com",
            "password": "strongpassword",
            "first_name": "Test",
            "last_name": "User",
            "role": "user"
        }
        register_response = requests.post(f"{BASE_URL}/register", json=user_payload)
        self.assertEqual(register_response.status_code, 200)
        self.assertIsInstance(response.json(), dict)

        # Step 2: Get Event details for Event with id = 1
        event_response = requests.get(f"{EVENTS_URL}/event/1")
        self.assertEqual(event_response.status_code, 200)
        event_data = event_response.json()
        self.assertIn("ID", event_data)
        print("Event details retrieved successfully.")

        # Step 3: Add a ticket for the Event with id = 1
        ticket_payload = {
            "price": 100,
            "row": "A",
            "seat_number": 1,
            "vip": False,
            "user_id": user_payload["email"],
            "event_id": 1
        }
        ticket_response = requests.post(f"{TICKETS_URL}/tickets", json=ticket_payload)
        self.assertEqual(ticket_response.status_code, 200)
        print("Ticket added successfully.")

if __name__ == "__main__":
    unittest.main()