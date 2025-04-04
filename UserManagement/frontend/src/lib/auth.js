import { writable } from 'svelte/store';
import { apiFetch } from './api';

export const user = writable(null);

export async function login(username, password) {
    const data = await apiFetch('/api/token', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: new URLSearchParams({ username, password }),
    });

    localStorage.setItem('token', data.access_token);
    user.set(data.user);
}

export function logout() {
    localStorage.removeItem('token');
    user.set(null);
}