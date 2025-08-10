package payroll

import (
	"employee-management/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// EmployeeService defines the interface for employee-related operations needed by the payroll service
type EmployeeService interface {
	ListEmployees(logger *logrus.Entry) ([]models.Employee, error)
}

// Service handles payroll-related business logic
type Service struct {
	repo            Repository
	employeeService EmployeeService
}

// NewService creates a new payroll service
func NewService(repo Repository, employeeService EmployeeService) *Service {
	return &Service{
		repo:            repo,
		employeeService: employeeService,
	}
}

// --- Salary Component ---

func (s *Service) CreateSalaryComponent(logger *logrus.Entry, data *models.SalaryComponentCreate) (*models.SalaryComponent, error) {
	return s.repo.CreateSalaryComponent(logger, data)
}

func (s *Service) GetSalaryComponentByID(logger *logrus.Entry, id uuid.UUID) (*models.SalaryComponent, error) {
	return s.repo.GetSalaryComponentByID(logger, id)
}

func (s *Service) ListSalaryComponents(logger *logrus.Entry) ([]models.SalaryComponent, error) {
	return s.repo.ListSalaryComponents(logger)
}

func (s *Service) UpdateSalaryComponent(logger *logrus.Entry, id uuid.UUID, data *models.SalaryComponentUpdate) (*models.SalaryComponent, error) {
	return s.repo.UpdateSalaryComponent(logger, id, data)
}

func (s *Service) DeleteSalaryComponent(logger *logrus.Entry, id uuid.UUID) error {
	return s.repo.DeleteSalaryComponent(logger, id)
}

// --- Employee Salary ---

func (s *Service) CreateEmployeeSalary(logger *logrus.Entry, data *models.EmployeeSalaryCreate) (*models.EmployeeSalary, error) {
	return s.repo.CreateEmployeeSalary(logger, data)
}

func (s *Service) GetEmployeeSalaries(logger *logrus.Entry, employeeID uuid.UUID) ([]models.EmployeeSalary, error) {
	return s.repo.GetEmployeeSalariesByEmployeeID(logger, employeeID)
}

func (s *Service) GetEmployeeSalary(logger *logrus.Entry, id uuid.UUID) (*models.EmployeeSalary, error) {
	return s.repo.GetEmployeeSalary(logger, id)
}

func (s *Service) UpdateEmployeeSalary(logger *logrus.Entry, id uuid.UUID, data *models.EmployeeSalaryUpdate) (*models.EmployeeSalary, error) {
	return s.repo.UpdateEmployeeSalary(logger, id, data)
}

func (s *Service) DeleteEmployeeSalary(logger *logrus.Entry, id uuid.UUID) error {
	return s.repo.DeleteEmployeeSalary(logger, id)
}

// --- Tax Bracket ---

func (s *Service) CreateTaxBracket(logger *logrus.Entry, data *models.TaxBracketCreate) (*models.TaxBracket, error) {
	return s.repo.CreateTaxBracket(logger, data)
}

func (s *Service) GetTaxBracket(logger *logrus.Entry, id uuid.UUID) (*models.TaxBracket, error) {
	return s.repo.GetTaxBracketByID(logger, id)
}

func (s *Service) GetTaxBrackets(logger *logrus.Entry, country string, year int) ([]models.TaxBracket, error) {
	return s.repo.GetTaxBrackets(logger, country, year)
}

func (s *Service) UpdateTaxBracket(logger *logrus.Entry, id uuid.UUID, data *models.TaxBracketUpdate) (*models.TaxBracket, error) {
	return s.repo.UpdateTaxBracket(logger, id, data)
}

func (s *Service) DeleteTaxBracket(logger *logrus.Entry, id uuid.UUID) error {
	return s.repo.DeleteTaxBracket(logger, id)
}

// --- Payroll Calculation ---

// CalculatePayrollInput represents the input for calculating payroll
type CalculatePayrollInput struct {
	PayPeriodStart time.Time `json:"pay_period_start" validate:"required"`
	PayPeriodEnd   time.Time `json:"pay_period_end"   validate:"required"`
	PaymentDate    time.Time `json:"payment_date"     validate:"required"`
	Country        string    `json:"country"          validate:"required"` // e.g., "USA"
}

// CalculatePayroll orchestrates the payroll calculation process
func (s *Service) CalculatePayroll(logger *logrus.Entry, input *CalculatePayrollInput) (*models.Payroll, error) {
	// 1. Create a new payroll record
	payroll, err := s.repo.CreatePayroll(logger, &models.PayrollCreate{
		PayPeriodStart: input.PayPeriodStart,
		PayPeriodEnd:   input.PayPeriodEnd,
		PaymentDate:    input.PaymentDate,
	})
	if err != nil {
		return nil, err
	}

	// 2. Get all employees
	employees, err := s.employeeService.ListEmployees(logger)
	if err != nil {
		return nil, err
	}

	// 3. Get tax brackets for the given country and year
	taxBrackets, err := s.repo.GetTaxBrackets(logger, input.Country, input.PayPeriodStart.Year())
	if err != nil {
		return nil, err
	}

	var totalGrossPay, totalDeductions, totalNetPay float64

	// 4. Iterate over each employee and calculate their payroll
	for _, employee := range employees {
		// Get employee's salary components
		salaries, err := s.repo.GetEmployeeSalariesByEmployeeID(logger, employee.ID)
		if err != nil {
			logger.WithError(err).WithField("employeeID", employee.ID).Error("Failed to get employee salaries")
			continue
		}

		// Calculate gross pay and deductions
		var grossPay, deductions float64
		var taxableEarnings float64
		for _, salary := range salaries {
			comp, err := s.repo.GetSalaryComponentByID(logger, salary.SalaryComponentID)
			if err != nil {
				continue
			}
			if comp.Type == "earning" {
				grossPay += salary.Amount
				if comp.IsTaxable {
					taxableEarnings += salary.Amount
				}
			} else if comp.Type == "deduction" {
				deductions += salary.Amount
			}
		}

		// Calculate tax
		taxAmount := s.calculateTax(taxableEarnings, taxBrackets)
		totalDeductionsForEmployee := deductions + taxAmount
		netPay := grossPay - totalDeductionsForEmployee

		// Create payroll detail record
		_, err = s.repo.CreatePayrollDetail(logger, &models.PayrollDetailCreate{
			PayrollID:       payroll.ID,
			EmployeeID:      employee.ID,
			GrossPay:        grossPay,
			TaxAmount:       taxAmount,
			OtherDeductions: deductions,
			NetPay:          netPay,
		})
		if err != nil {
			logger.WithError(err).Error("Failed to create payroll detail")
			continue
		}

		// Update totals
		totalGrossPay += grossPay
		totalDeductions += totalDeductionsForEmployee
		totalNetPay += netPay
	}

	// 5. Update the main payroll record with the calculated totals
	updatedPayroll, err := s.repo.UpdatePayroll(logger, payroll.ID, &models.PayrollUpdate{
		Status:          "calculated",
		TotalGrossPay:   totalGrossPay,
		TotalDeductions: totalDeductions,
		TotalNetPay:     totalNetPay,
	})
	if err != nil {
		return nil, err
	}

	return updatedPayroll, nil
}

// calculateTax calculates the tax amount based on taxable earnings and tax brackets
func (s *Service) calculateTax(taxableEarnings float64, brackets []models.TaxBracket) float64 {
	var tax float64
	for _, bracket := range brackets {
		if taxableEarnings > bracket.BracketMin {
			taxableAmountInBracket := taxableEarnings
			if taxableEarnings > bracket.BracketMax && bracket.BracketMax > 0 {
				taxableAmountInBracket = bracket.BracketMax
			}
			tax += (taxableAmountInBracket - bracket.BracketMin) * (bracket.TaxRate / 100)
		}
	}
	return tax
}

// --- Payroll Management ---

func (s *Service) GetPayrollByID(logger *logrus.Entry, id uuid.UUID) (*models.Payroll, error) {
	return s.repo.GetPayrollByID(logger, id)
}

func (s *Service) ListPayrolls(logger *logrus.Entry) ([]models.Payroll, error) {
	return s.repo.ListPayrolls(logger)
}

func (s *Service) GetPayrollDetails(logger *logrus.Entry, payrollID uuid.UUID) ([]models.PayrollDetail, error) {
	return s.repo.GetPayrollDetailsByPayrollID(logger, payrollID)
}

func (s *Service) ApprovePayroll(logger *logrus.Entry, id uuid.UUID) (*models.Payroll, error) {
	payroll, err := s.repo.GetPayrollByID(logger, id)
	if err != nil {
		return nil, err
	}
	if payroll.Status != "calculated" {
		return nil, errors.New("payroll must be in 'calculated' state to be approved")
	}
	return s.repo.UpdatePayroll(logger, id, &models.PayrollUpdate{Status: "approved"})
}

func (s *Service) ProcessPayroll(logger *logrus.Entry, id uuid.UUID) (*models.Payroll, error) {
	payroll, err := s.repo.GetPayrollByID(logger, id)
	if err != nil {
		return nil, err
	}
	if payroll.Status != "approved" {
		return nil, errors.New("payroll must be in 'approved' state to be processed")
	}

	details, err := s.repo.GetPayrollDetailsByPayrollID(logger, id)
	if err != nil {
		return nil, err
	}

	for _, detail := range details {
		// Create payslip
		_, err := s.repo.CreatePayslip(logger, &models.PayslipCreate{
			EmployeeID:     detail.EmployeeID,
			PayrollID:      payroll.ID,
			PayPeriodStart: payroll.PayPeriodStart,
			PayPeriodEnd:   payroll.PayPeriodEnd,
			GrossPay:       detail.GrossPay,
			TaxAmount:      detail.TaxAmount,
			Deductions:     map[string]float64{"other": detail.OtherDeductions}, // Simplified
			NetPay:         detail.NetPay,
		})
		if err != nil {
			logger.WithError(err).Error("Failed to create payslip")
		}
	}

	return s.repo.UpdatePayroll(logger, id, &models.PayrollUpdate{Status: "processed"})
}

func (s *Service) GetPayslip(logger *logrus.Entry, id uuid.UUID) (*models.Payslip, error) {
	return s.repo.GetPayslip(logger, id)
}
