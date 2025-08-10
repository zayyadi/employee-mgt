-- Delete seed data
DELETE FROM salary_components WHERE name IN (
    'Basic Salary', 'Housing Allowance', 'Transport Allowance', 
    'Overtime', 'Bonus', 'Tax Deduction', 'Pension Contribution', 
    'Health Insurance'
);

DELETE FROM leave_types WHERE name IN (
    'Annual Leave', 'Sick Leave', 'Maternity Leave', 
    'Paternity Leave', 'Bereavement Leave'
);
