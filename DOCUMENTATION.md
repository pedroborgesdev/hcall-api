# API Documentation

## Resources

- User Management
- Ticket Creation
- Status Updates
- JWT Communication

## Base URL

```
domain/api
```

## Authentication

### Login
- **Endpoint:** `POST /auth/enter`
- **Description:** Authenticates a user in the system
- **Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been loged in",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Email Not Registered (400):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```
  - Incorrect Password (400):
```json
{
    "message": "Password is incorrect",
    "status": false
}
```

### Register
- **Endpoint:** `POST /auth/register`
- **Description:** Registers a new user in the system
- **Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been registered",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Email Already Exists (400):
```json
{
    "message": "Email already exists",
    "status": false
}
```
  - Invalid Password (400):
```json
{
    "message": "Password is invalid",
    "status": false
}
```

## Users

### Get User Information
- **Endpoint:** `GET /user/fetch`
- **Description:** Retrieves information about a specific user or lists all users
- **Query Parameters:**
  - `email`: User's email (optional, example: `email=johndoe@example.com`)
- **Examples:**
  - List all users: `/user/fetch`
  - Get specific user: `/user/fetch?email=johndoe@example.com`
- **Responses:**
  - Success (200) - Specific User:
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "********",
    "created_at": "2023-01-01T12:00:00Z",
    "status": true
}
```
  - Success (200) - User List:
```json
{
    "users": [
        {
            "username": "John Doe",
            "email": "johndoe@example.com"
        },
        {
            "username": "Jane Smith",
            "email": "janesmith@example.com"
        }
    ],
    "status": true
}
```
  - User Not Found (404):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```

### Create User
- **Endpoint:** `POST /user/create`
- **Description:** Creates a new user in the system
- **Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "User has been created",
    "status": true
}
```
  - Email Already Exists (400):
```json
{
    "message": "Email already exists",
    "status": false
}
```
  - Invalid Data (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Remove User
- **Endpoint:** `POST /user/delete`
- **Description:** Removes a user from the system
- **Request Body:**
```json
{
    "email": "johndoe@example.com"
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
    "message": "Email aren't registered",
    "status": false
}
```

## Tickets

### Create Ticket
- **Endpoint:** `POST /ticket/create`
- **Description:** Creates a new technical support ticket
- **Request Body:**
```json
{
    "author_email": "johndoe@example.com",
    "ticket_name": "Printer Problem",
    "ticket_label": "Hardware",
    "ticket_equipment": "HP Printer",
    "ticket_explain": "The printer is not connecting to the network"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Ticket has been created",
    "status": true
}
```
  - Invalid Data (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```
  - Author Not Found (400):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```

### Update Ticket Status
- **Endpoint:** `POST /ticket/edit`
- **Description:** Updates the status of an existing ticket
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
    "message": "Ticket has been edited",
    "status": true
}
```
  - Ticket Not Found (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
  - Invalid Status (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Update Ticket History
- **Endpoint:** `POST /ticket/update`
- **Description:** Adds a new entry to the ticket's history
- **Request Body:**
```json
{
    "ticket_id": "ticket_028492wsd88178",
    "ticket_return": "Purchasing routers"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Update has been setting up",
    "status": true
}
```
  - Ticket Not Found (404):
```json
{
    "message": "Ticket hasn't found",
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
- **Query Parameters:**
  - `author`: Author's email (optional, example: `author=johndoe@example.com`)
  - `status`: Ticket status (optional, example: `status=pending`)
- **Valid Status Values:** `pending`, `doing`, `conclued`
- **Notes:**
  - If `status` parameter is not provided, tickets of all statuses will be returned
  - If `author` parameter is not provided, tickets from all authors will be returned
- **Examples:**
  - List all tickets: `/ticket/fetch`
  - List pending tickets: `/ticket/fetch?status=pending`
  - List tickets by author: `/ticket/fetch?author=johndoe@example.com`
  - List pending tickets by author: `/ticket/fetch?author=johndoe@example.com&status=pending`
- **Responses:**
  - Success (200):
```json
{
    "tickets": [
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
            "ticket_name": "Printer Problem",
            "ticket_status": "pending"
        },
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174001",
            "ticket_name": "Monitor Problem",
            "ticket_status": "doing"
        }
    ],
    "status": true
}
```
  - No Tickets Found (404):
```json
{
    "message": "Author don't have tickets",
    "status": false
}
```

### Get Ticket Information
- **Endpoint:** `GET /ticket/info`
- **Description:** Retrieves detailed information about a specific ticket, including its complete history
- **Query Parameters:**
  - `ticket_id`: Ticket ID (required, example: `ticket_id=ticket_123e4567-e89b-12d3-a456-426614174000`)
- **Examples:**
  - Get ticket information: `/ticket/info?ticket_id=ticket_123e4567-e89b-12d3-a456-426614174000`
- **Responses:**
  - Success (200):
```json
{
    "ticket_id": "ticket_028492wsd88178",
    "ticket_name": "Router Problem",
    "ticket_status": "doing",
    "ticket_explain": "Need to purchase new routers",
    "history": [
        {
            "ticket_return": "Purchasing routers",
            "ticket_date": "2023-07-15T14:30:45Z"
        },
        {
            "ticket_return": "Configuring notebook",
            "ticket_date": "2023-07-16T09:15:22Z"
        },
        {
            "ticket_return": "Clean screen",
            "ticket_date": "2023-07-17T11:45:30Z"
        }
    ],
    "status": true
}
```
  - Ticket Not Found (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
  - ID Not Provided (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Remove Ticket
- **Endpoint:** `POST /ticket/remove`
- **Description:** Removes a ticket from the system. Only the ticket author can remove it.
- **Request Body:**
```json
{
    "author_email": "johndoe@example.com",
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
  - Ticket Not Found or Incorrect Author (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
  - Invalid Data (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
``` 