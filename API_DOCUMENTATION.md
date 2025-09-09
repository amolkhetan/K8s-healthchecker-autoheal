# API Documentation

## Table of Contents
- [Payment Service](#payment-service)
- [Project Service](#project-service)
- [User Service](#user-service)

## Payment Service
Base URL: `/api/payment`

### POST /create-order
Creates a new payment order.

**Request Body:**
```json
{
  "amount": number,
  "currency": string,
  "projectId": string
}
```

**Response:**
```json
{
  "orderId": string,
  "currency": string,
  "amount": number
}
```

### POST /verify-payment
Verifies a payment after successful completion.

**Request Body:**
```json
{
  "orderCreationId": string,
  "razorpayPaymentId": string,
  "razorpayOrderId": string,
  "razorpaySignature": string,
  "user": {
    "name": string,
    "email": string,
    "phone": string,
    "amount": number,
    "projectId": string
  }
}
```

**Response:**
```json
{
  "msg": "success",
  "orderId": string,
  "paymentId": string,
  "paymentRecordId": string,
  "downloadLink": string
}
```

### POST /send-download-links
Sends all download links associated with a user's email.

**Request Body:**
```json
{
  "email": string
}
```

**Response:**
```json
{
  "message": "Download links sent successfully!"
}
```

### GET /stream/:paymentId
Streams content for a specific payment.

**Parameters:**
- paymentId: Payment ID to stream content for

## Project Service

### Project Domains
Base URL: `/api/project-domains`

#### POST /
Creates a new project domain.

**Request Body:**
```json
{
  "name": string,
  "description": string
}
```

#### GET /
Retrieves all project domains.

#### GET /:id
Retrieves a specific project domain.

**Parameters:**
- id: Project domain ID

#### PUT /:id
Updates a specific project domain.

**Parameters:**
- id: Project domain ID

**Request Body:**
```json
{
  "name": string,
  "description": string
}
```

### Project Content
Base URL: `/api/project-content`

#### POST /
Creates new project content.

**Request Body:**
```json
{
  "projectId": string,
  "downloadLink": string
}
```

#### GET /
Retrieves all project links.

#### PUT /:id
Updates specific project content.

**Parameters:**
- id: Project content ID

**Request Body:**
```json
{
  "downloadLink": string
}
```

### Projects
Base URL: `/api/projects`

#### POST /
Creates a new project.

**Request Body:**
```json
{
  "title": string,
  "description": string,
  "price": number,
  "domain": string,
  "features": string[],
  "technologies": string[]
}
```

#### GET /
Retrieves all projects.

#### GET /:id
Retrieves a specific project.

**Parameters:**
- id: Project ID

#### PUT /:id
Updates a specific project.

**Parameters:**
- id: Project ID

## User Service
Base URL: `/api/user-leads`

### POST /
Creates a new user lead.

**Request Body:**
```json
{
  "name": string,
  "email": string,
  "phone": string,
  "message": string
}
```

### GET /
Retrieves all user leads.

### GET /:id
Retrieves a specific user lead.

**Parameters:**
- id: User lead ID

### PUT /:id
Updates a specific user lead.

**Parameters:**
- id: User lead ID

**Request Body:**
```json
{
  "name": string,
  "email": string,
  "phone": string,
  "message": string
}
```

### DELETE /:id
Deletes a specific user lead.

**Parameters:**
- id: User lead ID

## Error Responses

All endpoints may return the following error responses:

### 400 Bad Request
```json
{
  "msg": "Error message describing the issue"
}
```

### 404 Not Found
```json
{
  "message": "Resource not found message"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error message"
}
```

## Authentication

Some endpoints may require authentication. When required, include the authentication token in the request header:

```
Authorization: Bearer <token>
```

## Rate Limiting

API endpoints may have rate limiting applied to prevent abuse. Please contact the system administrator for specific limits.

## Notes

1. All requests should include appropriate Content-Type headers (usually `application/json`).
2. All responses are in JSON format.
3. Dates are returned in ISO 8601 format.
4. All monetary values are in the smallest currency unit (e.g., cents for USD).
