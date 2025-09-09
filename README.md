# Slab.ai

## Environment Configuration

`.env` for each services in backend

```bash
MONGO_URI="mongodb://localhost:27017/slabai"
PORT=3003
RAZORPAY_KEY_ID="rzp_test_a6CEBoBbltCvzC"
RAZORPAY_KEY_SECRET="bq1a2NzTCviAwGeNCw9pHg43"
AWS_SECRET_ACCESS_KEY="asfasdf"
AWS_ACCESS_KEY_ID="asdsaf"
AWS_REGION="ap-south-1"
S3_BUCKET_NAME="slabai"
DOWNLOAD_URL="http://localhost:3003"
```

## User Service API Documentation

Base URL: `/api/user-leads`

### Endpoints

#### 1. Create User Lead
- **Method**: POST
- **Endpoint**: `/`
- **Description**: Creates a new user lead in the system
- **Request Body**:
  ```json
  {
    "name": "string",
    "countryCode": "string (+XX format)",
    "phoneNo": "string (10 digits)",
    "email": "string (valid email)"
  }
  ```
- **Response**:
  - Success (201):
    ```json
    {
      "id": "string",
      "name": "string",
      "countryCode": "string",
      "phoneNo": "string",
      "email": "string",
      "creationDateTime": "date"
    }
    ```
  - Error (400):
    ```json
    {
      "message": "Error message"
    }
    ```

#### 2. Get All User Leads
- **Method**: GET
- **Endpoint**: `/`
- **Description**: Retrieves all user leads from the system
- **Response**:
  - Success (200): Array of user lead objects
    ```json
    [
      {
        "id": "string",
        "name": "string",
        "countryCode": "string",
        "phoneNo": "string",
        "email": "string",
        "creationDateTime": "date"
      }
    ]
    ```
  - Error (500):
    ```json
    {
      "message": "Error message"
    }
    ```

#### 3. Get User Lead by ID
- **Method**: GET
- **Endpoint**: `/:id`
- **Parameters**: 
  - `id`: User lead ID (path parameter)
- **Response**:
  - Success (200):
    ```json
    {
      "id": "string",
      "name": "string",
      "countryCode": "string",
      "phoneNo": "string",
      "email": "string",
      "creationDateTime": "date"
    }
    ```
  - Error (404):
    ```json
    {
      "message": "User lead not found"
    }
    ```

#### 4. Update User Lead
- **Method**: PUT
- **Endpoint**: `/:id`
- **Parameters**: 
  - `id`: User lead ID (path parameter)
- **Request Body**:
  ```json
  {
    "name": "string (optional)",
    "countryCode": "string (optional)",
    "phoneNo": "string (optional)",
    "email": "string (optional)"
  }
  ```
- **Response**:
  - Success (200):
    ```json
    {
      "id": "string",
      "name": "string",
      "countryCode": "string",
      "phoneNo": "string",
      "email": "string",
      "creationDateTime": "date"
    }
    ```
  - Error (404):
    ```json
    {
      "message": "User lead not found"
    }
    ```

#### 5. Delete User Lead
- **Method**: DELETE
- **Endpoint**: `/:id`
- **Parameters**: 
  - `id`: User lead ID (path parameter)
- **Response**:
  - Success (200):
    ```json
    {
      "message": "User lead deleted successfully"
    }
    ```
  - Error (404):
    ```json
    {
      "message": "User lead not found"
    }
    ```

### Data Validation Rules
1. **Country Code**:
   - Must be in format: +XX (e.g., +91)
   - Must contain 1-3 digits after '+'

2. **Phone Number**:
   - Must be exactly 10 digits
   - Only numeric characters allowed

3. **Email**:
   - Must be a valid email format
   - Will be stored in lowercase
   - Must be unique in the system

### Notes
- All timestamps are in ISO format
- All string fields are automatically trimmed
- Email addresses are automatically converted to lowercase
- The API returns appropriate HTTP status codes for different scenarios
- All endpoints return JSON responses
