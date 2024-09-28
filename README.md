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
Name  VARCHARR
Location VARCHARR

StartDate  DateTime
Available_tickets  INT
Available_vip_tickets  INT

Needs to know if the user is admin or not.

## Ticket
/create_ticket   
/delete or /update  
/read

Needs Data from Event- and Userservice

## Payment
/check_balance  
/pay  
/add_money

Needs Data from Ticket- and Userservice

## Notification
/create  
/get_user_notifications  

Needs Data from Ticket-, Event- and Paymentservice

## Feedback
/create  
/read  
/update  
/delete  

Needs Data from Ticket-, Event- and Userservice.
## User
/create  
/update  
/read  
(/delete)  

