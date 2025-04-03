# HCall API

A high-performance, secure RESTful API for enterprise-grade support ticket management with advanced authentication and role-based access control. Built with Go and the Gin framework, designed for reliability, scalability, and security.

## Key Features

- 🔒 **Enterprise-grade Security**
  - JWT-based authentication with configurable expiration
  - Secure password policies with customizable complexity requirements
  - Role-based access control with granular permissions

- 🎫 **Comprehensive Ticket Management**
  - Complete lifecycle management from creation to resolution
  - Rich media support with secure image handling (base64 encoding)
  - Advanced filtering by author, status, date, and keywords
  - Detailed ticket history with timestamped audit trails

- ⚙️ **System Architecture**
  - High-performance REST API built with Go and Gin
  - ACID-compliant transactions for data integrity
  - Background workers for automated maintenance tasks
  - Structured, clean code with separation of concerns

## Documentation

Access our comprehensive API documentation:

- [Interactive Web Documentation](https://pedroborgesdev.github.io/hcall-api)
- [Local Documentation](DOCUMENTATION.md)

## Technology Stack

| Component       | Technology                          | Description                                |
|-----------------|-------------------------------------|--------------------------------------------|
| Language        | Go 1.16+                           | High-performance, concurrent programming   |
| Framework       | Gin Web Framework                   | Lightweight HTTP router with middleware    |
| Database        | PostgreSQL 12+                      | Robust, ACID-compliant relational database |
| Authentication  | JWT (JSON Web Tokens)               | Secure, stateless authentication           |
| ORM             | GORM                                | Powerful ORM with migrations and hooks     |
| Workers         | Native Go routines                  | Background task processing                 |

## Recent Updates

### [Unreleased]
- **Background Workers**: Automated ticket cleanup based on status and age
- **Database Optimizations**: Schema improvements for better performance
- **ACID Transactions**: Enhanced data integrity across operations
- **New Logs**: More detailed and beautiful logs
- **Defense against DoS attack**: Using Rate Limit system by IP to block multiple requests
- **Fix CORS Problems**: CORS errors in requests made by the browser or libraries like axios have been fixed

*Check back regularly for updates on new features and improvements.*

- Note: This version aren't stable!

## Prerequisites

Before installation, ensure you have:

- Go 1.16 or later
- PostgreSQL 12+ server
- Git for version control
- Understanding of RESTful APIs and JWT authentication

## Installation & Setup

### 1. Clone the repository
```bash
git clone https://github.com/pedroborgesdev/hcall-api.git
cd hcall-api
cd api
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Configure environment
```bash
cp .env.example .env
# Edit .env file with your specific configuration
```

### 4. Start the application
```bash
go run .
```

## Environment Configuration

The application is highly configurable through environment variables:

```ini
# Server Configuration
PORT=8080
HOST=localhost

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hcall
DB_SSL_MODE=disable

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Rate Limiting
RATE_LIMIT_REQUESTS=60
RATE_LIMIT_WINDOW=5

# Database Configuration
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
DB_CONN_TIMEOUT=5

# Other Preferences
USERNAME_MIN_CHAR=6
PASSWORD_MIN_CHAR=8
PASSWORD_SPECIAL=True
PASSWORD_DIGITS=True
PASSWORD_UPPERCASE=True
PASSWORD_LOWERCASE=True

# Workers Configuration
WORKER_TICKET_LOOPTIME=24
WORKER_TICKET_REMOVE_AFTER=30
WORKER_TICKET_REMOVE_STATUS=conclued

# Debug Modes
DEBUG=true
GIN_MODE=false
```

**⚠️ Security Note:** Never commit your `.env` file to version control. In production, use a strong, randomly generated JWT secret key.

## API Overview

### Authentication & Master Management
| Method | Endpoint                | Description                     | Authorized Roles |
|--------|-------------------------|---------------------------------|------------------|
| POST   | /api/auth/register      | User self-registration          | Public           |
| POST   | /api/auth/enter         | User login and token issuance   | Public           |
| POST   | /api/master/create      | Create initial master user      | Public (once)    |
| POST   | /api/master/delete      | Delete master user              | Public (auth)    |

### User Management
| Method | Endpoint                | Description                     | Authorized Roles |
|--------|-------------------------|---------------------------------|------------------|
| GET    | /api/user/fetch         | Retrieve user(s) information    | Admin, Master    |
| POST   | /api/user/create        | Create new user                 | Admin, Master    |
| POST   | /api/user/delete        | Delete existing user            | Admin, Master    |

### Ticket Management
| Method | Endpoint                | Description                     | Authorized Roles |
|--------|-------------------------|---------------------------------|------------------|
| POST   | /api/ticket/create      | Create new support ticket       | User             |
| GET    | /api/ticket/fetch       | List and filter tickets         | Admin, Master    |
| GET    | /api/ticket/count       | Get ticket counts by status     | All authenticated |
| GET    | /api/ticket/info        | Get detailed ticket information | Admin, Master    |
| POST   | /api/ticket/edit        | Update ticket status            | Admin, Master    |
| POST   | /api/ticket/update      | Add entry to ticket history     | Admin, Master    |
| POST   | /api/ticket/remove      | Delete ticket                   | User*, Admin, Master |

\* Users can only delete their own tickets

## Role-Based Access Control

The API implements a comprehensive role-based access control system:

| Role    | Description                           | Capabilities                                              |
|---------|---------------------------------------|------------------------------------------------------------|
| User    | Standard users who create tickets     | Create tickets, view own tickets, view ticket counts       |
| Admin   | Support staff who manage tickets      | View all tickets, update tickets, manage users             |
| Master  | System administrators with full access | All admin capabilities plus system configuration           |

## Security Architecture

The HCall API implements multiple layers of security:

- **Authentication**: JWT tokens with secure signing and controlled expiration
- **Password Security**: Enforced complexity requirements and bcrypt hashing
- **Access Control**: Strict role validation for each API endpoint
- **Data Protection**: ACID-compliant transactions for critical operations
- **API Security**: Input validation and sanitization to prevent injection attacks
- **Audit Trails**: Comprehensive logging and history tracking

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/) for high-performance API routing
- [JWT-Go](https://github.com/golang-jwt/jwt) for secure authentication
- [GORM](https://gorm.io/) for robust database operations
- [PostgreSQL](https://www.postgresql.org/) for reliable data storage