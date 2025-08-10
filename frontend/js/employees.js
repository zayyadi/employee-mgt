document.addEventListener('DOMContentLoaded', function () {
    const apiUrl = 'http://localhost:8080/api/v1/employees';
    const employeeTableBody = document.getElementById('employee-table-body');
    const employeeModal = new bootstrap.Modal(document.getElementById('employeeModal'));
    const employeeForm = document.getElementById('employeeForm');
    const createEmployeeBtn = document.getElementById('createEmployeeBtn');
    const modalTitle = document.getElementById('employeeModalLabel');

    // --- Data Fetching and Rendering ---

    async function fetchEmployees() {
        try {
            const response = await fetch(apiUrl);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const employees = await response.json();
            renderEmployees(employees);
        } catch (error) {
            console.error('Error fetching employees:', error);
        }
    }

    function renderEmployees(employees) {
        employeeTableBody.innerHTML = '';
        if (!employees || employees.length === 0) {
            return;
        }
        employees.forEach(employee => {
            const row = `
                <tr>
                    <td>${employee.ID.substring(0, 8)}...</td>
                    <td>${employee.first_name}</td>
                    <td>${employee.last_name}</td>
                    <td>${employee.email}</td>
                    <td>${employee.PositionID.substring(0, 8)}...</td>
                    <td>
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${employee.ID}">Edit</button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${employee.ID}">Delete</button>
                    </td>
                </tr>
            `;
            employeeTableBody.innerHTML += row;
        });
    }

    // --- Event Handlers ---

    createEmployeeBtn.addEventListener('click', () => {
        modalTitle.textContent = 'Add Employee';
        employeeForm.reset();
        document.getElementById('employeeId').value = '';
    });

    employeeTableBody.addEventListener('click', async (event) => {
        const target = event.target;
        const id = target.dataset.id;

        if (target.classList.contains('edit-btn')) {
            const employee = await getEmployeeById(id);
            if (employee) {
                modalTitle.textContent = 'Edit Employee';
                document.getElementById('employeeId').value = employee.ID;
                document.getElementById('firstName').value = employee.first_name;
                document.getElementById('lastName').value = employee.last_name;
                document.getElementById('email').value = employee.email;
                document.getElementById('position').value = employee.PositionID; // Assuming this is how it's stored
                document.getElementById('hireDate').value = new Date(employee.hire_date).toISOString().split('T')[0];
                employeeModal.show();
            }
        }

        if (target.classList.contains('delete-btn')) {
            if (confirm('Are you sure you want to delete this employee?')) {
                await deleteEmployee(id);
            }
        }
    });

    employeeForm.addEventListener('submit', async (event) => {
        event.preventDefault();
        const id = document.getElementById('employeeId').value;
        const employeeData = {
            first_name: document.getElementById('firstName').value,
            last_name: document.getElementById('lastName').value,
            email: document.getElementById('email').value,
            position_id: document.getElementById('position').value, // This should be a UUID
            hire_date: document.getElementById('hireDate').value,
            // You would need to gather all other required fields from the form
            // For simplicity, we are only sending a few fields.
            // The backend expects more fields, so this will need to be expanded.
            // Example of other fields:
            user_id: "00000000-0000-0000-0000-000000000000", // Placeholder
            employee_id: "EMP" + new Date().getTime(), // Placeholder
            date_of_birth: "1990-01-01", // Placeholder
            gender: "other", // Placeholder
            marital_status: "single", // Placeholder
            phone_number: "123-456-7890", // Placeholder
            address: "123 Main St", // Placeholder
            emergency_contact_name: "Jane Doe", // Placeholder
            emergency_contact_phone: "098-765-4321", // Placeholder
            department_id: "00000000-0000-0000-0000-000000000000", // Placeholder
            employment_status: "active", // Placeholder
        };

        if (id) {
            // Update
            await updateEmployee(id, employeeData);
        } else {
            // Create
            await createEmployee(employeeData);
        }
    });

    // --- API Functions ---

    async function createEmployee(data) {
        try {
            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data),
            });
            if (!response.ok) throw new Error('Failed to create employee');
            employeeModal.hide();
            fetchEmployees();
        } catch (error) {
            console.error('Error creating employee:', error);
        }
    }

    async function getEmployeeById(id) {
        try {
            const response = await fetch(`${apiUrl}/${id}`);
            if (!response.ok) throw new Error('Failed to get employee');
            return await response.json();
        } catch (error) {
            console.error('Error getting employee:', error);
        }
    }

    async function updateEmployee(id, data) {
        try {
            const response = await fetch(`${apiUrl}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data),
            });
            if (!response.ok) throw new Error('Failed to update employee');
            employeeModal.hide();
            fetchEmployees();
        } catch (error) {
            console.error('Error updating employee:', error);
        }
    }

    async function deleteEmployee(id) {
        try {
            const response = await fetch(`${apiUrl}/${id}`, { method: 'DELETE' });
            if (!response.ok) throw new Error('Failed to delete employee');
            fetchEmployees();
        } catch (error) {
            console.error('Error deleting employee:', error);
        }
    }

    // --- Initial Load ---
    fetchEmployees();
});
