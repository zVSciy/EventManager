import unittest
import requests

BASE_URL = "https://auth_api:8000"
EVENTS_URL = "http://events_api:8000"
TICKETS_URL = "http://tickets_api:8000"
REVIEWS_URL = "http://review_api:8083"

user_payload = {
    "email": "testuser@example.com",
    "password": "strongpassword",
    "first_name": "Test",
    "last_name": "User",
    "role": "user"
}

review_payload = {
    "user_id": user_payload["email"],
    "content": "Amazing event!",
    "rating": 5,
    "event_id": 1
}

# Initialize global variables to store created ticket and review IDs
ticket_id_created = None
review_id_created = None

class TestIntegrationUserWorkflow(unittest.TestCase):        
    def test_1_register_user(self):
        # Step 1: Register a user
        register_response = requests.post(f"{BASE_URL}/register", json=user_payload, verify=False) # Verify=False to ignore SSL warnings
        if register_response.status_code == 400:
            print("NOTE: User already exists, skipping registration.")
            
        else:
            self.assertEqual(register_response.status_code, 200)
            self.assertEqual(register_response.json(), {"message": "User registered successfully"})
            print("User registered successfully.")

    def test_2_get_event_details(self):
        # Step 2: Get Event details for Event with id = 1
        event_response = requests.get(f"{EVENTS_URL}/event/1")
        self.assertEqual(event_response.status_code, 200)
        event_data = event_response.json()
        self.assertIn("ID", event_data)
        print("Event details retrieved successfully.")

    def test_3_add_ticket(self):    
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
        global ticket_id_created
        ticket_id_created = ticket_response.json().get('id')  # Store the ticket ID
        print("Ticket created successfully.")

    def test_4_get_tickets_of_user_for_event(self):
        # Step 4: Get all tickets for the user for the event with id = 1
        self.assertIsNotNone(ticket_id_created, "Ticket ID is not set. Run test_add_ticket first.")
        ticket_response = requests.get(f"{TICKETS_URL}/tickets/user/{user_payload['email']}?event_id=1")
        self.assertEqual(ticket_response.status_code, 200)
        tickets = ticket_response.json()
        self.assertIsInstance(tickets, list)
        print("Ticket details retrieved successfully.")
    
    def test_5_cancel_tickets(self):
        # Step 5: Cancel the ticket created in Step 3
        self.assertIsNotNone(ticket_id_created, "Ticket ID is not set. Run test_add_ticket first.")
        ticket_id = ticket_id_created
        response = requests.delete(f"{TICKETS_URL}/tickets/user/{user_payload['email']}/ticket/{ticket_id}/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['msg'], 'Ticket cancelled successfully.')
        print("Ticket cancelled successfully.")

    def test_6_write_review(self):
        # Step 6: Write a review for the event with id = 1
        review_response = requests.post(f"{REVIEWS_URL}/reviews/create", json=review_payload)
        self.assertEqual(review_response.status_code, 200)
        review_data = review_response.json()
        global review_id_created
        review_id_created = review_data.get('id')  # Store the review ID
        print("Review created successfully.")

    def test_7_get_review(self):
        # Step 7: Get the review for the event with id = 1
        self.assertIsNotNone(review_id_created, "Review ID is not set. Run test_write_review first.")
        review_id = review_id_created
        review_response = requests.get(f"{REVIEWS_URL}/reviews/{review_id}")
        self.assertEqual(review_response.status_code, 200)
        review_data = review_response.json()
        self.assertEqual(review_data["content"], "Amazing event!")
        self.assertEqual(review_data["rating"], 5)
        print("Review details retrieved successfully.")

    def test_8_delete_review(self):
        # Step (): Delete the review for the event with id = 1
        self.assertIsNotNone(review_id_created, "Review ID is not set. Run test_write_review first.")
        review_id = review_id_created
        review_response = requests.delete(f"{REVIEWS_URL}/reviews/{review_id}")
        self.assertEqual(review_response.status_code, 200)
        review_data = review_response.json()
        self.assertEqual(review_data["detail"], "Review deleted successfully")
        print("Review deleted successfully.")

if __name__ == "__main__":
    unittest.main()