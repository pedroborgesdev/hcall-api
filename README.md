# Ticket Management API

A RESTful API for technical support ticket management, developed with Node.js and Express. This API allows you to create, manage, and track support tickets, as well as manage system users.

## Features

- 🔐 Secure JWT Authentication
- 👥 Complete User Management
- 🎫 Ticket Creation and Tracking
- 📊 Status and History Updates
- 🔍 Advanced Search and Filters
- 📱 RESTful and Responsive API

## Documentation

Complete API documentation is available at:
- [Online Documentation](https://pedroborgesdev.github.io/hcall-api)
- [DOCUMENTATION.md](DOCUMENTATION.md)

## Main Endpoints

### Authentication
- `POST /auth/enter` - User Login
- `POST /auth/register` - New User Registration

### Users
- `GET /user/fetch` - List Users
- `POST /user/create` - Create User
- `POST /user/delete` - Remove User

### Tickets
- `POST /ticket/create` - Create Ticket
- `POST /ticket/edit` - Update Status
- `POST /ticket/update` - Update History
- `GET /ticket/fetch` - List Tickets
- `GET /ticket/info` - Ticket Details
- `POST /ticket/remove` - Remove Ticket

## Usage Example

```javascript
// Login
axios.post('domain/api/auth/enter', {
  username: "John Doe",
  email: "johndoe@example.com",
  password: "Password123"
})
.then(response => {
  const token = response.data.token;
  console.log('Login successful:', response.data);
})
.catch(error => {
  console.error('Login error:', error.response.data);
});
```

## Security

- JWT Authentication
- Encrypted Passwords
- Data Validation
- Protection Against Common Attacks

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
