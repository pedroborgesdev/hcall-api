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
- [Online Documentation](https://your-domain.com/docs)
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

## Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Your Name - [@your_twitter](https://twitter.com/your_twitter) - email@example.com

Project Link: [https://github.com/your-username/api-tickets](https://github.com/your-username/api-tickets) 