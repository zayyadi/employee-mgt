CREATE TABLE positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    department_id UUID NOT NULL REFERENCES departments(id),
    description TEXT NOT NULL,
    requirements TEXT NOT NULL,
    salary_range_min DECIMAL(10, 2) NOT NULL,
    salary_range_max DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
