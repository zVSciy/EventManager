# Review Service

## Overview
The Review Service allows you to create, retrieve, update, and delete reviews for various events.

## Prerequisites
- Docker
- Docker Compose
- Postman (for manual testing)

## Project Structure


## Running the Service

**Build and start the Docker container:**

   Navigate to the `Review` directory and run the following command:

   ```sh
   docker-compose up --build
   ```

   This will build the Docker container and start the Review Service. The service will be available at `http://localhost:8083`.

## Running Tests

### Automated tests with pytest:

The tests will automatically run when the Docker container is started. To run the tests manually, you can execute the following command inside the container:

   ```sh
   docker-compose run api pytest
   ```

### Manual Testing with Postman

Set up Postman:

Open Postman and create a new collection for the Review Service.

You can test every endpoint listed below:

## Feedback | Showcase 

### GET /reviews/{review_id}
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
**Request:**
```json
{
    "user_id": 4,
    "content": "Shit event!!!",
    "rating": 1,
    "event_id": 1
}
```

**Response:**
```json
{
    "event_id": 1,
    "content": "Shit event!!!",
    "rating": 1,
    "user_id": 4,
    "id": 10
}
```

### DELETE /reviews/{review_id} 
**Response:**
```json
{
    "detail": "Review deleted successfully"
}
```

### Database Data
- ID INT PRIMARY KEY 
- User INT
- Comment Text
- Rating INT
- Event INT  

> Needs Data from Ticket-, Event- and Userservice.

## Troubleshooting

- Ensure that Docker and Docker Compose are correctly installed.
- Check the Docker logs for errors: `docker-compose logs`