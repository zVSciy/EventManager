CREATE DATABASE EventManagement;
USE EventManagement;
CREATE TABLE Events (
    ID INT AUTO_INCREMENT PRIMARY KEY,   
    Name VARCHAR(255) NOT NULL,          
    Location VARCHAR(255),               
    Organisator VARCHAR(255),            
    StartDate DATETIME,                  
    Available_normal_tickets INT,        
    Available_vip_tickets INT,           
    Canceled BOOLEAN DEFAULT FALSE       
);
