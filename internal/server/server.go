package server

import (
	"employee-management/internal/attendance"
	"employee-management/internal/auth"
	"employee-management/internal/database"
	"employee-management/internal/department"
	"employee-management/internal/employee"
	"employee-management/internal/leave"
	"employee-management/internal/middleware"
	"employee-management/internal/payroll"
	"employee-management/internal/position"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server represents the HTTP server
type Server struct {
	router            *gin.Engine
	db                *database.DB
	authHandler       *auth.Handler
	employeeHandler   *employee.Handler
	departmentHandler *department.Handler
	positionHandler   *position.Handler
	attendanceHandler *attendance.Handler
	leaveHandler      *leave.Handler
	payrollHandler    *payroll.Handler
}

// NewServer creates a new server instance
func NewServer(db *database.DB) *Server {
	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode)

	// Create router
	router := gin.New()

	// Add middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// Add custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Add custom validators here if needed
		_ = v
	}

	// Initialize auth service and handler
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret" // Fallback for development
	}

	authService := auth.NewService(db, jwtSecret)
	authHandler := auth.NewHandler(authService)

	employeeRepo := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepo)
	employeeHandler := employee.NewHandler(employeeService)

	departmentService := department.NewService(db)
	departmentHandler := department.NewHandler(departmentService)

	positionService := position.NewService(db)
	positionHandler := position.NewHandler(positionService)

	attendanceService := attendance.NewService(db)
	attendanceHandler := attendance.NewHandler(attendanceService)

	leaveRepo := leave.NewRepository(db)
	leaveService := leave.NewService(leaveRepo)
	leaveHandler := leave.NewHandler(leaveService)

	payrollRepo := payroll.NewRepository(db)
	payrollService := payroll.NewService(payrollRepo, employeeService)
	payrollHandler := payroll.NewHandler(payrollService)

	return &Server{
		router:            router,
		db:                db,
		authHandler:       authHandler,
		employeeHandler:   employeeHandler,
		departmentHandler: departmentHandler,
		positionHandler:   positionHandler,
		attendanceHandler: attendanceHandler,
		leaveHandler:      leaveHandler,
		payrollHandler:    payrollHandler,
	}
}

// setupRoutes sets up all the routes for the application
func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", s.healthCheck)

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", s.login)
			auth.POST("/logout", s.logout)
			auth.POST("/register", s.register)
			auth.POST("/forgot-password", s.forgotPassword)
			auth.POST("/reset-password", s.resetPassword)
			auth.POST("/refresh", s.refreshToken)
		}

		// Employee routes
		employees := v1.Group("/employees")
		{
			employees.GET("/", s.listEmployees)
			employees.GET("/:id", s.getEmployee)
			employees.POST("/", s.createEmployee)
			employees.PUT("/:id", s.updateEmployee)
			employees.DELETE("/:id", s.deleteEmployee)
			employees.GET("/search", s.searchEmployees)
		}

		// Department routes
		departments := v1.Group("/departments")
		{
			departments.GET("/", s.listDepartments)
			departments.GET("/:id", s.getDepartment)
			departments.POST("/", s.createDepartment)
			departments.PUT("/:id", s.updateDepartment)
			departments.DELETE("/:id", s.deleteDepartment)
		}

		// Position routes
		positions := v1.Group("/positions")
		{
			positions.GET("/", s.listPositions)
			positions.GET("/:id", s.getPosition)
			positions.POST("/", s.createPosition)
			positions.PUT("/:id", s.updatePosition)
			positions.DELETE("/:id", s.deletePosition)
		}

		// Attendance routes
		attendance := v1.Group("/attendance")
		{
			attendance.GET("/", s.listAttendance)
			attendance.GET("/:id", s.getAttendance)
			attendance.POST("/check-in", s.checkIn)
			attendance.POST("/check-out", s.checkOut)
			attendance.POST("/", s.createAttendance)
			attendance.PUT("/:id", s.updateAttendance)
		}

		// Leave routes
		leave := v1.Group("/leave")
		{
			// Leave Types
			leaveTypes := leave.Group("/types")
			{
				leaveTypes.GET("/", s.listLeaveTypes)
				leaveTypes.POST("/", s.createLeaveType)
				leaveTypes.GET("/:id", s.getLeaveType)
				leaveTypes.PUT("/:id", s.updateLeaveType)
				leaveTypes.DELETE("/:id", s.deleteLeaveType)
			}

			// Leave Requests
			leaveRequests := leave.Group("/requests")
			{
				leaveRequests.GET("/", s.listLeaveRequests)
				leaveRequests.POST("/", s.createLeaveRequest)
				leaveRequests.GET("/:id", s.getLeaveRequest)
				leaveRequests.PUT("/:id/approve", s.approveLeaveRequest)
				leaveRequests.PUT("/:id/reject", s.rejectLeaveRequest)
			}
		}

		// Payroll routes
		payrollRoutes := v1.Group("/payroll")
		{
			payrollRoutes.POST("/calculate", s.calculatePayroll)
			payrollRoutes.GET("/", s.listPayrolls)
			payrollRoutes.GET("/:id", s.getPayroll)
			payrollRoutes.POST("/:id/approve", s.approvePayroll)
			payrollRoutes.POST("/:id/process", s.processPayroll)

			// Salary Components
			components := payrollRoutes.Group("/components")
			{
				components.POST("/", s.createSalaryComponent)
				components.GET("/", s.listSalaryComponents)
				components.GET("/:id", s.getSalaryComponent)
			}

			// Employee Salaries
			employeeSalaries := payrollRoutes.Group("/employee-salaries")
			{
				employeeSalaries.POST("/", s.createEmployeeSalary)
				employeeSalaries.GET("/:employeeId", s.getEmployeeSalaries)
			}

			// Tax Brackets
			taxBrackets := payrollRoutes.Group("/tax-brackets")
			{
				taxBrackets.POST("/", s.createTaxBracket)
				taxBrackets.GET("/", s.getTaxBrackets)
			}
		}

		// Payslip routes
		payslips := v1.Group("/payslips")
		{
			payslips.GET("/:id", s.getPayslip)
		}

		// Report routes
		reports := v1.Group("/reports")
		{
			reports.GET("/", s.listReports)
			reports.GET("/:type", s.generateReport)
			reports.POST("/:type/export", s.exportReport)
		}

		// Notification routes
		notifications := v1.Group("/notifications")
		{
			notifications.GET("/", s.listNotifications)
			notifications.PUT("/:id/read", s.markNotificationAsRead)
			notifications.PUT("/read-all", s.markAllNotificationsAsRead)
		}
	}
}

// Run starts the HTTP server
func (s *Server) Run() error {
	// Setup routes
	s.setupRoutes()

	// Create HTTP server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start server
	return server.ListenAndServe()
}

// healthCheck handles health check requests
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// Auth handlers
func (s *Server) login(c *gin.Context)          { s.authHandler.Login(c) }
func (s *Server) logout(c *gin.Context)         { s.authHandler.Logout(c) }
func (s *Server) register(c *gin.Context)       { s.authHandler.Register(c) }
func (s *Server) forgotPassword(c *gin.Context) { s.authHandler.ForgotPassword(c) }
func (s *Server) resetPassword(c *gin.Context)  { s.authHandler.ResetPassword(c) }
func (s *Server) refreshToken(c *gin.Context)   { s.authHandler.RefreshToken(c) }
func (s *Server) listEmployees(c *gin.Context) {
	s.employeeHandler.ListEmployees(c)
}
func (s *Server) getEmployee(c *gin.Context) {
	s.employeeHandler.GetEmployee(c)
}
func (s *Server) createEmployee(c *gin.Context) {
	s.employeeHandler.CreateEmployee(c)
}
func (s *Server) updateEmployee(c *gin.Context) {
	s.employeeHandler.UpdateEmployee(c)
}
func (s *Server) deleteEmployee(c *gin.Context) {
	s.employeeHandler.DeleteEmployee(c)
}
func (s *Server) searchEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "search employees endpoint"})
}
func (s *Server) listDepartments(c *gin.Context) {
	s.departmentHandler.ListDepartments(c)
}
func (s *Server) getDepartment(c *gin.Context) {
	s.departmentHandler.GetDepartment(c)
}
func (s *Server) createDepartment(c *gin.Context) {
	s.departmentHandler.CreateDepartment(c)
}
func (s *Server) updateDepartment(c *gin.Context) {
	s.departmentHandler.UpdateDepartment(c)
}
func (s *Server) deleteDepartment(c *gin.Context) {
	s.departmentHandler.DeleteDepartment(c)
}
func (s *Server) listPositions(c *gin.Context) {
	s.positionHandler.ListPositions(c)
}
func (s *Server) getPosition(c *gin.Context) {
	s.positionHandler.GetPosition(c)
}
func (s *Server) createPosition(c *gin.Context) {
	s.positionHandler.CreatePosition(c)
}
func (s *Server) updatePosition(c *gin.Context) {
	s.positionHandler.UpdatePosition(c)
}
func (s *Server) deletePosition(c *gin.Context) {
	s.positionHandler.DeletePosition(c)
}
func (s *Server) listAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list attendance endpoint"})
}
func (s *Server) getAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get attendance endpoint"})
}
func (s *Server) checkIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "check in endpoint"})
}
func (s *Server) checkOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "check out endpoint"})
}
func (s *Server) createAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create attendance endpoint"})
}
func (s *Server) updateAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update attendance endpoint"})
}

// Leave handlers
func (s *Server) listLeaveTypes(c *gin.Context)      { s.leaveHandler.ListLeaveTypes(c) }
func (s *Server) createLeaveType(c *gin.Context)     { s.leaveHandler.CreateLeaveType(c) }
func (s *Server) getLeaveType(c *gin.Context)        { s.leaveHandler.GetLeaveType(c) }
func (s *Server) updateLeaveType(c *gin.Context)     { s.leaveHandler.UpdateLeaveType(c) }
func (s *Server) deleteLeaveType(c *gin.Context)     { s.leaveHandler.DeleteLeaveType(c) }
func (s *Server) listLeaveRequests(c *gin.Context)   { s.leaveHandler.ListLeaveRequests(c) }
func (s *Server) createLeaveRequest(c *gin.Context)  { s.leaveHandler.CreateLeaveRequest(c) }
func (s *Server) getLeaveRequest(c *gin.Context)     { s.leaveHandler.GetLeaveRequest(c) }
func (s *Server) approveLeaveRequest(c *gin.Context) { s.leaveHandler.ApproveLeaveRequest(c) }
func (s *Server) rejectLeaveRequest(c *gin.Context)  { s.leaveHandler.RejectLeaveRequest(c) }

// Payroll Handlers
func (s *Server) calculatePayroll(c *gin.Context)      { s.payrollHandler.CalculatePayroll(c) }
func (s *Server) listPayrolls(c *gin.Context)          { s.payrollHandler.ListPayrolls(c) }
func (s *Server) getPayroll(c *gin.Context)            { s.payrollHandler.GetPayroll(c) }
func (s *Server) approvePayroll(c *gin.Context)        { s.payrollHandler.ApprovePayroll(c) }
func (s *Server) processPayroll(c *gin.Context)        { s.payrollHandler.ProcessPayroll(c) }
func (s *Server) createSalaryComponent(c *gin.Context) { s.payrollHandler.CreateSalaryComponent(c) }
func (s *Server) listSalaryComponents(c *gin.Context)  { s.payrollHandler.ListSalaryComponents(c) }
func (s *Server) getSalaryComponent(c *gin.Context)    { s.payrollHandler.GetSalaryComponent(c) }
func (s *Server) createEmployeeSalary(c *gin.Context)  { s.payrollHandler.CreateEmployeeSalary(c) }
func (s *Server) getEmployeeSalaries(c *gin.Context)   { s.payrollHandler.GetEmployeeSalaries(c) }
func (s *Server) createTaxBracket(c *gin.Context)      { s.payrollHandler.CreateTaxBracket(c) }
func (s *Server) getTaxBrackets(c *gin.Context)        { s.payrollHandler.GetTaxBrackets(c) }
func (s *Server) getPayslip(c *gin.Context)            { s.payrollHandler.GetPayslip(c) }

func (s *Server) listReports(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list reports endpoint"})
}
func (s *Server) generateReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "generate report endpoint"})
}
func (s *Server) exportReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "export report endpoint"})
}
func (s *Server) listNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list notifications endpoint"})
}
func (s *Server) markNotificationAsRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "mark notification as read endpoint"})
}
func (s *Server) markAllNotificationsAsRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "mark all notifications as read endpoint"})
}
