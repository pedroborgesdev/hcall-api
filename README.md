# HCall API

A robust and secure API for managing support tickets and user authentication. Built with Go and Gin framework.

## Features

- 🔐 JWT-based authentication
- 👥 Role-based access control (User, Admin, Master)
- 🎫 Ticket management system
- 📸 Image upload support (base64 encoded)
- 📝 Ticket history tracking
- 🔄 Status updates and notifications

## Api Documentation
Complete API documentation is available at:
- [Online Documentation](https://pedroborgesdev.github.io/hcall-api)
- [documentation.md](documentation.md)

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT

## Prerequisites

- Go 1.16 or higher
- PostgreSQL

## Installation

1. Clone the repository:
```bash
git clone https://github.com/pedroborgesdev/hcall-api.git
cd hcall
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the application:
```bash
go run .
```

## Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/enter` - Login user
- `POST /api/master/create` - Create master user
- `POST /api/master/delete` - Delete master user

### Users
- `GET /api/user/fetch` - Get user information
- `POST /api/user/create` - Create new user
- `POST /api/user/delete` - Remove user

### Tickets
- `POST /api/ticket/create` - Create new ticket
- `GET /api/ticket/fetch` - List tickets
- `GET /api/ticket/info` - Get ticket details
- `POST /api/ticket/edit` - Update ticket status
- `POST /api/ticket/update` - Update ticket history
- `POST /api/ticket/remove` - Remove ticket

## User Roles and Permissions

### User
- Create and remove their own tickets
- Access authentication endpoints

### Admin
- All user permissions
- Manage other users
- View and update all tickets
- Cannot create tickets

### Master
- Full system access
- Manage all users and tickets
- Create users with any role

## Security

- All endpoints (except authentication) require JWT token
- Passwords are hashed using bcrypt
- Role-based access control
- Input validation and sanitization
- Rate limiting support

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please open an issue in the GitHub repository or contact the maintainers.

## Acknowledgments

- [Gin Framework](https://gin-gonic.com/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [GORM](https://gorm.io/) 