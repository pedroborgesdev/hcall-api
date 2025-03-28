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

## Authentication Requirements

All API endpoints require a valid JWT token in the Authorization header except for the `/auth` endpoints. The token must be included as a Bearer token in the following format:

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

2. **Admin**
   - Administrative role with broader access
   - Access to all endpoints except:
     - Cannot create tickets
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
    "email": "master@example.com",
    "password": "StrongPassword123"
}
```
- **Responses:**
  - Success (200):
```json
{
    "message": "Master user has been created",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "status": true
}
```
  - Master Already Exists (403):
```json
{
    "message": "Master user already exists",
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
    "email": "master@example.com",
    "password": "StrongPassword123"
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
    "message": "Invalid master credentials",
    "status": false
}
```
  - Master Not Found (404):
```json
{
    "message": "Master user not found",
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
    "message": "User has been loged in",
    "jwt_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
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
    "status": false
}
```
  - No Users Found (404):
```json
{
    "message": "No users found with specified role",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
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
    "message": "Invalid data",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
    "status": false
}
```

### Remove User
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
    "message": "Email aren't registered",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
    "status": false
}
```

## Tickets

### Create Ticket
- **Endpoint:** `POST /ticket/create`
- **Description:** Creates a new ticket in the system
- **Authorized Roles:** `user`, `admin`, `master`
- **Request Body:**
```json
{
    "ticket_name": "Router Problem",
    "ticket_explain": "Need to configure the router in room 302",
    "tickes_images": [
        {
            "image_name": "router_front.jpg",
            "image_content": "base64_encoded_image_data...",
            "image_type": "image/jpeg"
        },
        {
            "image_name": "router_back.jpg",
            "image_content": "base64_encoded_image_data...",
            "image_ype": "image/jpeg"
        }
    ]
}
```
- **Notes:**
  - The `ticket_images` field is optional and can contain multiple images
  - Each image must include `image_name`, `image_content` (base64 encoded), and `image_type` fields
  - Supported image types: `image/jpeg`, `image/png`
  - Maximum file size per image: 5MB
  - Maximum number of images per ticket: 10
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
    "message": "Invalid data",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "Admin role not authorized to create tickets",
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
    "message": "Invalid data",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
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
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
    "status": false
}
```

### List Tickets
- **Endpoint:** `GET /ticket/fetch`
- **Description:** Lists tickets from a specific author or all system tickets
- **Authorized Roles:** `admin`, `master`
- **Query Parameters:**
  - `author`: Author's email (optional, example: `author=johndoe@example.com`)
  - `status`: Ticket status (optional, example: `status=pending`)
  - `date`: Tickets that were created after the date (optional, example: `date=2025-03-27`)
- **Valid Status Values:** `pending`, `doing`, `conclued`
- **Notes:**
  - If `status` parameter is not provided, tickets of all statuses will be returned
  - If `author` parameter is not provided, tickets from all authors will be returned
- **Examples:**
  - List all tickets: `/ticket/fetch`
  - List pending tickets: `/ticket/fetch?status=pending`
  - List tickets by author: `/ticket/fetch?author=johndoe@example.com`
  - List pending tickets by author: `/ticket/fetch?author=johndoe@example.com&status=pending`
  - List tickets by author were created after the date: `/ticket/fetch?author=johndoe@example.com&date=2025-01-20`
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
  - Invalid Date Format (404):
```json
{
    "message": "Invalid date format",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
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
    "ticket_id": "ticket_028492wsd88178",
    "ticket_name": "Router Problem",
    "ticket_status": "doing",
    "ticket_explain": "Need to purchase new routers",
    "ticket_images": [
        {
            "image_id": "img_12345678",
            "image_name": "router_front.jpg",
            "image_url": "base64_encoded_image_data...",
            "image_type": "image/jpeg",
            "image_uploaded_at": "2023-07-15T13:30:22Z"
        },
        {
            "image_id": "img_872394737",
            "image_name": "router_back.png",
            "image_url": "base64_encoded_image_data...",
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
        },
        {
            "ticket_return": "Clean screen",
            "ticket_date": "2023-07-17T11:45:30Z"
        }
    ],
    "ticket_date": "2025-03-27T22:37:14.722128-03:00",
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
    "message": "Invalid data",
    "status": false
}
```
  - Unauthorized Role (403):
```json
{
    "message": "User role not authorized for this endpoint",
    "status": false
}
```

### Remove Ticket
- **Endpoint:** `POST /ticket/remove`
- **Description:** Removes a ticket from the system. Only the ticket author can remove it.
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
    "message": "Invalid data",
    "status": false
}
```
  - Not Ticket Author (403): 
```json
{
    "message": "User can only remove their own tickets",
    "status": false
}
```

## Axios Examples

The following examples demonstrate how to use our API with Axios, a popular HTTP client for JavaScript. Each example includes code for making requests and handling responses.

### Authentication Flow Example

```javascript
const axios = require('axios');

// Function to obtain a JWT token
async function getToken() {
    try {
        const response = await axios.post('domain/api/auth/register', {
            "user_name": "John Doe",
            "user_email": "johndoe@example.com",
            "user_password": "Password123"
        });

        // Extract token from response
        const token = response.data.jwt_token;
        console.log('Token received:', token);

        // Return token for later use
        return token;
    } catch (error) {
        console.error('Error obtaining token:', error);
        throw error;
    }
}

// Function to make authenticated requests using JWT token
async function makeAuthenticatedRequest(token) {
    try {
        const response = await axios.get('domain/api/user/fetch', {
            headers: {
                'Authorization': `Bearer ${token}` // Pass token in Authorization header
            }
        });

        console.log('Response from authenticated request:', response.data);
    } catch (error) {
        console.error('Error in authenticated request:', error);
    }
}

// Execute complete flow
async function completeFlow() {
    try {
        // 1. Get token
        const token = await getToken();

        // 2. Make authenticated request with token
        await makeAuthenticatedRequest(token);
    } catch (error) {
        console.error('Error in complete flow:', error);
    }
}

// Start flow
completeFlow();
```

### Endpoint Examples

#### Login

```javascript
const axios = require('axios');

async function login() {
    try {
        const response = await axios.post('domain/api/auth/enter', {
            "user_email": "johndoe@example.com",
            "user_password": "Password123"
        });
        
        console.log('Login successful:', response.data);
        return response.data.token;
    } catch (error) {
        console.error('Login error:', error.response?.data || error.message);
    }
}

login();
```

#### Create Master User

```javascript
const axios = require('axios');

async function createMasterUser() {
    try {
        const response = await axios.post('domain/api/master/create', {
            "master_email": "admin@example.com",
            "master_password": "StrongPassword123"
        });
        
        console.log('Master user created:', response.data);
        return response.data.jwt_token;
    } catch (error) {
        console.error('Error creating master user:', error.response?.data || error.message);
    }
}

createMasterUser();
```

#### Delete Master User

```javascript
const axios = require('axios');

async function deleteMasterUser() {
    try {
        const response = await axios.post('domain/api/master/delete', {
            "master_email": "admin@example.com",
            "master_password": "StrongPassword123"
        });
        
        console.log('Master user deleted:', response.data);
        return response.data.status;
    } catch (error) {
        console.error('Error deleting master user:', error.response?.data || error.message);
    }
}

deleteMasterUser();
```

#### Create User

```javascript
const axios = require('axios');

async function createUser(token) {
    try {
        const response = await axios.post('domain/api/user/create', 
            {
                "user_name": "John Doe",
                "user_email": "johndoe@example.com",
                "user_password": "Password123",
                "user_role": "user"
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('User created:', response.data);
    } catch (error) {
        console.error('Error creating user:', error.response?.data || error.message);
    }
}

// Example usage
// createUser(jwtToken);
```

#### Get User Information

```javascript
const axios = require('axios');

async function getUserInfo(token, email = null, role = null) {
    try {
        // Build query parameters
        let url = 'domain/api/user/fetch';
        const params = {};
        
        if (email) params.email = email;
        if (role) params.role = role;
        
        const response = await axios.get(url, {
            params: params,
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        console.log('User information:', response.data);
    } catch (error) {
        console.error('Error fetching user info:', error.response?.data || error.message);
    }
}

// Example usage
// getUserInfo(jwtToken); // Get all users
// getUserInfo(jwtToken, 'johndoe@example.com'); // Get specific user
// getUserInfo(jwtToken, null, 'admin'); // Get users with admin role
```

#### Remove User

```javascript
const axios = require('axios');

async function removeUser(token, email) {
    try {
        const response = await axios.post('domain/api/user/delete', 
            {
                "user_email": email
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('User removed:', response.data);
    } catch (error) {
        console.error('Error removing user:', error.response?.data || error.message);
    }
}

// Example usage
// removeUser(jwtToken, 'johndoe@example.com');
```

#### Create Ticket

```javascript
const axios = require('axios');

async function createTicket(token, ticketData) {
    try {
        const response = await axios.post('domain/api/ticket/create', 
            ticketData,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('Ticket created:', response.data);
    } catch (error) {
        console.error('Error creating ticket:', error.response?.data || error.message);
    }
}

// Example usage
// const ticketData = {
//     "ticket_name": "Router Problem",
//     "ticket_explain": "Need to configure the router in room 302",
//     "ticket_images": [
//         {
//             "image_name": "router_front.jpg",
//             "image_content": "[base64 encoded image data]",
//             "image_type": "image/jpeg"
//         },
//         {
//             "image_name": "router_back.jpg", 
//             "image_content": "[base64 encoded image data]",
//             "image_type": "image/jpeg"
//         }
//     ]
// };
// createTicket(jwtToken, ticketData);
```

#### Update Ticket Status

```javascript
const axios = require('axios');

async function updateTicketStatus(token, ticketId, status) {
    try {
        const response = await axios.post('domain/api/ticket/edit', 
            {
                "ticket_id": ticketId,
                "ticket_status": status
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('Ticket status updated:', response.data);
    } catch (error) {
        console.error('Error updating ticket status:', error.response?.data || error.message);
    }
}

// Example usage
// updateTicketStatus(jwtToken, 'ticket_123e4567-e89b-12d3-a456-426614174000', 'doing');
```

#### Update Ticket History

```javascript
const axios = require('axios');

async function updateTicketHistory(token, ticketId, update) {
    try {
        const response = await axios.post('domain/api/ticket/update', 
            {
                "ticket_id": ticketId,
                "ticket_return": update
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('Ticket history updated:', response.data);
    } catch (error) {
        console.error('Error updating ticket history:', error.response?.data || error.message);
    }
}

// Example usage
// updateTicketHistory(jwtToken, 'ticket_028492wsd88178', 'Purchasing routers');
```

#### List Tickets

```javascript
const axios = require('axios');

async function listTickets(token, author = null, status = null) {
    try {
        // Build query parameters
        let url = 'domain/api/ticket/fetch';
        const params = {};
        
        if (author) params.author = author;
        if (status) params.status = status;
        
        const response = await axios.get(url, {
            params: params,
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        console.log('Tickets:', response.data);
    } catch (error) {
        console.error('Error fetching tickets:', error.response?.data || error.message);
    }
}

// Example usage
// listTickets(jwtToken); // Get all tickets
// listTickets(jwtToken, 'johndoe@example.com'); // Get tickets by author
// listTickets(jwtToken, null, 'pending'); // Get pending tickets
```

#### Get Ticket Information

```javascript
const axios = require('axios');

async function getTicketInfo(token, ticketId) {
    try {
        const response = await axios.get('domain/api/ticket/info', {
            params: {
                ticket_id: ticketId
            },
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        console.log('Ticket information:', response.data);
    } catch (error) {
        console.error('Error fetching ticket info:', error.response?.data || error.message);
    }
}

// Example usage
// getTicketInfo(jwtToken, 'ticket_028492wsd88178');
```

#### Remove Ticket

```javascript
const axios = require('axios');

async function removeTicket(token, ticketId) {
    try {
        const response = await axios.post('domain/api/ticket/remove', 
            {
                "ticket_id": ticketId
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );
        
        console.log('Ticket removed:', response.data);
    } catch (error) {
        console.error('Error removing ticket:', error.response?.data || error.message);
    }
}

// Example usage
// removeTicket(jwtToken, 'ticket_123e4567-e89b-12d3-a456-426614174000');
``` 