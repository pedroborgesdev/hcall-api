# HCall API Documentation

## Resources

- User Management
- Ticket Creation
- Status Updates
- JWT Communication

## Base URL

```
domain/api
```

## Environment Configuration

The API relies on environment variables for configuration. These can be set in a `.env` file in the root directory. Below are the key environment variables used:

### Database Configuration
- `DB_HOST`: Database host (default: "localhost")
- `DB_PORT`: Database port (default: "5432")
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_SSLMODE`: SSL mode for database connection (default: "disable")

### Security Settings
- `USERNAME_MIN_CHAR`: Minimum characters for username (default: 6)
- `PASSWORD_MIN_CHAR`: Minimum characters for password (default: 8)
- `PASSWORD_SPECIAL`: Require special characters in password (default: True)
- `PASSWORD_DIGITS`: Require digits in password (default: True)
- `PASSWORD_UPPERCASE`: Require uppercase letters in password (default: True)
- `PASSWORD_LOWERCASE`: Require lowercase letters in password (default: True)

### JWT Configuration
- `JWT_SECRET`: Secret key used for JWT token signing
- `JWT_EXPIRATION_HOURS`: Hours until JWT token expires (default: 24)

### Worker Configuration
- `WORKER_TICKET_LOOPTIME`: Hours between ticket worker runs (default: 24)
- `WORKER_TICKET_REMOVE_AFTER`: Days after which to remove tickets (default: 30)
- `WORKER_TICKET_REMOVE_STATUS`: Status of tickets to remove (default: "conclued")

### Server Configuration
- `PORT`: Port on which to run the API server (default: 8080)

## Authentication Requirements

All API endpoints require a valid JWT token in the Authorization header except for the `/auth` endpoints and `/master` endpoints. The token must be included as a Bearer token in the following format:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

JWT tokens are obtained through the login or register endpoints. These tokens expire after 24 hours and must be renewed by logging in again.

The JWT token is used to identify the user making the request. All actions performed by the API will be associated with the user identified by the token, so there is no need to send user identification in the request body.

## User Roles and Permissions

The API supports three user roles with different access levels:

1. **User**
   - Default role for regular users
   - Access limited to:
     - Authentication endpoints (`/auth/*`)
     - Creating tickets (`/ticket/create`)
     - Removing tickets (`/ticket/remove`)
     - Viewing ticket count (`/ticket/count`)

2. **Admin**
   - Administrative role with broader access
   - Access to all endpoints except:
     - Can only create users with the "user" role
   - When using `/user/create`, can only create users with the "user" role

3. **Master**
   - Highest privileged role
   - Full access to all endpoints
   - Created using the `/master/create` endpoint

## Master Creation

### Create Master User
- **Endpoint:** `POST /master/create`
- **Description:** Creates a user with master privileges
- **Authorized Roles:** None (public endpoint)
- **Note:** This endpoint only works when there is no master user registered in the system yet. Once a master user exists, the endpoint will be disabled.
- **Request Body:**
```json
{
    "master_email": "master@example.com",
    "master_password": "StrongPassword123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Master user has been created",
    "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Master Already Exists (403):
```json
{
    "message": "Master user already exists",
    "reason": "error message",
    "status": false
}
```

### Delete Master User
- **Endpoint:** `POST /master/delete`
- **Description:** Removes a user with master privileges
- **Authorized Roles:** None (public endpoint)
- **Note:** This endpoint requires valid master credentials for the account to be deleted.
- **Request Body:**
```json
{
    "master_email": "master@example.com",
    "master_password": "StrongPassword123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Master user has been deleted",
    "status": true
}
```
  - Invalid Credentials (403):
```json
{
    "message": "Invalid master password",
    "reason": "error message",
    "status": false
}
```
  - Master Not Found (404):
```json
{
    "message": "Master user not found",
    "reason": "error message",
    "status": false
}
```

## Authentication

### Login
- **Endpoint:** `POST /auth/enter`
- **Description:** Authenticates a user in the system
- **Authorized Roles:** None (public endpoint)
- **Request Body:**
```json
{
    "user_email": "johndoe@example.com",
    "user_password": "Password123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been logged in",
    "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Email Not Registered/Incorrect Password (400):
```json
{
    "message": "Invalid credentials",
    "reason": "error message",
    "status": false
}
```

### Register
- **Endpoint:** `POST /auth/register`
- **Description:** Registers a new user in the system
- **Authorized Roles:** None (public endpoint)
- **Request Body:**
```json
{
    "user_name": "John Doe",
    "user_email": "johndoe@example.com",
    "user_password": "Password123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been registered",
    "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Email Already Exists (400):
```json
{
    "message": "Error message",
    "reason": "error details",
    "status": false
}
```

## Users

### Get User Information
- **Endpoint:** `GET /user/fetch`
- **Description:** Retrieves information about a specific user or lists all users
- **Authorized Roles:** `admin`, `master`
- **Query Parameters:**
  - `email`: User's email (optional, example: `email=johndoe@example.com`)
  - `role`: User's role (optional, example: `role=admin`)
- **Valid Role Values:** `user`, `admin`, `master`
- **Examples:**
  - List all users: `/user/fetch`
  - Get specific user: `/user/fetch?email=johndoe@example.com`
  - List users with specific role: `/user/fetch?role=admin`
  - List specific user with role: `/user/fetch?email=johndoe@example.com&role=admin`
- **Notes:**
  - If `role` parameter is provided, only users with that role will be returned
  - If both `email` and `role` parameters are provided, the user must match both criteria
- **Responses:**
  - Success (200) - Specific User:
```json
{
    "user_name": "John Doe",
    "user_email": "johndoe@example.com",
    "user_password": "********",
    "user_created_at": "2023-01-01T12:00:00Z",
    "user_role": "user",
    "status": true
}
```
  - Success (200) - User List:
```json
{
    "users": [
        {
            "user_name": "John Doe",
            "user_email": "johndoe@example.com",
            "user_role": "user"
        },
        {
            "user_name": "Jane Smith",
            "user_email": "janesmith@example.com",
            "user_role": "admin"
        }
    ],
    "status": true
}
```
  - User Not Found (404):
```json
{
    "message": "Email aren't registered",
    "reason": "error message",
    "status": false
}
```
  - No Users Found (404):
```json
{
    "message": "No users found with specified role",
    "reason": "error message",
    "status": false
}
```

### Create User
- **Endpoint:** `POST /user/create`
- **Description:** Creates a new user in the system
- **Authorized Roles:** `admin`, `master`
- **Notes:** 
  - Admin users can only create users with the "user" role
  - Master users can create users with any role (user or admin)
- **Request Body:**
```json
{
    "user_name": "John Doe",
    "user_email": "johndoe@example.com",
    "user_password": "Password123",
    "user_role": "user"
}
```
- **Valid Role Values:** `user`, `admin` (Master users only)
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been created",
    "status": true
}
```
  - Error (400):
```json
{
    "message": "User creation failed",
    "reason": "error message",
    "status": false
}
```

### Delete User
- **Endpoint:** `POST /user/delete`
- **Description:** Removes a user from the system
- **Authorized Roles:** `admin`, `master`
- **Request Body:**
```json
{
    "user_email": "johndoe@example.com"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been deleted",
    "status": true
}
```
  - User Not Found (404):
```json
{
    "message": "User not found",
    "reason": "error message",
    "status": false
}
```
  - Internal Server Error (500):
```json
{
    "message": "Internal server error",
    "reason": "error message",
    "status": false
}
```

## Tickets

### Create Ticket
- **Endpoint:** `POST /ticket/create`
- **Description:** Creates a new ticket in the system
- **Authorized Roles:** `user`
- **Request Body:**
```json
{
    "ticket_name": "Router Problem",
    "ticket_explain": "Need to configure the router in room 302",
    "ticket_images": [
        {
            "image_name": "router_front.jpg",
            "image_content": "base64_encoded_image_data...",
            "image_type": "image/jpeg"
        },
        {
            "image_name": "router_back.jpg",
            "image_content": "base64_encoded_image_data...",
            "image_type": "image/jpeg"
        }
    ]
}
```
- **Notes:**
  - The `ticket_images` field is optional and can contain multiple images
  - Each image must include `image_name`, `image_content` (base64 encoded), and `image_type` fields
  - Supported image types: `image/jpeg`, `image/png`, `image/gif`
- **Responses:**
  - Success (200):
```json
{
    "message": "Ticket has been created",
    "status": true
}
```
  - Error (400):
```json
{
    "message": "Ticket creation failed",
    "reason": "error message",
    "status": false
}
```

### Update Ticket Status
- **Endpoint:** `POST /ticket/edit`
- **Description:** Updates the status of an existing ticket
- **Authorized Roles:** `admin`, `master`
- **Request Body:**
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
    "ticket_status": "doing"
}
```
- **Valid Status Values:** `pending`, `doing`, `conclued`
- **Responses:**
  - Success (200):
```json
{
    "message": "Ticket status has been updated",
    "status": true
}
```
  - Error (400):
```json
{
    "message": "Ticket status update failed",
    "reason": "error message",
    "status": false
}
```

### Update Ticket History
- **Endpoint:** `POST /ticket/update`
- **Description:** Adds a new entry to the ticket's history
- **Authorized Roles:** `admin`, `master`
- **Request Body:**
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
    "ticket_return": "Purchasing routers"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Ticket history has been added",
    "status": true
}
```
  - Error (400):
```json
{
    "message": "Ticket history add failed",
    "reason": "error message",
    "status": false
}
```
- **Notes:**
  - Updated history can be viewed through the `/ticket/info` endpoint
  - Each update is recorded with date and time
  - History is displayed in chronological order

### List Tickets
- **Endpoint:** `GET /ticket/fetch`
- **Description:** Lists tickets from a specific author or all system tickets
- **Authorized Roles:** `admin`, `master`
- **Query Parameters:**
  - `author`: Author's email (optional, example: `author=johndoe@example.com`)
  - `status`: Ticket status (optional, example: `status=pending`)
  - `date`: Tickets that were created after the date (optional, example: `date=2025-03-27`)
  - `name`: Ticket name for filtering (optional, example: `name=Router`)
- **Valid Status Values:** `pending`, `doing`, `conclued`
- **Notes:**
  - If `status` parameter is not provided, tickets of all statuses will be returned
  - If `author` parameter is not provided, tickets from all authors will be returned
- **Examples:**
  - List all tickets: `/ticket/fetch`
  - List pending tickets: `/ticket/fetch?status=pending`
  - List tickets by author: `/ticket/fetch?author=johndoe@example.com`
  - List pending tickets by author: `/ticket/fetch?author=johndoe@example.com&status=pending`
  - List tickets by author created after the date: `/ticket/fetch?author=johndoe@example.com&date=2025-01-20`
  - List tickets by name: `/ticket/fetch?name=Router`
- **Responses:**
  - Success (200):
```json
{
    "tickets": [
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
            "ticket_name": "Printer Problem",
            "ticket_status": "pending",
            "ticket_author": "John Doe",
            "ticket_date": "2023-07-15T14:30:45Z"
        },
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174001",
            "ticket_name": "Monitor Problem",
            "ticket_status": "doing",
            "ticket_author": "Jane Smith",
            "ticket_date": "2023-07-16T09:15:22Z"
        }
    ],
    "status": true
}
```
  - No Tickets Found (404):
```json
{
    "message": "No tickets found",
    "status": false
}
```
  - Invalid Date Format (404):
```json
{
    "message": "Invalid date format",
    "reason": "error message",
    "status": false
}
```

### Get Tickets Count
- **Endpoint:** `GET /ticket/count`
- **Description:** Lists tickets count by status
- **Authorized Roles:** `user`, `admin`, `master`
- **Responses:**
  - Success (200):
```json
{
    "tickets_total": 5,
    "tickets_pending": 2,
    "tickets_doing": 3,
    "tickets_conclued": 0,
    "status": true
}
```
  - Error (404):
```json
{
    "message": "Ticket not found",
    "reason": "error message",
    "status": false
}
```

### Get Ticket Information
- **Endpoint:** `GET /ticket/info`
- **Description:** Retrieves detailed information about a specific ticket, including its complete history
- **Authorized Roles:** `admin`, `master`
- **Query Parameters:**
  - `ticket_id`: Ticket ID (required, example: `ticket_id=ticket_123e4567-e89b-12d3-a456-426614174000`)
- **Responses:**
  - Success (200):
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
    "ticket_name": "Router Problem",
    "tickt_status": "doing",
    "ticket_explain": "Need to purchase new routers",
    "ticket_email": "johndoe@example.com",
    "ticket_images": [
        {
            "image_id": "img_12345678",
            "image_name": "router_front.jpg",
            "image_base64": "base64_encoded_image_data...",
            "image_type": "image/jpeg",
            "image_uploaded_at": "2023-07-15T13:30:22Z"
        },
        {
            "image_id": "img_87654321",
            "image_name": "router_back.png",
            "image_base64": "base64_encoded_image_data...",
            "image_type": "image/png",
            "image_uploaded_at": "2023-07-15T13:30:22Z"
        }
    ],
    "ticket_history": [
        {
            "ticket_return": "Purchasing routers",
            "ticket_date": "2023-07-15T14:30:45Z"
        },
        {
            "ticket_return": "Configuring notebook",
            "ticket_date": "2023-07-16T09:15:22Z"
        }
    ],
    "ticket_date": "2023-07-15T13:30:22Z",
    "status": true
}
```
  - Ticket Not Found (404):
```json
{
    "message": "Ticket hasn't found",
    "reason": "error message",
    "status": false
}
```
  - ID Not Provided (400):
```json
{
    "message": "Invalid data",
    "reason": "Ticket ID is required",
    "status": false
}
```

### Remove Ticket
- **Endpoint:** `POST /ticket/remove`
- **Description:** Removes a ticket from the system
- **Authorized Roles:** `user` (only their own tickets), `admin`, `master`
- **Request Body:**
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Ticket has been removed",
    "status": true
}
```
  - Error (400):
```json
{
    "message": "No permission to delete",
    "reason": "error message",
    "status": false
}
```

