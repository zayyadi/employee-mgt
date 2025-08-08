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

## API Documentation
API documentation is available at `/api/docs` when the server is running.

## License
MIT
