# Payment Service API

API for managing payments.

## Version: 1.0

### Base Path
`/api/v1`

<br><br>

# Accounts

## Create Account
**POST** `/accounts`  

#### Request:
- **Body**:  
  ```json
  {
    "user_id": "string"
  }
  ```
- **Header**:  
  - `Authorization` (required): Authorization token  

#### Responses:
- **200**: Account created successfully  
  ```json
  {
    "message": "string",
    "user_id": "string"
  }
  ```

- **Error Responses**
    - **400**: Invalid request
    - **401**: Unauthorized
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```

## Get Account Details
**GET** `/accounts/{user_id}`  

#### Request:
- **Path Parameters**:  
  - `user_id` (required): User ID  
- **Header**:  
  - `Authorization` (required): Authorization token  

#### Responses:
- **200**: Account details  
  ```json
  {
    "balance": 0,
    "createdAt": "string",
    "currency": "string",
    "user_id": "string"
  }
  ```

- **Error Responses**
    - **401**: Unauthorized
    - **404**: Account not found
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```

<br><br>

# Payments

## Get Payments by Account
**GET** `/accounts/{user_id}/payments`  

#### Request:
- **Path Parameters**:  
  - `user_id` (required): User ID  
- **Header**:  
  - `Authorization` (required): Authorization token  

#### Responses:
- **200**: List of payments  
  ```json
  [
    {
      "amount": 0,
      "created_at": "string",
      "currency": "string",
      "id": "string",
      "payment_reference": "string",
      "processed_at": "string",
      "recipient_id": "string",
      "status": "string",
      "user_id": "string"
    }
  ]
  ```

- **Error Responses**
    - **401**: Unauthorized
    - **404**: User not found
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```

## Get Payment
**GET** `/accounts/{user_id}/payments/{id}`  

#### Request:
- **Path Parameters**:  
  - `user_id` (required): User ID  
  - `id` (required): Payment ID  
- **Header**:  
  - `Authorization` (required): Authorization token  

#### Responses:
- **200**: Payment details  
  ```json
  {
    "amount": 0,
    "created_at": "string",
    "currency": "string",
    "id": "string",
    "payment_reference": "string",
    "processed_at": "string",
    "recipient_id": "string",
    "status": "string",
    "user_id": "string"
  }
  ```

- **Error Responses**
    - **400**: Invalid payment ID
    - **401**: Unauthorized
    - **404**: Payment not found
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```

## Create Payment
**POST** `/payments`  

#### Request:
- **Header**:  
  - `Idempotency-Key` (required): Unique key to prevent duplicate payments  
  - `Authorization` (required): Authorization token  
- **Body**:  
  ```json
  {
    "amount": 0,
    "currency": "string",
    "payment_reference": "string",
    "recipient_id": "string",
    "user_id": "string"
  }
  ```

#### Responses:
- **201**: Payment created successfully  
  ```json
  {
    "id": "string",
    "status": "string"
  }
  ```

- **Error Responses**
    - **400**: Invalid request or missing idempotency key
    - **401**: Unauthorized
    - **409**: Duplicate idempotency key
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```

<br><br>

# Health

## Health Check
**GET** `/health`  

#### Responses:
- **200**: Service is healthy  
  ```json
  {
    "message": "string"
  }
  ```

- **Error Responses**
    - **500**: Internal Server Error
  ```json
  {
    "error": "string"
  }
  ```