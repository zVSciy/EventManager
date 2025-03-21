import React, { useState, useEffect, useRef } from 'react';
import './App.css';
let eventID;

// Use the proxy with /api prefix
const URL = '/app_review/api';

const API_URLS = {
  submitReview: `${URL}/reviews/create`,
  getReview: (id) => `${URL}/reviews/${id}`,
  getReviews: (eventId) => `${URL}/reviews/event/${eventId}`,
  updateReview: (id) => `${URL}/reviews/${id}`,
  deleteReview: (id) => `${URL}/reviews/${id}`,
  getAllReviews: `${URL}/reviews/`, 
};

function App() {
  const [selectedEndpoint, setSelectedEndpoint] = useState('submitReview');
  const [formData, setFormData] = useState({
    user_id: '',
    content: '',
    rating: '',
    review_id: '',
  });
  const [response, setResponse] = useState(null);

  const reviewFormRef = useRef(null);
  const getReviewFormRef = useRef(null);
  const getReviewsFormRef = useRef(null);
  const updateReviewFormRef = useRef(null);
  const deleteReviewFormRef = useRef(null);

  // Extract event ID from URL on component mount
  useEffect(() => {
    // Get event ID from URL parameters
    const urlParams = new URLSearchParams(window.location.search);
    const eventIdFromUrl = urlParams.get('eventId');
    
    if (eventIdFromUrl) {
      // Store as integer
      eventID = parseInt(eventIdFromUrl, 10);
      sessionStorage.setItem('eventId', eventID);
    } else {
      // Fallback to session storage
      const storedEventId = sessionStorage.getItem('eventId');
      if (storedEventId) {
        eventID = parseInt(storedEventId, 10);
      }
    }
    console.log("Event ID:", eventID);
  }, []);

  useEffect(() => {
    function updateButtonState(formRef, buttonId) {
      const form = formRef.current;
      const button = document.getElementById(buttonId);

      if (form && button) {
        if (selectedEndpoint === 'getAllReviews' || selectedEndpoint === 'getReviews' || form.checkValidity()) {
          button.classList.remove('btn-disabled');
          button.classList.add('solana-primary');
        } else {
          button.classList.remove('solana-primary');
          button.classList.add('btn-disabled');
        }
      }
    }
  
    const forms = [
      { ref: reviewFormRef, buttonId: 'submitReviewButton' },
      { ref: getReviewFormRef, buttonId: 'getReviewButton' },
      { ref: getReviewsFormRef, buttonId: 'getReviewsButton' },
      { ref: updateReviewFormRef, buttonId: 'updateReviewButton' },
      { ref: deleteReviewFormRef, buttonId: 'deleteReviewButton' },
    ];
  
    forms.forEach(({ ref, buttonId }) => {
      const updateState = () => updateButtonState(ref, buttonId);
      if (ref.current) {
        ref.current.addEventListener('input', updateState);
      }
      // Initial check on page load
      updateButtonState(ref, buttonId);
    });
  
    return () => {
      forms.forEach(({ ref, buttonId }) => {
        const updateState = () => updateButtonState(ref, buttonId);
        if (ref.current) {
          ref.current.removeEventListener('input', updateState);
        }
      });
    };
  }, [selectedEndpoint]);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    let url = '';
    let options = {};

    // Create a copy of form data and add event_id for submission
    const submissionData = { ...formData };
    
    // For endpoints that need event_id, use the one from URL
    if (selectedEndpoint === 'submitReview') {
      submissionData.event_id = eventID;
    }

    switch (selectedEndpoint) {
      case 'submitReview':
        url = API_URLS.submitReview;
        options = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(submissionData),
        };
        break;
      case 'getReview':
        url = API_URLS.getReview(formData.review_id);
        options = { method: 'GET' };
        break;
      case 'getReviews':
        url = API_URLS.getReviews(eventID);
        options = { method: 'GET' };
        break;
      case 'getAllReviews':
        url = API_URLS.getAllReviews;
        options = { method: 'GET' };
        break;
      case 'updateReview':
        url = API_URLS.updateReview(formData.review_id);
        options = {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(submissionData),
        };
        break;
      case 'deleteReview':
        url = API_URLS.deleteReview(formData.review_id);
        options = { method: 'DELETE' };
        break;
      default:
        break;
    }

    try {
      const response = await fetch(url, options);
      const data = await response.json();
      setResponse(data);
    } catch (error) {
      console.error('Error:', error);
      setResponse({ error: 'An error occurred' });
    }
  };

  return (
    <div className="App">
      <nav className="bg-gray-800 p-4">
        <div className="container mx-auto flex justify-between items-center">
          <button className="text-white text-xl font-bold bg-transparent border-0">Review App</button>
          <div>
            <a 
              href="/app_event/"
              className="text-gray-300 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >
              Back to Events
            </a>
          </div>
        </div>
      </nav>
      <div className="container mx-auto p-4">
        <div className="flex flex-wrap -mx-2">
          <div className="w-full px-2 mb-4 flex flex-col">
            <h1 className="text-3xl font-bold mb-4 text-solana-primary">Review Management</h1>
            {eventID && (
              <div className="mb-4 bg-gray-800 p-2 rounded">
                <p className="text-white">Current Event ID: <span className="font-bold">{eventID}</span></p>
              </div>
            )}
            <div className="mb-4">
              <label htmlFor="endpoint" className="block text-sm font-medium">Select Endpoint:</label>
              <select
                id="endpoint"
                name="endpoint"
                className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                value={selectedEndpoint}
                onChange={(e) => setSelectedEndpoint(e.target.value)}
              >
                <option value="submitReview">Submit Review</option>
                <option value="getReview">Get Review by ID</option>
                <option value="getReviews">Get Reviews by Event ID</option>
                <option value="getAllReviews">Get All Reviews</option>
                <option value="updateReview">Update Review</option>
                <option value="deleteReview">Delete Review</option>
              </select>
            </div>
            <form
              id="reviewForm"
              ref={reviewFormRef}
              className="bg-gray-800 p-4 rounded-lg shadow-md flex-grow"
              onSubmit={handleSubmit}
            >
              {selectedEndpoint === 'submitReview' && (
                <>
                  <div className="mb-4">
                    <label htmlFor="user_id" className="block text-sm font-medium">User ID:</label>
                    <input
                      type="number"
                      id="user_id"
                      name="user_id"
                      min="0"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.user_id}
                      onChange={handleInputChange}
                      required
                    />
                  </div>
                  <div className="mb-4">
                    <label htmlFor="content" className="block text-sm font-medium">Content:</label>
                    <textarea
                      id="content"
                      name="content"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.content}
                      onChange={handleInputChange}
                      required
                    ></textarea>
                  </div>
                  <div className="mb-4">
                    <label htmlFor="rating" className="block text-sm font-medium">Rating:</label>
                    <input
                      type="number"
                      id="rating"
                      name="rating"
                      min="1"
                      max="5"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.rating}
                      onChange={handleInputChange}
                      required
                    />
                  </div>
                </>
              )}
              {selectedEndpoint === 'getReview' && (
                <div className="mb-4">
                  <label htmlFor="review_id" className="block text-sm font-medium">Review ID:</label>
                  <input
                    type="number"
                    id="review_id"
                    name="review_id"
                    min="0"
                    className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                    value={formData.review_id}
                    onChange={handleInputChange}
                    required
                  />
                </div>
              )}
              {selectedEndpoint === 'getReviews' && (
                <div className="mb-4 text-center">
                  <p className="text-sm font-medium">Using Event ID: <span className="font-bold">{eventID}</span></p>
                  <p className="text-sm text-gray-400">(Automatically retrieved from URL)</p>
                </div>
              )}
              {selectedEndpoint === 'getAllReviews' && (
                <div className="mb-4 text-center">
                  <p className="text-sm font-medium">Click the button to get all reviews</p>
                </div>
              )}
              {selectedEndpoint === 'updateReview' && (
                <>
                  <div className="mb-4">
                    <label htmlFor="review_id" className="block text-sm font-medium">Review ID:</label>
                    <input
                      type="number"
                      id="review_id"
                      name="review_id"
                      min="0"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.review_id}
                      onChange={handleInputChange}
                      required
                    />
                  </div>
                  <div className="mb-4">
                    <label htmlFor="content" className="block text-sm font-medium">Content:</label>
                    <textarea
                      id="content"
                      name="content"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.content}
                      onChange={handleInputChange}
                      required
                    ></textarea>
                  </div>
                  <div className="mb-4">
                    <label htmlFor="rating" className="block text-sm font-medium">Rating:</label>
                    <input
                      type="number"
                      id="rating"
                      name="rating"
                      min="1"
                      max="5"
                      className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                      value={formData.rating}
                      onChange={handleInputChange}
                      required
                    />
                  </div>
                </>
              )}
              {selectedEndpoint === 'deleteReview' && (
                <div className="mb-4">
                  <label htmlFor="review_id" className="block text-sm font-medium">Review ID:</label>
                  <input
                    type="number"
                    id="review_id"
                    name="review_id"
                    min="0"
                    className="mt-1 block w-full bg-gray-700 border border-gray-600 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-solana-primary"
                    value={formData.review_id}
                    onChange={handleInputChange}
                    required
                  />
                </div>
              )}
            <button type="submit" id="submitReviewButton" className="w-full btn-disabled text-white font-bold py-2 px-4 rounded-md">
              {selectedEndpoint === 'submitReview' && 'Submit Review'}
              {selectedEndpoint === 'getReview' && 'Get Review'}
              {selectedEndpoint === 'getReviews' && 'Get Reviews'}
              {selectedEndpoint === 'getAllReviews' && 'Get All Reviews'}
              {selectedEndpoint === 'updateReview' && 'Update Review'}
              {selectedEndpoint === 'deleteReview' && 'Delete Review'}
            </button>
            </form>
            {response && (
              <div className="mt-4 bg-gray-800 p-4 rounded-lg shadow-md">
                <h2 className="text-xl font-bold mb-2 text-solana-primary">Response</h2>
                <pre className="bg-gray-900 p-2 rounded-md text-left text-white">
                  {JSON.stringify(response, null, 2)}
                </pre>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;