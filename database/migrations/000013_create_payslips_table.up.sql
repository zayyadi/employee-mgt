CREATE TABLE payslips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    employee_id UUID NOT NULL REFERENCES employees(id),
    payroll_id UUID NOT NULL REFERENCES payroll(id),
    pay_period_start DATE NOT NULL,
    pay_period_end DATE NOT NULL,
    gross_pay DECIMAL(10, 2) NOT NULL,
    tax_amount DECIMAL(10, 2) NOT NULL,
    deductions JSONB,
    net_pay DECIMAL(10, 2) NOT NULL,
    file_path VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
