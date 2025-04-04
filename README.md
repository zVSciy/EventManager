# EventManager
The `EventManager`is a software engineering project of HTL Villach's 5AHITS (Year 2024/2025) that provides a management tool for concerts and other kinds of events. 

## Table of Contents
- [Concepts](#concepts)
  - [Process](#process)
- [API Endpoints](#api-endpoints)
  - [Event](#event)
  - [Ticket](#ticket)
  - [Payment](#payment)
  - [Notification](#notification)
  - [Feedback](#feedback)
  - [User](#user)

# Provisioning
This section covers all important point about provisioning and testing the `EventManager`.

## Requirements
* Ubuntu 24.04 Server (recommended)
* Docker Runtime (more security with Docker Rootless)
* Source code of the Github repository

## Start the EventManager
1. Clone the source code from GitHub
```bash
git clone https://github.com/zVSciy/EventManager.git
```

2. Run the stack using `docker-compose-test.yaml`
```bash
docker compose -f docker-compose-test.yaml up [parameters]
```

3. Wait until every container is started and all databases as well as APIs are healthy. After the integration tests had run (exit with code 0), the `EventManager` may be used by opening the website `https://<server-ip>:8080/`

If you are not familiar with Docker, here are some parameters that might be helpful.
* **-d** &rarr; Detached mode; Output of containers is not shown in the command line which causes the CLI to be further accessible.
* **--build** &rarr; All containers, including their environments, are rebuild.

## Running Rootless Docker
To start off, if you have not installed the Docker Runtime yet, mind Docker's installation guide: https://docs.docker.com/engine/install/ubuntu/

The following commands for rootless Docker originate from Docker's official documentation, see: https://docs.docker.com/engine/security/rootless/#install. Keep in mind that the given commands should be executed with the user that is responsible for Docker later on.

1. Install `uidmap`
```bash
sudo apt-get install -y uidmap
``` 
2. Shutdown Docker's system deamon
```bash
sudo systemctl disable --now docker.service docker.socket
```

3. Install `Dockerd-Rootless-Setup`
```bash
dockerd-rootless-setuptool.sh install
```

4. Define the variable `DOCKER_HOST`
```bash
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
```

The Docker engine is now placed in the current user's directory which means that the Docker runtime uses user permissions instead of `root` ones. As a result, the docker commands do not require `sudo` anymore when using a simple user.

## Additional Notes
- Each service runs its Unittest automatically when started.
- Integration tests are defined in `./Integration_Tests`.
- The NGINX proxy resolves the following services:
  - User Management &rarr; `/`
  - Event &rarr; `/app_event`
  - Ticket &rarr; `/app_ticket`
  - Notification &rarr; `/app_notification`
  - Review &rarr; `/app_review` 
- The `Payment` service has not been integrated into other components as it was submitted way too late.

# Concepts
In the following, the process and concept of this application is introduced.
## Process
![image](Process_Detailed.drawio.svg)

The project consists of the following six microservices:
- User (User management)
- Event
- Ticket
- Payment
- Notification
- Review (Feedback)

Each one fulfills crucial functionalities that are combined in order to achieve overall functionality of the `EventManager`.

# API Endpoints
A short documentation of the API endpoint of each service is stated below.
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

### Go to [Payment API Specification](Payment/api-spec.md)


## Notification

### GET /notification
**Request:**
```json
{
  "skip":0,
  "limit":10,
}
```
> skip and limit is optional, default is listed above

**Response:**
```json
[
    {
        "id": 1,
        "timestamp": 1729235305,
        "paymentId": 2,
        "eventId": 1,
        "description": "testdesc",
        "userId":2,
        "status": "active",
        "ticketId": 4
    },
    {
        "id": 2,
        "timestamp": 1729235305,
        "paymentId": 2,
        "eventId": 1,
        "description": "testdesc",
        "userId":2,
        "status": "active",
        "ticketId": 4
    },
    {
        "id": 3,
        "timestamp": 1729235305,
        "paymentId": 2,
        "eventId": 1,
        "description": "testdesc",
        "userId":2,
        "status": "active",
        "ticketId": 4
    }
]
```

### GET /notification/\<id\>
**Response:**
```json
{
    "id": 2,
    "timestamp": 1729235305,
    "paymentId": 2,
    "eventId": 1,
    "description": "testdesc",
    "userId":2,
    "status": "active",
    "ticketId": 4
}
```


### POST /notification  
**Request:**
```json
{
    "description": "testdesc",
    "status": "active",
    "timestamp": "2024-02-17T00:15:05",
    "eventId": "1",
    "paymentId": "2",
    "ticketId": "4",
    "userId":2,
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
    "userId":2,
    "status": "active",
    "ticketId": 4
}
```

### PUT /notification  
**Request:**
```json
{
    "timestamp": 123213,
    "paymentId": 2,
    "userId": 2,
    "eventId": 5,
    "description": "edited",
    "status": "disabled",
    "ticketId": 45
}
```
> Note the timestamp is in seconds, the example timestamp is somewhere around 12:00 04.10.2024  
**Response:**
```json
{
    "id": 2,
    "timestamp": 123213,
    "paymentId": 2,
    "userId": 2,
    "eventId": 5,
    "description": "edited",
    "status": "disabled",
    "ticketId": 45
}
```



### Database Data
id INT  PRIMARY KEY  AUTO INCREMENT  
status   VARCHAR  
description  Text  
timestamp  int  
userId  INT  (References to User.Id)  
eventId INT (References to Event.Id)  
paymentId INT (References to Payment.Id)  
ticketId INT (References to Ticket.Id)  

> Needs Data from Ticket-, Event- and Paymentservice  

## Feedback
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
- ID INT  PRIMARY KEY 
- User  INT
- Comment  Text
- Rating INT
- Event INT  

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

