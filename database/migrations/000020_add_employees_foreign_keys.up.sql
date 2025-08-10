-- Add foreign key constraints to employees table
ALTER TABLE employees
    ADD CONSTRAINT employees_department_id_fkey FOREIGN KEY (department_id) REFERENCES departments(id),
    ADD CONSTRAINT employees_position_id_fkey FOREIGN KEY (position_id) REFERENCES positions(id),
    ADD CONSTRAINT employees_manager_id_fkey FOREIGN KEY (manager_id) REFERENCES employees(id);
