CREATE TABLE payroll (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pay_period_start DATE NOT NULL,
    pay_period_end DATE NOT NULL,
    payment_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    total_gross_pay DECIMAL(12, 2) NOT NULL DEFAULT 0,
    total_deductions DECIMAL(12, 2) NOT NULL DEFAULT 0,
    total_net_pay DECIMAL(12, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
