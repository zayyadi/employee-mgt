-- Remove foreign key constraints from employees table
ALTER TABLE employees
    DROP CONSTRAINT IF EXISTS employees_department_id_fkey,
    DROP CONSTRAINT IF EXISTS employees_position_id_fkey,
    DROP CONSTRAINT IF EXISTS employees_manager_id_fkey;
