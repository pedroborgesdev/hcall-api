# HCall API

A high-performance, secure API for comprehensive support ticket management with robust authentication. Built with Go and the Gin framework, designed for scalability and enterprise-grade security.

## Features

- 🔐 **JWT-based authentication** with secure token handling
- 👥 **Granular role-based access control** (User, Admin, Master)
- 🎫 **Full ticket lifecycle management** from creation to resolution
- 📸 **Secure image handling** with base64 encoding support
- 📝 **Comprehensive audit trails** with ticket history tracking
- 🔄 **Real-time status updates** and notification system
- ⚡ **High performance** with Go's concurrency model

## API Documentation

Access our comprehensive API documentation:

- [Interactive Online Documentation](https://pedroborgesdev.github.io/hcall-api)
- [Local Documentation](DOCUMENTATION.md)

## Tech Stack

| Component       | Technology                          |
|-----------------|-------------------------------------|
| Language        | Go 1.16+                           |
| Framework       | Gin                                 |
| Database        | PostgreSQL                          |
| Authentication  | JWT (JSON Web Tokens)               |
| ORM             | GORM                                |

## NEWS

### [Unreleased]
- Workers are added (for remove tickets after dates)
- Added a new endpoint (ticket/count) to count tickets
- Updated database schema

*Check back regularly for updates on new features and improvements.*

## Prerequisites

Before installation, ensure you have:

- Go 1.16 or later
- PostgreSQL 12+ server
- Basic understanding of REST APIs
- Environment configuration access

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
# Configure your environment variables in .env
```

### 4. Start the application
```bash
go run .
```
## Configuration

The application requires the following environment variables in your `.env` file:

```ini
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hcall
DB_SSLMODE=disable

# Password Policy
USERNAME_MIN_CHAR=6
PASSWORD_MIN_CHAR=8
PASSWORD_SPECIAL=True
PASSWORD_DIGITS=True
PASSWORD_UPPERCASE=True
PASSWORD_LOWERCASE=True

# Ticket Worker Settings
WORKER_TICKET_LOOPTIME=20
WORKER_TICKET_REMOVE_AFTER=10
WORKER_TICKET_REMOVE_STATUS=conclued

# Application Settings
PORT=8080
JWT_SECRET=mysecretkeyonhere
JWT_EXPIRATION_HOURS=24
```

**Security Note:** Always keep your `.env` file secure and never commit it to version control. The JWT_SECRET should be a strong, randomly generated string in production environments.

## API Endpoints

### Authentication
| Method | Endpoint                | Description                     |
|--------|-------------------------|---------------------------------|
| POST   | /api/auth/register      | Register new user               |
| POST   | /api/auth/enter         | User login                      |
| POST   | /api/master/create      | Create master user (privileged) |
| POST   | /api/master/delete      | Delete master user              |

### User Management
| Method | Endpoint                | Description                     |
|--------|-------------------------|---------------------------------|
| GET    | /api/user/fetch         | Retrieve user information       |
| POST   | /api/user/create        | Create new user                 |
| POST   | /api/user/delete        | Remove user                     |

### Ticket Operations
| Method | Endpoint                | Description                     |
|--------|-------------------------|---------------------------------|
| POST   | /api/ticket/create      | Create new support ticket       |
| GET    | /api/ticket/fetch       | List tickets                    |
| GET    | /api/ticket/count       | Count of tickets                |
| GET    | /api/ticket/info        | Get ticket details              |
| POST   | /api/ticket/edit        | Update ticket status            |
| POST   | /api/ticket/update      | Update ticket history           |
| POST   | /api/ticket/remove      | Delete ticket                   |

## Authorization Matrix

| Role  | User Management | Ticket Access | Ticket Creation | Admin Functions | Master Functions |
|-------|-----------------|---------------|------------------|-----------------|------------------|
| User  | Self-only       | Own tickets   | ✓                | ✗               | ✗                |
| Admin | Full            | All tickets   | ✗                | ✓               | ✗                |
| Master| Full            | All tickets   | ✓                | ✓               | ✓                |

## Security Features

- **End-to-end encryption**: JWT tokens with strong signing
- **Password security**: bcrypt hashing
- **RBAC implementation**: Strict role validation
- **CORS policies**: Strict origin validation

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for full details.

## Support & Contribution

For support requests:
- Open an issue in our [GitHub repository](https://github.com/pedroborgesdev/hcall-api/issues)
- Contact the maintainers directly

We welcome contributions! Please follow our contribution guidelines.

## Acknowledgments

- [Gin Framework](https://gin-gonic.com/) for high-performance routing
- [JWT-Go](https://github.com/golang-jwt/jwt) for secure authentication
- [GORM](https://gorm.io/) for database operations
- PostgreSQL for reliable data storage