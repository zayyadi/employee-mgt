package payroll

import (
	"employee-management/internal/database"
	"employee-management/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Repository defines the interface for payroll data operations
type Repository interface {
	// Salary Component methods
	CreateSalaryComponent(logger *logrus.Entry, data *models.SalaryComponentCreate) (*models.SalaryComponent, error)
	GetSalaryComponentByID(logger *logrus.Entry, id uuid.UUID) (*models.SalaryComponent, error)
	ListSalaryComponents(logger *logrus.Entry) ([]models.SalaryComponent, error)
	UpdateSalaryComponent(logger *logrus.Entry, id uuid.UUID, data *models.SalaryComponentUpdate) (*models.SalaryComponent, error)
	DeleteSalaryComponent(logger *logrus.Entry, id uuid.UUID) error

	// Employee Salary methods
	CreateEmployeeSalary(logger *logrus.Entry, data *models.EmployeeSalaryCreate) (*models.EmployeeSalary, error)
	GetEmployeeSalariesByEmployeeID(logger *logrus.Entry, employeeID uuid.UUID) ([]models.EmployeeSalary, error)
	GetEmployeeSalary(logger *logrus.Entry, id uuid.UUID) (*models.EmployeeSalary, error)
	UpdateEmployeeSalary(logger *logrus.Entry, id uuid.UUID, data *models.EmployeeSalaryUpdate) (*models.EmployeeSalary, error)
	DeleteEmployeeSalary(logger *logrus.Entry, id uuid.UUID) error

	// Tax Bracket methods
	CreateTaxBracket(logger *logrus.Entry, data *models.TaxBracketCreate) (*models.TaxBracket, error)
	GetTaxBracketByID(logger *logrus.Entry, id uuid.UUID) (*models.TaxBracket, error)
	GetTaxBrackets(logger *logrus.Entry, country string, year int) ([]models.TaxBracket, error)
	UpdateTaxBracket(logger *logrus.Entry, id uuid.UUID, data *models.TaxBracketUpdate) (*models.TaxBracket, error)
	DeleteTaxBracket(logger *logrus.Entry, id uuid.UUID) error

	// Payroll methods
	CreatePayroll(logger *logrus.Entry, data *models.PayrollCreate) (*models.Payroll, error)
	GetPayrollByID(logger *logrus.Entry, id uuid.UUID) (*models.Payroll, error)
	ListPayrolls(logger *logrus.Entry) ([]models.Payroll, error)
	UpdatePayroll(logger *logrus.Entry, id uuid.UUID, data *models.PayrollUpdate) (*models.Payroll, error)

	// Payroll Detail methods
	CreatePayrollDetail(logger *logrus.Entry, data *models.PayrollDetailCreate) (*models.PayrollDetail, error)
	GetPayrollDetailsByPayrollID(logger *logrus.Entry, payrollID uuid.UUID) ([]models.PayrollDetail, error)

	// Payslip methods
	CreatePayslip(logger *logrus.Entry, data *models.PayslipCreate) (*models.Payslip, error)
	GetPayslip(logger *logrus.Entry, id uuid.UUID) (*models.Payslip, error)
}

type repository struct {
	db *database.DB
}

// NewRepository creates a new payroll repository
func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

func logQuery(logger *logrus.Entry, query string, startTime time.Time) {
	logger.WithFields(logrus.Fields{
		"query":    query,
		"duration": time.Since(startTime),
	}).Debug("Executed payroll query")
}

// --- Salary Component ---

func (r *repository) CreateSalaryComponent(logger *logrus.Entry, data *models.SalaryComponentCreate) (*models.SalaryComponent, error) {
	startTime := time.Now()
	var comp models.SalaryComponent
	query := `INSERT INTO salary_components (name, type, is_taxable, is_recurring)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, name, type, is_taxable, is_recurring, created_at, updated_at`
	err := r.db.QueryRow(query, data.Name, data.Type, data.IsTaxable, data.IsRecurring).Scan(
		&comp.ID, &comp.Name, &comp.Type, &comp.IsTaxable, &comp.IsRecurring, &comp.CreatedAt, &comp.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &comp, err
}

func (r *repository) GetSalaryComponentByID(logger *logrus.Entry, id uuid.UUID) (*models.SalaryComponent, error) {
	startTime := time.Now()
	var comp models.SalaryComponent
	query := `SELECT id, name, type, is_taxable, is_recurring, created_at, updated_at
			  FROM salary_components WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&comp.ID, &comp.Name, &comp.Type, &comp.IsTaxable, &comp.IsRecurring, &comp.CreatedAt, &comp.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &comp, err
}

func (r *repository) ListSalaryComponents(logger *logrus.Entry) ([]models.SalaryComponent, error) {
	startTime := time.Now()
	var comps []models.SalaryComponent
	query := `SELECT id, name, type, is_taxable, is_recurring, created_at, updated_at
			  FROM salary_components`
	rows, err := r.db.Query(query)
	logQuery(logger, query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comp models.SalaryComponent
		if err := rows.Scan(&comp.ID, &comp.Name, &comp.Type, &comp.IsTaxable, &comp.IsRecurring, &comp.CreatedAt, &comp.UpdatedAt); err != nil {
			return nil, err
		}
		comps = append(comps, comp)
	}
	return comps, nil
}

func (r *repository) UpdateSalaryComponent(logger *logrus.Entry, id uuid.UUID, data *models.SalaryComponentUpdate) (*models.SalaryComponent, error) {
	startTime := time.Now()
	var comp models.SalaryComponent
	query := `UPDATE salary_components
			  SET name = $1, type = $2, is_taxable = $3, is_recurring = $4, updated_at = NOW()
			  WHERE id = $5
			  RETURNING id, name, type, is_taxable, is_recurring, created_at, updated_at`
	err := r.db.QueryRow(query, data.Name, data.Type, data.IsTaxable, data.IsRecurring, id).Scan(
		&comp.ID, &comp.Name, &comp.Type, &comp.IsTaxable, &comp.IsRecurring, &comp.CreatedAt, &comp.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &comp, err
}

func (r *repository) DeleteSalaryComponent(logger *logrus.Entry, id uuid.UUID) error {
	startTime := time.Now()
	query := "DELETE FROM salary_components WHERE id = $1"
	_, err := r.db.Exec(query, id)
	logQuery(logger, query, startTime)
	return err
}

// --- Employee Salary ---

func (r *repository) CreateEmployeeSalary(logger *logrus.Entry, data *models.EmployeeSalaryCreate) (*models.EmployeeSalary, error) {
	startTime := time.Now()
	var s models.EmployeeSalary
	query := `INSERT INTO employee_salaries (employee_id, salary_component_id, amount, effective_date, end_date)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, employee_id, salary_component_id, amount, effective_date, end_date, created_at, updated_at`
	err := r.db.QueryRow(query, data.EmployeeID, data.SalaryComponentID, data.Amount, data.EffectiveDate, data.EndDate).Scan(
		&s.ID, &s.EmployeeID, &s.SalaryComponentID, &s.Amount, &s.EffectiveDate, &s.EndDate, &s.CreatedAt, &s.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &s, err
}
func (r *repository) GetEmployeeSalary(logger *logrus.Entry, id uuid.UUID) (*models.EmployeeSalary, error) {
	startTime := time.Now()
	var s models.EmployeeSalary
	query := `SELECT id, employee_id, salary_component_id, amount, effective_date, end_date, created_at, updated_at
			  FROM employee_salaries WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&s.ID, &s.EmployeeID, &s.SalaryComponentID, &s.Amount, &s.EffectiveDate, &s.EndDate, &s.CreatedAt, &s.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &s, err
}

func (r *repository) GetEmployeeSalariesByEmployeeID(logger *logrus.Entry, employeeID uuid.UUID) ([]models.EmployeeSalary, error) {
	startTime := time.Now()
	var salaries []models.EmployeeSalary
	query := `SELECT id, employee_id, salary_component_id, amount, effective_date, end_date, created_at, updated_at
			  FROM employee_salaries WHERE employee_id = $1 AND (end_date IS NULL OR end_date > NOW())`
	rows, err := r.db.Query(query, employeeID)
	logQuery(logger, query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s models.EmployeeSalary
		if err := rows.Scan(&s.ID, &s.EmployeeID, &s.SalaryComponentID, &s.Amount, &s.EffectiveDate, &s.EndDate, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		salaries = append(salaries, s)
	}
	return salaries, nil
}

func (r *repository) UpdateEmployeeSalary(logger *logrus.Entry, id uuid.UUID, data *models.EmployeeSalaryUpdate) (*models.EmployeeSalary, error) {
	startTime := time.Now()
	var s models.EmployeeSalary
	query := `UPDATE employee_salaries
			  SET amount = $1, end_date = $2, updated_at = NOW()
			  WHERE id = $3
			  RETURNING id, employee_id, salary_component_id, amount, effective_date, end_date, created_at, updated_at`
	err := r.db.QueryRow(query, data.Amount, data.EndDate, id).Scan(
		&s.ID, &s.EmployeeID, &s.SalaryComponentID, &s.Amount, &s.EffectiveDate, &s.EndDate, &s.CreatedAt, &s.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &s, err
}

func (r *repository) DeleteEmployeeSalary(logger *logrus.Entry, id uuid.UUID) error {
	startTime := time.Now()
	query := "DELETE FROM employee_salaries WHERE id = $1"
	_, err := r.db.Exec(query, id)
	logQuery(logger, query, startTime)
	return err
}

// --- Tax Bracket ---

func (r *repository) CreateTaxBracket(logger *logrus.Entry, data *models.TaxBracketCreate) (*models.TaxBracket, error) {
	startTime := time.Now()
	var tb models.TaxBracket
	query := `INSERT INTO tax_brackets (country, tax_year, bracket_min, bracket_max, tax_rate)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, country, tax_year, bracket_min, bracket_max, tax_rate, created_at, updated_at`
	err := r.db.QueryRow(query, data.Country, data.TaxYear, data.BracketMin, data.BracketMax, data.TaxRate).Scan(
		&tb.ID, &tb.Country, &tb.TaxYear, &tb.BracketMin, &tb.BracketMax, &tb.TaxRate, &tb.CreatedAt, &tb.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &tb, err
}

func (r *repository) GetTaxBracketByID(logger *logrus.Entry, id uuid.UUID) (*models.TaxBracket, error) {
	startTime := time.Now()
	var tb models.TaxBracket
	query := `SELECT id, country, tax_year, bracket_min, bracket_max, tax_rate, created_at, updated_at
			  FROM tax_brackets WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&tb.ID, &tb.Country, &tb.TaxYear, &tb.BracketMin, &tb.BracketMax, &tb.TaxRate, &tb.CreatedAt, &tb.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &tb, err
}

func (r *repository) GetTaxBrackets(logger *logrus.Entry, country string, year int) ([]models.TaxBracket, error) {
	startTime := time.Now()
	var brackets []models.TaxBracket
	query := `SELECT id, country, tax_year, bracket_min, bracket_max, tax_rate, created_at, updated_at
			  FROM tax_brackets WHERE country = $1 AND tax_year = $2 ORDER BY bracket_min ASC`
	rows, err := r.db.Query(query, country, year)
	logQuery(logger, query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tb models.TaxBracket
		if err := rows.Scan(&tb.ID, &tb.Country, &tb.TaxYear, &tb.BracketMin, &tb.BracketMax, &tb.TaxRate, &tb.CreatedAt, &tb.UpdatedAt); err != nil {
			return nil, err
		}
		brackets = append(brackets, tb)
	}
	return brackets, nil
}

func (r *repository) UpdateTaxBracket(logger *logrus.Entry, id uuid.UUID, data *models.TaxBracketUpdate) (*models.TaxBracket, error) {
	startTime := time.Now()
	var tb models.TaxBracket
	query := `UPDATE tax_brackets
			  SET country = $1, tax_year = $2, bracket_min = $3, bracket_max = $4, tax_rate = $5, updated_at = NOW()
			  WHERE id = $6
			  RETURNING id, country, tax_year, bracket_min, bracket_max, tax_rate, created_at, updated_at`
	err := r.db.QueryRow(query, data.Country, data.TaxYear, data.BracketMin, data.BracketMax, data.TaxRate, id).Scan(
		&tb.ID, &tb.Country, &tb.TaxYear, &tb.BracketMin, &tb.BracketMax, &tb.TaxRate, &tb.CreatedAt, &tb.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &tb, err
}

func (r *repository) DeleteTaxBracket(logger *logrus.Entry, id uuid.UUID) error {
	startTime := time.Now()
	query := "DELETE FROM tax_brackets WHERE id = $1"
	_, err := r.db.Exec(query, id)
	logQuery(logger, query, startTime)
	return err
}

// --- Payroll ---

func (r *repository) CreatePayroll(logger *logrus.Entry, data *models.PayrollCreate) (*models.Payroll, error) {
	startTime := time.Now()
	var p models.Payroll
	query := `INSERT INTO payroll (pay_period_start, pay_period_end, payment_date)
			  VALUES ($1, $2, $3)
			  RETURNING id, pay_period_start, pay_period_end, payment_date, status, total_gross_pay, total_deductions, total_net_pay, created_at, updated_at`
	err := r.db.QueryRow(query, data.PayPeriodStart, data.PayPeriodEnd, data.PaymentDate).Scan(
		&p.ID, &p.PayPeriodStart, &p.PayPeriodEnd, &p.PaymentDate, &p.Status, &p.TotalGrossPay, &p.TotalDeductions, &p.TotalNetPay, &p.CreatedAt, &p.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &p, err
}

func (r *repository) GetPayrollByID(logger *logrus.Entry, id uuid.UUID) (*models.Payroll, error) {
	startTime := time.Now()
	var p models.Payroll
	query := `SELECT id, pay_period_start, pay_period_end, payment_date, status, total_gross_pay, total_deductions, total_net_pay, created_at, updated_at
			  FROM payroll WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&p.ID, &p.PayPeriodStart, &p.PayPeriodEnd, &p.PaymentDate, &p.Status, &p.TotalGrossPay, &p.TotalDeductions, &p.TotalNetPay, &p.CreatedAt, &p.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &p, err
}

func (r *repository) ListPayrolls(logger *logrus.Entry) ([]models.Payroll, error) {
	startTime := time.Now()
	var payrolls []models.Payroll
	query := `SELECT id, pay_period_start, pay_period_end, payment_date, status, total_gross_pay, total_deductions, total_net_pay, created_at, updated_at
			  FROM payroll`
	rows, err := r.db.Query(query)
	logQuery(logger, query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Payroll
		if err := rows.Scan(&p.ID, &p.PayPeriodStart, &p.PayPeriodEnd, &p.PaymentDate, &p.Status, &p.TotalGrossPay, &p.TotalDeductions, &p.TotalNetPay, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		payrolls = append(payrolls, p)
	}
	return payrolls, nil
}

func (r *repository) UpdatePayroll(logger *logrus.Entry, id uuid.UUID, data *models.PayrollUpdate) (*models.Payroll, error) {
	startTime := time.Now()
	var p models.Payroll
	query := `UPDATE payroll
			  SET status = $1, total_gross_pay = $2, total_deductions = $3, total_net_pay = $4, updated_at = NOW()
			  WHERE id = $5
			  RETURNING id, pay_period_start, pay_period_end, payment_date, status, total_gross_pay, total_deductions, total_net_pay, created_at, updated_at`
	err := r.db.QueryRow(query, data.Status, data.TotalGrossPay, data.TotalDeductions, data.TotalNetPay, id).Scan(
		&p.ID, &p.PayPeriodStart, &p.PayPeriodEnd, &p.PaymentDate, &p.Status, &p.TotalGrossPay, &p.TotalDeductions, &p.TotalNetPay, &p.CreatedAt, &p.UpdatedAt,
	)
	logQuery(logger, query, startTime)
	return &p, err
}

// --- Payroll Detail ---

func (r *repository) CreatePayrollDetail(logger *logrus.Entry, data *models.PayrollDetailCreate) (*models.PayrollDetail, error) {
	startTime := time.Now()
	var pd models.PayrollDetail
	query := `INSERT INTO payroll_details (payroll_id, employee_id, gross_pay, tax_amount, other_deductions, net_pay)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, payroll_id, employee_id, gross_pay, tax_amount, other_deductions, net_pay, created_at`
	err := r.db.QueryRow(query, data.PayrollID, data.EmployeeID, data.GrossPay, data.TaxAmount, data.OtherDeductions, data.NetPay).Scan(
		&pd.ID, &pd.PayrollID, &pd.EmployeeID, &pd.GrossPay, &pd.TaxAmount, &pd.OtherDeductions, &pd.NetPay, &pd.CreatedAt,
	)
	logQuery(logger, query, startTime)
	return &pd, err
}

func (r *repository) GetPayrollDetailsByPayrollID(logger *logrus.Entry, payrollID uuid.UUID) ([]models.PayrollDetail, error) {
	startTime := time.Now()
	var details []models.PayrollDetail
	query := `SELECT id, payroll_id, employee_id, gross_pay, tax_amount, other_deductions, net_pay, created_at
			  FROM payroll_details WHERE payroll_id = $1`
	rows, err := r.db.Query(query, payrollID)
	logQuery(logger, query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pd models.PayrollDetail
		if err := rows.Scan(&pd.ID, &pd.PayrollID, &pd.EmployeeID, &pd.GrossPay, &pd.TaxAmount, &pd.OtherDeductions, &pd.NetPay, &pd.CreatedAt); err != nil {
			return nil, err
		}
		details = append(details, pd)
	}
	return details, nil
}

// --- Payslip ---

func (r *repository) CreatePayslip(logger *logrus.Entry, data *models.PayslipCreate) (*models.Payslip, error) {
	startTime := time.Now()
	var ps models.Payslip
	query := `INSERT INTO payslips (employee_id, payroll_id, pay_period_start, pay_period_end, gross_pay, tax_amount, deductions, net_pay, file_path)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			  RETURNING id, employee_id, payroll_id, pay_period_start, pay_period_end, gross_pay, tax_amount, deductions, net_pay, file_path, created_at`
	err := r.db.QueryRow(query, data.EmployeeID, data.PayrollID, data.PayPeriodStart, data.PayPeriodEnd, data.GrossPay, data.TaxAmount, data.Deductions, data.NetPay, data.FilePath).Scan(
		&ps.ID, &ps.EmployeeID, &ps.PayrollID, &ps.PayPeriodStart, &ps.PayPeriodEnd, &ps.GrossPay, &ps.TaxAmount, &ps.Deductions, &ps.NetPay, &ps.FilePath, &ps.CreatedAt,
	)
	logQuery(logger, query, startTime)
	return &ps, err
}

func (r *repository) GetPayslip(logger *logrus.Entry, id uuid.UUID) (*models.Payslip, error) {
	startTime := time.Now()
	var ps models.Payslip
	query := `SELECT id, employee_id, payroll_id, pay_period_start, pay_period_end, gross_pay, tax_amount, deductions, net_pay, file_path, created_at
			  FROM payslips WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&ps.ID, &ps.EmployeeID, &ps.PayrollID, &ps.PayPeriodStart, &ps.PayPeriodEnd, &ps.GrossPay, &ps.TaxAmount, &ps.Deductions, &ps.NetPay, &ps.FilePath, &ps.CreatedAt,
	)
	logQuery(logger, query, startTime)
	return &ps, err
}
