# Payment Service API
API for managing payments

## Version: 1.0


### /accounts

#### POST
##### Summary:

Create Account

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| account | body | Account ID | Yes | [models.AccountRequest](#models.AccountRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Account created successfully | [models.CreateAccountResponse](#models.CreateAccountResponse) |
| 400 | Invalid request | [models.ErrorResponse](#models.ErrorResponse) |

### /accounts/{user_id}

#### GET
##### Summary:

Get Account Details

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| user_id | path | User ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Account details | [models.Account](#models.Account) |
| 404 | Account not found | [models.ErrorResponse](#models.ErrorResponse) |
| 500 | Internal Server Error | [models.ErrorResponse](#models.ErrorResponse) |

### /accounts/{user_id}/payments

#### GET
##### Summary:

Get Payments by Account

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| user_id | path | UserID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | List of payments | [ [models.Payment](#models.Payment) ] |
| 404 | User not found | [models.ErrorResponse](#models.ErrorResponse) |
| 500 | Internal Server Error | [models.ErrorResponse](#models.ErrorResponse) |

### /health

#### GET
##### Summary:

Health Check

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Service is healthy | [models.HealthCheckResponse](#models.HealthCheckResponse) |

### /payments

#### POST
##### Summary:

Create Payment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Idempotency-Key | header | Unique key to prevent duplicate payments | Yes | string |
| payment | body | Payment details | Yes | [models.PaymentRequest](#models.PaymentRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Payment created successfully | [models.CreatePaymentResponse](#models.CreatePaymentResponse) |
| 400 | Invalid request or missing idempotency key | [models.ErrorResponse](#models.ErrorResponse) |
| 409 | Duplicate idempotency key | [models.ErrorResponse](#models.ErrorResponse) |

### /payments/{id}

#### GET
##### Summary:

Get Payment

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Payment ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Payment details | [models.Payment](#models.Payment) |
| 400 | Invalid payment ID | [models.ErrorResponse](#models.ErrorResponse) |
| 404 | Payment not found | [models.ErrorResponse](#models.ErrorResponse) |
| 500 | Internal Server Error | [models.ErrorResponse](#models.ErrorResponse) |

### Models


#### models.Account

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| balance | number |  | No |
| createdAt | string |  | No |
| currency | string |  | No |
| user_id | string |  | Yes |

#### models.AccountRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| user_id | string |  | No |

#### models.CreateAccountResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| user_id | string |  | No |

#### models.CreatePaymentResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| status | string |  | No |

#### models.ErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### models.HealthCheckResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### models.Payment

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | Yes |
| created_at | string |  | No |
| currency | string |  | Yes |
| id | string |  | No |
| payment_reference | string |  | No |
| processed_at | string |  | No |
| recipient_id | string |  | Yes |
| status | string |  | No |
| user_id | string |  | Yes |

#### models.PaymentRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | Yes |
| currency | string |  | Yes |
| payment_reference | string |  | No |
| recipient_id | string |  | Yes |
| user_id | string |  | Yes |