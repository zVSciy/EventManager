const API_BASE_URL = 'https://localhost:8443';

export async function apiFetch(endpoint, options = {}) {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, options);
    if (!response.ok) {
        throw new Error(response.statusText);
    }
    return response.json();
}
