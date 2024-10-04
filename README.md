# EventManager
Am boutta readme 

# Concepts
## Process
![image](https://github.com/user-attachments/assets/8a319b5d-f45f-47e2-beca-79070e116fe3)


## Microservices Idea
![image](https://github.com/user-attachments/assets/7d39198e-47e7-498f-9850-06615f22271d)

# API Endpoints
## Event
/create  
/read  
/update  
/check_available_tickets  

### Database Data
ID  INT  PRIMARY KEY  
Name  VARCHAR  
Location VARCHAR  
Organisator VARCHAR  
StartDate  DateTime  
Available_tickets  INT  
Available_vip_tickets  INT  

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
/check_balance  
/pay  
/add_money

### Database Data
ID INT  PRIMARY KEY  
User INT  
Bank VARCHAR  
Balance INT  

> Needs Data from Ticket- and Userservice

## Notification

### GET /notification
**Response:**
```json
{
  "code": 200,
  "response":
    [
      {
        "id": 68,
        "description": "Very very very very very very long string, it may be called text by now.",
        "date": 1728035912,
        "eventId": 123,
        "paymentId": 123,
        "ticketId": 1234
      },
      {
        "id": 70,
        "description": "Less long string.",
        "date": 1728035912,
        "eventId": 123,
        "paymentId": 123,
        "ticketId": 1234
      },
    ],
}
```

### GET /notification/\<id\>
**Response:**
```json
{
  "code": 200,
  "response":
    [
      {
        "id": 68,
        "description": "Very very very very very very long string, it may be called text by now.",
        "date": 1728035912,
        "eventId": 123,
        "paymentId": 123,
        "ticketId": 1234
      }
    ],
}
```


### POST /notification  
**Request:**
```json
{
  "userId": 1,
  "description": "Very very very very very very long string",
  "date": 1728035912
}
```
> Note the timestamp is in seconds, the example timestamp is somewhere around 12:00 04.10.2024  
**Response:**
```json
{
  "code": 201,
  "response": "OK",
  "affectedId": 123,
}
```


### Database Data
ID INT  PRIMARY KEY  AUTO INCREMENT  
Description  Text  
Date  Datetime  
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

