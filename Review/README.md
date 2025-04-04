# Review Service

The Review Service is a component of the EventManager application that allows users to submit, retrieve, update, and delete reviews for events.

## Architecture

The application consists of two main components:
- **Backend**: A REST API built with FastAPI that handles review data operations
- **Frontend**: A React application that provides a user interface for interacting with the review service

## Backend API

The backend service exposes the following endpoints:

### GET /reviews/{review_id}
Retrieves a specific review by its ID.

**Response:**
```json
{
    "content": "Great event!",
    "event_id": 1,
    "rating": 5,
    "user_id": 1,
    "id": 1
}
```

### GET /reviews/
Retrieves all reviews in the system.

**Response:**
```json
[
    {
        "event_id": 1,
        "content": "Great event!",
        "rating": 5,
        "user_id": 1,
        "id": 1
    },
    {
        "event_id": 2,
        "content": "Not bad",
        "rating": 3,
        "user_id": 2,
        "id": 2
    },
    {
        "event_id": 1,
        "content": "Could be better",
        "rating": 2,
        "user_id": 3,
        "id": 3
    },
    {
        "event_id": 2,
        "content": "Loved it!",
        "rating": 5,
        "user_id": 4,
        "id": 4
    }
]
```

### GET /reviews/event/{event_id}
Retrieves all reviews for a specific event.

**Response:**
```json
[
    {
        "event_id": 3,
        "content": "Terrible experience",
        "rating": 1,
        "user_id": 6,
        "id": 6
    },
    {
        "event_id": 3,
        "content": "Pretty good",
        "rating": 4,
        "user_id": 7,
        "id": 7
    }
]
```

### POST /reviews
Creates a new review.

**Request:**
```json
{
    "user_id": 4,
    "content": "Amazing event!",
    "rating": 5,
    "event_id": 1
}
```

**Response:**
```json
{
    "event_id": 1,
    "content": "Amazing event!",
    "rating": 5,
    "user_id": 4,
    "id": 10
}
```

### PUT /reviews/{review_id}
Updates an existing review.

**Request:**
```json
{
    "user_id": 4,
    "content": "Updated opinion: Not as good as I initially thought",
    "rating": 3,
    "event_id": 1
}
```

**Response:**
```json
{
    "event_id": 1,
    "content": "Updated opinion: Not as good as I initially thought",
    "rating": 3,
    "user_id": 4,
    "id": 10
}
```

### DELETE /reviews/{review_id}
Deletes a review.

**Response:**
```json
{
    "detail": "Review deleted successfully"
}
```

## Database Schema

The review service uses a database with the following schema:

| Field   | Type | Description               |
|---------|------|---------------------------|
| id      | INT  | Primary key               |
| user_id | INT  | User ID who left review   |
| content | TEXT | Review content            |
| rating  | INT  | Rating (typically 1-5)    |
| event_id| INT  | Event ID being reviewed   |

> Note: The review service requires data from Ticket, Event, and User services.

## Frontend Implementation

The frontend is built with React and provides a user interface for interacting with the review service API. It includes:

1. A form interface for submitting, retrieving, updating, and deleting reviews
2. Dropdown selection for choosing the API endpoint to interact with
3. Dynamic form fields that change based on the selected endpoint
4. Response display for showing API results

### Frontend API Integration

The frontend application communicates with the backend through a set of predefined API URLs:

```javascript
const API_URLS = {
  submitReview: `${API_BASE_URL}/reviews`,
  getReview: (id) => `${API_BASE_URL}/reviews/${id}`,
  getReviews: (eventId) => `${API_BASE_URL}/reviews/event/${eventId}`,
  getAllReviews: `${API_BASE_URL}/reviews/`,
  updateReview: (id) => `${API_BASE_URL}/reviews/${id}`,
  deleteReview: (id) => `${API_BASE_URL}/reviews/${id}`
};
```

### Authentication

The frontend includes support for authentication when making API requests. While the specific authentication mechanism is implementation-dependent, the application handles authentication through:

1. Allowing users to select their identity
2. Passing user information in API requests
3. Potentially handling authorization tokens for secured endpoints

## Getting Started

### Backend Setup
1. Install required dependencies
2. Configure your database connection
3. Start the FastAPI server

### Frontend Setup
1. Navigate to the frontend directory
2. Install dependencies with `npm install`
3. Start the development server with `npm start`

## Docker Setup

The Review Service can be easily deployed using Docker and Docker Compose.

### Prerequisites

- Docker installed on your machine
- Docker Compose installed on your machine

### Docker Configuration

The application is containerized with the following services:

- **review_api**: The backend FastAPI service running on port 8083
- **review_web**: The frontend React application running on port 3000

The services are defined in the `docker-compose.yml` file:

```yml
version: '3.8'

networks:
  my-network:
    driver: bridge

services:
  review_api:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8083:8083"
    command: >
      sh -c "exec uvicorn main:app --host 0.0.0.0 --port 8083 --reload"
    networks:
      - my-network

  review_web:
    build:
      context: ./frontend
      dockerfile: Dockerfile.dev
    ports:
      - "3000:3000"
    volumes:
      - ./frontend:/app
      - /app/node_modules
    environment:
      - CHOKIDAR_USEPOLLING=true
      - WDS_SOCKET_PORT=0
      - NODE_ENV=development
    networks:
      - my-network
    depends_on:
      - review_api
```

### Running with Docker

1. **Build and start the containers:**

   ```bash
   docker-compose up --build
   ```

   This command builds the images (if they don't exist) and starts the containers.
   
   To run the containers in detached mode (background):
   
   ```bash
   docker-compose up -d --build
   ```

2. **Access the services:**

   - Backend API: http://localhost:8083
   - Frontend UI: http://localhost:3000

3. **Stop the containers:**

   ```bash
   docker-compose down
   ```

4. **View logs:**

   ```bash
   docker-compose logs
   ```
   
   To follow the logs in real-time:
   
   ```bash
   docker-compose logs -f
   ```

### Development with Docker

The setup includes hot-reloading for both frontend and backend:

- Backend changes will automatically reload the FastAPI application
- Frontend changes will be reflected immediately in the browser 

The frontend container uses volumes to mount the local `frontend` directory into the container, allowing for real-time code updates without rebuilding the container.

## Testing

The application includes test configurations for both frontend and backend components. Frontend tests use Jest with React Testing Library.
