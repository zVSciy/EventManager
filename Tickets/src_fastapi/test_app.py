import unittest
from fastapi.testclient import TestClient
from main import app  # Import FastAPI application

client = TestClient(app)

class TestTicketCreationEditing(unittest.TestCase):
    def test_create_tickets(self):
        test_ticket = {
            "price": 200,
            "row": "A",
            "seat_number": 12,
            "vip": "true",
            "user_id": 9999,
            "event_id": 1
        }
        response = client.post('/tickets/', json = test_ticket)
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), dict)

    def test_edit_tickets_not_found(self):
        ticket_id = 100000
        update_ticket = {
            "price": 9999,
            "row": "B",
            "seat_number": 10,
            "vip": True,
            "user_id": 9999,
            "event_id": 1
        }
        response = client.put(f'/tickets/{ticket_id}/', json = update_ticket)
        self.assertEqual(response.status_code, 404)
        self.assertEqual(response.json()['detail']['msg'], f"Ticket with `id`: `{ticket_id}` doesn't exist.")

    def test_edit_tickets(self):
        ticket_id = 1
        update_ticket = {
            "price": 9999,
            "row": "B",
            "seat_number": 10,
            "vip": True,
            "user_id": 9999,
            "event_id": 1
        }
        response = client.put(f'/tickets/{ticket_id}/', json = update_ticket)
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), dict)

class TestTicketShow(unittest.TestCase):
    def test_get_all_tickets(self):
        response = client.get('/tickets/')
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), list)

    def test_get_tickets_for_event(self):
        event_id = 1
        response = client.get(f'/tickets?event_id={event_id}')
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), list)

class TestTicketDeletion(unittest.TestCase):
    def test_cancel_tickets_not_found(self):
        ticket_id = 100000
        response = client.delete(f'/tickets/{ticket_id}/')
        self.assertEqual(response.status_code, 404)
        self.assertEqual(response.json()['detail']['msg'], f"Ticket with `id`: `{ticket_id}` doesn't exist.")
    
    def test_cancel_tickets(self):
        ticket_id = 1
        response = client.delete(f'/tickets/{ticket_id}/')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['msg'], 'Ticket cancelled successfully.')

if __name__ == "__main__":
    unittest.main()