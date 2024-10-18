# EventManager
Am boutta readme 

# Concepts
## Process
![image](https://github.com/user-attachments/assets/8a319b5d-f45f-47e2-beca-79070e116fe3)


## Microservices Idea
![image](https://github.com/user-attachments/assets/7d39198e-47e7-498f-9850-06615f22271d)

# API Endpoints
## Event
**GET** /event/  
**GET** /event/\<event-id>

**Response**
```json
{
  "code": 200,
  "reponse": [
      {
        "ID": 1,
        "name": "Markforster Konzert",
        "location": "Ruine Landskron",
        "organisator": "Gemeinde Finkenstein",
        "startdate": "2024-09-29T20:00:00Z",
        "available_normal_tickets": 100,
        "available_vip_tickets": 20,
        "canceled": false
      },
      {
        "ID": 2,
        "name": "Ramstein Konzert",
        "location": "Klagenfurter Stadion",
        "organisator": "Klagenfurter Konzertverein",
        "startdate": "2024-08-29T22:10:00Z",
        "available_normal_tickets": 2000,
        "available_vip_tickets": 200,
        "canceled": false
      }
    ],
}
```
**GET** /event/available_tickets/\<event-id>
```json
{
  "code": 200,
  "reponse": [
      {
        "ID": 1,
        "available_normal_tickets": 100,
        "available_vip_tickets": 20,
      }

    ],
}
```
**POST** /event/  
**Payload**
```json
{
  "name": "Markforster Konzert",
  "location": "Ruine Landskron",
  "organisator": "Gemeinde Finkenstein",
  "startdate": "2024-09-29T20:00:00Z",
  "available_normal_tickets": 100,
  "available_vip_tickets": 20,
}
```

**Response** 
```json
{
  "code": 200,
  "reponse": "Event was created successfully"
  "eventID": 1
}
```

**PUT** /event/\<event-id>/  
**Payload**
```json
{
  "name": "Bla Konzert",
  "location": "Ruine Finkenstein",
  "organisator": "Gemeinde Finkenstein",
  "startdate": "2024-09-29T20:00:00Z",
  "available_normal_tickets": 10,
  "available_vip_tickets": 50,

}
```

**Response**
```json
{
  "code": 200,
  "reponse": "Event was updated successfully"
  "eventID": 1
}
```

**PUT** /event/cancel/\<event-id>  
**Payload**
```json
{
  "canceled": true

}
```

**Response**
```json
{
  "code": 200,
  "reponse": "Event was canceled successfully"
  "eventID": 1
}
```

**PUT** /event/updateTicket/<event-id>  
**Payload**
```json
{
  "available_normal_tickets": 20
  "abailable_vip_tickets": 5

}
```

**Response**
```json
{
  "code": 200,
  "reponse": "Available tickets were updated successfully"
  "eventID": 1
}
```



### Database Data
ID  INT  AUTO_INCREMENT  PRIMARY KEY  
Name  VARCHAR  
Location VARCHAR  
Organisator VARCHAR  
StartDate  DateTime  
Available_normal_tickets  INT  
Available_vip_tickets  INT  
Canceled  BOOLEAN  Default False  

> Needs to know if the user is admin or not.

## Ticket
**GET** /ticket/  
**GET** /ticket/\<ticket-id\>

**Reponse:**
```json
{
  "code": 200,
  "reponse": [
      {
        "row": "A",
        "price": 12.99,
        "seatNumber": 3,
        "state": "paid",
        "vip": true,
        "user": 1,
        "event": 1 
      },
      {
        "row": "A",
        "price": 12.99,
        "seatNumber": 4,
        "state": "paid",
        "vip": true,
        "user": 1,
        "event": 1 
      }
    ],
}
```

**POST** /ticket

**Payload:**
```json
{
  "row": "A",
  "price": 12.99,
  "seatNumber": 3,
  "state": "awaiting payment",
  "vip": true,
  "user": 1,
  "event": 1 
}
```
**Reponse:**
```json
{
  "code": 201,
  "reponse": "Successfully created object!",
  "createdId": 1
}
```

**PUT** /ticket/\<ticket-id\>  

**Payload:**
```json
{
  "row": "A",
  "price": 12.99,
  "seatNumber": 3,
  "state": "paid",
  "vip": true,
  "user": 1,
  "event": 1 
}
```
**Reponse:**
```json
{
  "code": 200,
  "reponse": "Successfully updated object!",
  "affectedId": 1
}
```

### Database Data
ID  INT  PRIMARY KEY  
Row  VARCHAR  NULL  
Price  INT  
SeatNumber  INT  NULL  
State  VARCHAR  
VIP  BOOLEAN  
User  INT  (References to User.id)  
Event INT  (References to Event.id)  

> Needs Data from Event- and Userservice

## Payment

#### Endpoints Exclusive to Own Interface:

### POST /login
```http
POST /login
Content-Type: application/json

{
  "username": "user123",
  "password": "password123"
}
```
**Response:**
```json
{
  "token": "session-token-123"
}
```

### POST /logout
```http
POST /logout
Content-Type: application/json
Authorization: Bearer session-token-123
```
**Response:**
```json
{
  "message": "Logged out successfully"
}
```

#### Additional Endpoints:

### GET /accounts/{userId}
```http
GET /accounts/user123
```
**Response:**
```json
{
  "userId": "user123",
  "accountStatus": "exists"
}
```

### POST /accounts
```http
POST /accounts
Content-Type: application/json

{
  "userId": "user123"
}
```
**Response:**
```json
{
  "userId": "user123",
  "accountStatus": "created"
}
```

### POST /payments
```http
POST /payments
Content-Type: application/json

{
  "userId": "user123",
  "amount": 100,
  "currency": "EUR"
}
```
**Response:**
```json
{
  "paymentId": "payment123",
  "status": "initiated"
}
```

### POST /payments/{paymentId}/process
```http
POST /payments/payment123/process
Content-Type: application/json
```
**Response:**
```json
{
  "paymentId": "payment123",
  "status": "processed"
}
```

### Database Data
#### Collections
- Accounts
  - ``id``, ``userId``, ``accountStatus``, ``balance``, ``createdAt``

- Payments
  - ``id``, ``userId``, ``amount``, ``currency``, ``status``, ``createdAt``, ``processedAt``


> Needs Data from Ticket- and Userservice

## Notification

### GET /notification
**Response:**
```json
[
  {
      "id": 3,
      "timestamp": 1729235305,
      "paymentId": 2,
      "eventId": 1,
      "description": "testdesc",
      "status": "active",
      "ticketId": 4
  },
  {
      "id": 3,
      "timestamp": 1729235305,
      "paymentId": 2,
      "eventId": 1,
      "description": "testdesc",
      "status": "active",
      "ticketId": 4
  }
],
```

### GET /notification/\<id\>
**Response:**
```json
{
    "id": 3,
    "timestamp": 1729235305,
    "paymentId": 2,
    "eventId": 1,
    "description": "testdesc",
    "status": "active",
    "ticketId": 4
}
```


### POST /notification  
**Request:**
```json
{
  "userId": 1,
  "description": "Very very very very very very long string",
  "status": "active",
  "timestamp": 1728035912
}
```
> Note the timestamp is in seconds, the example timestamp is somewhere around 12:00 04.10.2024  
**Response:**
```json
{
    "id": 3,
    "timestamp": 1729235305,
    "paymentId": 2,
    "eventId": 1,
    "description": "testdesc",
    "status": "active",
    "ticketId": 4
}
```

### PUT /notification  
**Request:**
```json
{
  "id": 68,
  "description": "different text",
  "status": "hidden",
  "timestamp": 1728035912,
  "eventId": 123,
  "paymentId": 123,
  "ticketId": 1234
}
```
> Note the timestamp is in seconds, the example timestamp is somewhere around 12:00 04.10.2024  
**Response:**
```json
{
    "id": 3,
    "timestamp": 1729235305,
    "paymentId": 2,
    "eventId": 1,
    "description": "different text",
    "status": "active",
    "ticketId": 4
}
```



### Database Data
ID INT  PRIMARY KEY  AUTO INCREMENT  
Status   VARCHAR  
Description  Text  
Timestamp  INT  
UserId  INT  (References to User.Id)  
EventId INT (References to Event.Id)  
PaymentId INT (References to Payment.Id)  
TicketId INT (References to Ticket.Id)  

> Needs Data from Ticket-, Event- and Paymentservice  

## Feedback
/create  
/read  
/update  
/delete  

### Database Data
ID INT  PRIMARY KEY  
User  INT  
Comment  Text  
Event INT  
Date  Datetime  

> Needs Data from Ticket-, Event- and Userservice.

## User
/create  
/update  
/read  
(/delete)  

### Database Data
ID  INT  PRIMARY KEY  
Firstname  VARCHAR  
Lastname  VARCHAR  
Email  VARCHAR 

