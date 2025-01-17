// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "post": {
                "tags": [
                    "accounts"
                ],
                "summary": "Create Account",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.CreateAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{user_id}": {
            "get": {
                "tags": [
                    "accounts"
                ],
                "summary": "Get Account Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account details",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    },
                    "404": {
                        "description": "Account not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{user_id}/payments": {
            "get": {
                "tags": [
                    "payments"
                ],
                "summary": "Get Payments by Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of payments",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Payment"
                            }
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Service is healthy",
                        "schema": {
                            "$ref": "#/definitions/models.HealthCheckResponse"
                        }
                    }
                }
            }
        },
        "/payments": {
            "post": {
                "tags": [
                    "payments"
                ],
                "summary": "Create Payment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unique key to prevent duplicate payments",
                        "name": "Idempotency-Key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Payment details",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Payment created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.CreatePaymentResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request or missing idempotency key",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Duplicate idempotency key",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payments/{id}": {
            "get": {
                "tags": [
                    "payments"
                ],
                "summary": "Get Payment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Payment details",
                        "schema": {
                            "$ref": "#/definitions/models.Payment"
                        }
                    },
                    "400": {
                        "description": "Invalid payment ID",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Payment not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Account": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "balance": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.AccountRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateAccountResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.CreatePaymentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "recipient_id",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "payment_reference": {
                    "type": "string"
                },
                "processed_at": {
                    "type": "string"
                },
                "recipient_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.PaymentRequest": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "recipient_id",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                },
                "payment_reference": {
                    "type": "string"
                },
                "recipient_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Payment Service API",
	Description:      "API for managing payments",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
