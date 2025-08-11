import { browser } from '$app/environment';

const baseURL = browser ? 'http://localhost:8080/api' : 'http://backend:8080/api';

// A simple fetch wrapper
export async function apiFetch(path: string, options: RequestInit = {}) {
  const defaultHeaders = {
    'Content-Type': 'application/json',
  };

  const mergedOptions: RequestInit = {
    ...options,
    headers: {
      ...defaultHeaders,
      ...options.headers,
    },
  };

  const response = await fetch(`${baseURL}${path}`, mergedOptions);

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({ message: 'An unknown error occurred' }));
    throw new Error(errorData.message || 'API request failed');
  }

  // Handle responses with no content
  if (response.status === 204) {
    return null;
  }

  return response.json();
}
