# Employee Management System with Payroll - Project Design

## Overview
This document outlines the complete design and architecture for a modular monolith employee management system with an integrated payroll system. The application will be built using Go with the Gin framework and PostgreSQL as the database.

## Technology Stack
- **Backend**: Go with Gin framework
- **Database**: PostgreSQL
- **Authentication**: JWT-based authentication with 2FA support
- **Frontend**: HTML/CSS/JavaScript with Bootstrap
- **API Documentation**: Swagger/OpenAPI

## Modular Architecture Design

### Core Modules
1. **Authentication Module**: User management, roles, login/logout, password reset, 2FA
2. **Employee Management Module**: Employee profiles, departments, positions, organizational chart
3. **Attendance Module**: Time tracking, check-in/check-out, attendance reports
4. **Leave Management Module**: Leave requests, approvals, balance tracking, leave policies
5. **Performance Module**: Performance reviews, goal tracking, feedback system
6. **Document Module**: Document storage, categorization, access control
7. **Payroll Module**: Salary calculation, tax management, payslip generation
8. **Report Module**: Analytics dashboard, export functionality
9. **Notification Module**: Email/SMS notifications for various events

### Module Structure
Each module follows this pattern:
```
module/
├── handlers/       # HTTP handlers
├── services/       # Business logic
├── repositories/   # Data access layer
├── models/         # Data structures
├── dto/            # Data transfer objects
├── middleware/     # Module-specific middleware
└── utils/          # Utility functions
```

## Database Schema Design

### Core Tables

#### 1. users
```sql
- id (UUID, PK)
- username (VARCHAR)
- email (VARCHAR, UNIQUE)
- password_hash (VARCHAR)
- role (VARCHAR)
- is_active (BOOLEAN)
- last_login (TIMESTAMP)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 2. employees
```sql
- id (UUID, PK)
- user_id (UUID, FK to users)
- employee_id (VARCHAR, UNIQUE)
- first_name (VARCHAR)
- last_name (VARCHAR)
- date_of_birth (DATE)
- gender (VARCHAR)
- marital_status (VARCHAR)
- phone_number (VARCHAR)
- email (VARCHAR)
- address (TEXT)
- emergency_contact_name (VARCHAR)
- emergency_contact_phone (VARCHAR)
- department_id (UUID, FK to departments)
- position_id (UUID, FK to positions)
- hire_date (DATE)
- employment_status (VARCHAR)
- manager_id (UUID, FK to employees)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 3. departments
```sql
- id (UUID, PK)
- name (VARCHAR)
- description (TEXT)
- manager_id (UUID, FK to employees)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 4. positions
```sql
- id (UUID, PK)
- title (VARCHAR)
- department_id (UUID, FK to departments)
- description (TEXT)
- requirements (TEXT)
- salary_range_min (DECIMAL)
- salary_range_max (DECIMAL)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 5. attendance
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- check_in_time (TIMESTAMP)
- check_out_time (TIMESTAMP)
- date (DATE)
- status (VARCHAR)
- notes (TEXT)
- created_at (TIMESTAMP)
```

#### 6. leave_types
```sql
- id (UUID, PK)
- name (VARCHAR)
- description (TEXT)
- max_days_per_year (INTEGER)
- is_accrued (BOOLEAN)
- created_at (TIMESTAMP)
```

#### 7. leave_requests
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- leave_type_id (UUID, FK to leave_types)
- start_date (DATE)
- end_date (DATE)
- reason (TEXT)
- status (VARCHAR)  -- pending, approved, rejected
- approved_by (UUID, FK to employees)
- approved_at (TIMESTAMP)
- created_at (TIMESTAMP)
```

#### 8. performance_reviews
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- reviewer_id (UUID, FK to employees)
- review_date (DATE)
- rating (INTEGER)
- comments (TEXT)
- goals (TEXT)
- created_at (TIMESTAMP)
```

#### 9. documents
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- name (VARCHAR)
- description (TEXT)
- file_path (VARCHAR)
- category (VARCHAR)
- uploaded_by (UUID, FK to employees)
- created_at (TIMESTAMP)
```

#### 10. salary_components
```sql
- id (UUID, PK)
- name (VARCHAR)
- type (VARCHAR)  -- earning, deduction
- is_taxable (BOOLEAN)
- is_recurring (BOOLEAN)
- created_at (TIMESTAMP)
```

#### 11. employee_salaries
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- salary_component_id (UUID, FK to salary_components)
- amount (DECIMAL)
- effective_date (DATE)
- end_date (DATE)
- created_at (TIMESTAMP)
```

#### 12. tax_brackets
```sql
- id (UUID, PK)
- country (VARCHAR)
- tax_year (INTEGER)
- bracket_min (DECIMAL)
- bracket_max (DECIMAL)
- tax_rate (DECIMAL)
- created_at (TIMESTAMP)
```

#### 13. payroll
```sql
- id (UUID, PK)
- pay_period_start (DATE)
- pay_period_end (DATE)
- payment_date (DATE)
- status (VARCHAR)  -- draft, calculated, approved, processed
- total_gross_pay (DECIMAL)
- total_deductions (DECIMAL)
- total_net_pay (DECIMAL)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 14. payroll_details
```sql
- id (UUID, PK)
- payroll_id (UUID, FK to payroll)
- employee_id (UUID, FK to employees)
- gross_pay (DECIMAL)
- tax_amount (DECIMAL)
- other_deductions (DECIMAL)
- net_pay (DECIMAL)
- created_at (TIMESTAMP)
```

#### 15. payslips
```sql
- id (UUID, PK)
- employee_id (UUID, FK to employees)
- payroll_id (UUID, FK to payroll)
- pay_period_start (DATE)
- pay_period_end (DATE)
- gross_pay (DECIMAL)
- tax_amount (DECIMAL)
- deductions (JSON)
- net_pay (DECIMAL)
- file_path (VARCHAR)
- created_at (TIMESTAMP)
```

#### 16. reports
```sql
- id (UUID, PK)
- name (VARCHAR)
- type (VARCHAR)
- description (TEXT)
- query (TEXT)
- created_at (TIMESTAMP)
```

#### 17. notifications
```sql
- id (UUID, PK)
- user_id (UUID, FK to users)
- title (VARCHAR)
- message (TEXT)
- is_read (BOOLEAN)
- created_at (TIMESTAMP)
```

## API Design

### Authentication Endpoints
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `POST /api/v1/auth/register` - User registration (admin only)
- `POST /api/v1/auth/forgot-password` - Password reset request
- `POST /api/v1/auth/reset-password` - Password reset
- `POST /api/v1/auth/refresh` - Token refresh
- `POST /api/v1/auth/2fa/setup` - Setup 2FA
- `POST /api/v1/auth/2fa/verify` - Verify 2FA code

### Employee Management Endpoints
- `GET /api/v1/employees` - List all employees
- `GET /api/v1/employees/{id}` - Get employee details
- `POST /api/v1/employees` - Create new employee
- `PUT /api/v1/employees/{id}` - Update employee
- `DELETE /api/v1/employees/{id}` - Delete employee
- `GET /api/v1/employees/search` - Search employees

### Department Endpoints
- `GET /api/v1/departments` - List all departments
- `GET /api/v1/departments/{id}` - Get department details
- `POST /api/v1/departments` - Create department
- `PUT /api/v1/departments/{id}` - Update department
- `DELETE /api/v1/departments/{id}` - Delete department

### Position Endpoints
- `GET /api/v1/positions` - List all positions
- `GET /api/v1/positions/{id}` - Get position details
- `POST /api/v1/positions` - Create position
- `PUT /api/v1/positions/{id}` - Update position
- `DELETE /api/v1/positions/{id}` - Delete position

### Attendance Endpoints
- `GET /api/v1/attendance` - List attendance records
- `GET /api/v1/attendance/{id}` - Get attendance record
- `POST /api/v1/attendance/check-in` - Employee check-in
- `POST /api/v1/attendance/check-out` - Employee check-out
- `POST /api/v1/attendance` - Create attendance record (admin)
- `PUT /api/v1/attendance/{id}` - Update attendance record

### Leave Management Endpoints
- `GET /api/v1/leave/types` - List leave types
- `GET /api/v1/leave/requests` - List leave requests
- `GET /api/v1/leave/requests/{id}` - Get leave request
- `POST /api/v1/leave/requests` - Create leave request
- `PUT /api/v1/leave/requests/{id}/approve` - Approve leave request
- `PUT /api/v1/leave/requests/{id}/reject` - Reject leave request

### Performance Endpoints
- `GET /api/v1/performance/reviews` - List performance reviews
- `GET /api/v1/performance/reviews/{id}` - Get performance review
- `POST /api/v1/performance/reviews` - Create performance review
- `PUT /api/v1/performance/reviews/{id}` - Update performance review

### Document Endpoints
- `GET /api/v1/documents` - List documents
- `GET /api/v1/documents/{id}` - Get document
- `POST /api/v1/documents` - Upload document
- `DELETE /api/v1/documents/{id}` - Delete document

### Payroll Endpoints
- `GET /api/v1/payroll` - List payroll records
- `POST /api/v1/payroll/calculate` - Calculate payroll for period
- `GET /api/v1/payroll/{id}` - Get payroll details
- `POST /api/v1/payroll/{id}/approve` - Approve payroll
- `POST /api/v1/payroll/{id}/process` - Process payroll
- `GET /api/v1/payslips/{id}` - Get payslip
- `POST /api/v1/payslips/{id}/send` - Send payslip to employee

### Report Endpoints
- `GET /api/v1/reports` - List available reports
- `GET /api/v1/reports/{type}` - Generate specific report
- `POST /api/v1/reports/{type}/export` - Export report in specific format

### Notification Endpoints
- `GET /api/v1/notifications` - List notifications
- `PUT /api/v1/notifications/{id}/read` - Mark notification as read
- `PUT /api/v1/notifications/read-all` - Mark all notifications as read

## Frontend Structure

### Admin Dashboard
- Employee management interface
- Department and position management
- Attendance monitoring
- Leave approval system
- Performance review management
- Payroll processing interface
- Report generation tools
- User management

### Employee Portal
- Personal profile management
- Attendance tracking
- Leave request submission
- Performance review viewing
- Document access
- Payslip viewing
- Notification center

## Security Features
- Role-based access control (RBAC)
- JWT token authentication
- Two-factor authentication
- Password encryption with bcrypt
- Input validation and sanitization
- Rate limiting
- CORS protection
- SQL injection prevention
- Cross-site scripting (XSS) protection

## Deployment Considerations
- Docker containerization
- Environment-based configuration
- Database migrations
- Health check endpoints
- Logging and monitoring
- Backup and recovery procedures
- SSL/TLS encryption

## Development Phases

### Phase 1: Project Setup and Authentication
- Project structure setup
- Database configuration
- Authentication module implementation
- User management
- Basic employee CRUD operations

### Phase 2: Core Employee Management
- Department and position management
- Attendance tracking system
- Organizational chart implementation

### Phase 3: Leave and Performance Management
- Leave request system
- Approval workflows
- Performance review module

### Phase 4: Document Management
- File upload and storage
- Document categorization
- Access control implementation

### Phase 5: Payroll Engine
- Salary calculation logic
- Tax computation system
- Salary component management

### Phase 6: Payslip and Banking Integration
- Payslip generation
- Bank transfer integration
- Payment processing

### Phase 7: Reporting System
- Analytics dashboard
- Report generation engine
- Export functionality (PDF, Excel)

### Phase 8: Notification and Communication
- Email/SMS notification system
- Real-time updates
- Communication channels

### Phase 9: Frontend Implementation
- Admin dashboard development
- Employee portal development
- Responsive design implementation

### Phase 10: Testing and Deployment
- Unit testing
- Integration testing
- Security audit
- Performance optimization
- Documentation
- Deployment setup

## Testing Strategy
- Unit tests for each module
- Integration tests for API endpoints
- End-to-end tests for critical workflows
- Security testing
- Performance testing
- Load testing

## Monitoring and Logging
- Application logging
- Error tracking
- Performance monitoring
- Database query optimization
- API response time tracking

This comprehensive design document provides a roadmap for developing a robust, scalable employee management system with an integrated payroll module using a modular monolith architecture.
