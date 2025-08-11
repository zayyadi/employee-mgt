import { writable } from 'svelte/store';
import { apiFetch } from '$lib/utils/api';

// Define the shape of the store's state
interface EmployeeStore {
  employees: any[]; // Replace 'any' with a proper Employee type later
  loading: boolean;
  error: string | null;
}

// Create the writable store with an initial state
const { subscribe, set, update } = writable<EmployeeStore>({
  employees: [],
  loading: false,
  error: null,
});

// Function to fetch all employees
async function fetchAll() {
  update(state => ({ ...state, loading: true, error: null }));
  try {
    const employees = await apiFetch('/employees');
    update(state => ({ ...state, employees: employees, loading: false }));
  } catch (e: any) {
    update(state => ({ ...state, error: e.message, loading: false }));
  }
}

// Expose the store's methods
export default {
  subscribe,
  fetchAll,
};
