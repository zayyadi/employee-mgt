# Employee Management System with Payroll

A modular monolith employee management system with integrated payroll functionality built with Go and Gin framework.

## Features
- Employee management (profiles, departments, positions)
- Attendance tracking
- Leave management
- Performance reviews
- Document management
- Payroll calculation and processing
- Tax management
- Payslip generation
- Reporting and analytics
- Role-based access control
- Two-factor authentication
- Web dashboard with responsive UI

## Technology Stack
- **Backend**: Go with Gin framework
- **Database**: PostgreSQL
- **Authentication**: JWT with 2FA support
- **Frontend**: HTML/CSS/JavaScript with Bootstrap
- **Documentation**: Swagger/OpenAPI

## Project Structure
```
employee-management/
├── cmd/
│   └── server/
├── internal/
│   ├── auth/
│   ├── employee/
│   ├── department/
│   ├── position/
│   ├── attendance/
│   ├── leave/
│   ├── performance/
│   ├── document/
│   ├── payroll/
│   ├── report/
│   ├── notification/
│   ├── database/
│   └── middleware/
├── pkg/
├── web/
│   ├── css/
│   ├── js/
│   ├── assets/
│   ├── components/
│   ├── pages/
│   ├── index.html
│   ├── login.html
│   └── employee-profile.html
├── configs/
├── migrations/
├── docs/
├── scripts/
└── tests/
```

## Getting Started
1. Install Go 1.19+
2. Install PostgreSQL
3. Clone the repository
4. Configure environment variables
5. Run database migrations
6. Start the server

## Web Dashboard
The application includes a complete web dashboard with:
- Admin dashboard with analytics
- Employee management interface
- Department and position management
- Attendance tracking
- Leave management system
- Payroll processing interface
- Reporting tools

To access the web dashboard:
1. Start the server
2. Navigate to `http://localhost:8080/web/login.html` to login
3. Use any username and password to access the dashboard (demo mode)

## API Documentation
API documentation is available at `/api/docs` when the server is running.

## License
MIT
