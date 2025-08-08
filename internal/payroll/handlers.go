package payroll

import (
	"employee-management/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler handles HTTP requests for payroll
type Handler struct {
	service *Service
}

// NewHandler creates a new payroll handler
func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// --- Salary Component Handlers ---

func (h *Handler) CreateSalaryComponent(c *gin.Context) {
	var input models.SalaryComponentCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comp, err := h.service.CreateSalaryComponent(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create salary component"})
		return
	}
	c.JSON(http.StatusCreated, comp)
}

func (h *Handler) GetSalaryComponent(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	comp, err := h.service.GetSalaryComponentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Salary component not found"})
		return
	}
	c.JSON(http.StatusOK, comp)
}

func (h *Handler) ListSalaryComponents(c *gin.Context) {
	comps, err := h.service.ListSalaryComponents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list salary components"})
		return
	}
	c.JSON(http.StatusOK, comps)
}

// --- Employee Salary Handlers ---

func (h *Handler) CreateEmployeeSalary(c *gin.Context) {
	var input models.EmployeeSalaryCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	salary, err := h.service.CreateEmployeeSalary(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee salary"})
		return
	}
	c.JSON(http.StatusCreated, salary)
}

func (h *Handler) GetEmployeeSalaries(c *gin.Context) {
	employeeID, _ := uuid.Parse(c.Param("employeeId"))
	salaries, err := h.service.GetEmployeeSalaries(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get employee salaries"})
		return
	}
	c.JSON(http.StatusOK, salaries)
}

// --- Tax Bracket Handlers ---

func (h *Handler) CreateTaxBracket(c *gin.Context) {
	var input models.TaxBracketCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bracket, err := h.service.CreateTaxBracket(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tax bracket"})
		return
	}
	c.JSON(http.StatusCreated, bracket)
}

func (h *Handler) GetTaxBrackets(c *gin.Context) {
	// In a real app, you'd get country and year from query params
	brackets, err := h.service.GetTaxBrackets("USA", 2024)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tax brackets"})
		return
	}
	c.JSON(http.StatusOK, brackets)
}

// --- Payroll Handlers ---

func (h *Handler) CalculatePayroll(c *gin.Context) {
	var input CalculatePayrollInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payroll, err := h.service.CalculatePayroll(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate payroll", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payroll)
}

func (h *Handler) ListPayrolls(c *gin.Context) {
	payrolls, err := h.service.ListPayrolls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list payrolls"})
		return
	}
	c.JSON(http.StatusOK, payrolls)
}

func (h *Handler) GetPayroll(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	payroll, err := h.service.GetPayrollByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payroll not found"})
		return
	}

	details, err := h.service.GetPayrollDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get payroll details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payroll": payroll,
		"details": details,
	})
}

func (h *Handler) ApprovePayroll(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	payroll, err := h.service.ApprovePayroll(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve payroll", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payroll)
}

func (h *Handler) ProcessPayroll(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	payroll, err := h.service.ProcessPayroll(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payroll", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payroll)
}

// --- Payslip Handlers ---

func (h *Handler) GetPayslip(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	payslip, err := h.service.GetPayslip(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payslip not found"})
		return
	}
	c.JSON(http.StatusOK, payslip)
}
