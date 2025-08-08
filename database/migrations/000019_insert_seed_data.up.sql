-- Insert default leave types
INSERT INTO leave_types (name, description, max_days_per_year, is_accrued) VALUES
    ('Annual Leave', 'Paid time off for vacation and personal activities', 20, true),
    ('Sick Leave', 'Paid time off for illness or medical appointments', 10, true),
    ('Maternity Leave', 'Leave for childbirth and bonding with newborn', 90, false),
    ('Paternity Leave', 'Leave for bonding with newborn child', 14, false),
    ('Bereavement Leave', 'Leave for dealing with death of family member', 5, false)
ON CONFLICT (name) DO NOTHING;

-- Insert default salary components
INSERT INTO salary_components (name, type, is_taxable, is_recurring) VALUES
    ('Basic Salary', 'earning', true, true),
    ('Housing Allowance', 'earning', true, true),
    ('Transport Allowance', 'earning', true, true),
    ('Overtime', 'earning', true, false),
    ('Bonus', 'earning', true, false),
    ('Tax Deduction', 'deduction', false, true),
    ('Pension Contribution', 'deduction', false, true),
    ('Health Insurance', 'deduction', false, true)
ON CONFLICT (name) DO NOTHING;
