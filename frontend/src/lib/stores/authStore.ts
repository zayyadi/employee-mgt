import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { apiFetch } from '$lib/utils/api';

// Define the shape of the auth state
interface AuthState {
  token: string | null;
  user: any | null; // Replace 'any' with a proper User type later
}

// Function to get initial state from localStorage
function getInitialState(): AuthState {
  if (browser) {
    const token = localStorage.getItem('authToken');
    // In a real app, you might also store and retrieve user info
    if (token) {
      return { token, user: null }; // Initially user is null, would be fetched with token
    }
  }
  return { token: null, user: null };
}

// Create the writable store
const authStore = writable<AuthState>(getInitialState());

// Login function
async function login(email, password) {
  const response = await apiFetch('/auth/login', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
  });

  if (response.token) {
    authStore.set({ token: response.token, user: response.user || null });
    if (browser) {
      localStorage.setItem('authToken', response.token);
    }
  }
}

// Logout function
function logout() {
  authStore.set({ token: null, user: null });
  if (browser) {
    localStorage.removeItem('authToken');
  }
}

export default {
  subscribe: authStore.subscribe,
  login,
  logout,
};
