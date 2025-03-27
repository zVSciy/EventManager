import unittest
from fastapi.testclient import TestClient
from main import app  # Import FastAPI application

client = TestClient(app)

class TestTicketCreationEditing(unittest.TestCase):
    def test_create_tickets(self):
        global ticket_id_created
        test_ticket = {
            "price": 200,
            "row": "A",
            "seat_number": 12,
            "vip": True,
            "user_id": 9999,
            "event_id": 1
        }
        response = client.post('/tickets/', json = test_ticket)
        self.assertEqual(response.status_code, 200)
        self.assertIsInstance(response.json(), dict)
        ticket_id_created = response.json().get('id')  # Store the ticket ID

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
        self.assertIsNotNone(ticket_id_created, "Ticket ID is not set. Run test_create_tickets first.")
        ticket_id = ticket_id_created
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
        self.assertIsNotNone(ticket_id_created, "Ticket ID is not set. Run test_create_tickets first.")
        ticket_id = ticket_id_created
        response = client.delete(f'/tickets/{ticket_id}/')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['msg'], 'Ticket cancelled successfully.')

if __name__ == "__main__":
    unittest.main()