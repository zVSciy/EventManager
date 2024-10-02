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

### Standard Data
ID  INT
Name  VARCHAR  
Location VARCHAR  
Organisator VARCHAR  
StartDate  DateTime  
Available_tickets  INT  
Available_vip_tickets  INT  

Needs to know if the user is admin or not.

## Ticket
/create_ticket   
/delete or /update  
/read

### Standard Data
ID  INT  
Row  VARCHAR  
Price  INT  
SeatNumber  INT  
VIP  BOOLEAN  
User  INT  
Event INT  

Needs Data from Event- and Userservice

## Payment
/check_balance  
/pay  
/add_money

Needs Data from Ticket- and Userservice

## Notification
/create  
/get_user_notifications  

### Standard Data
ID INT  
User  INT  
Description  Text  

Needs Data from Ticket-, Event- and Paymentservice

## Feedback
/create  
/read  
/update  
/delete  

### Standard Data
ID INT    
User  INT  
Comment  Text  
Event INT  
Date  Datetime  


Needs Data from Ticket-, Event- and Userservice.

## User
/create  
/update  
/read  
(/delete)  

### Standard Data
ID  INT  
Firstname  VARCHAR  
Lastname  VARCHAR  
Email  VARCHAR 


