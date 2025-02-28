import { json } from '@sveltejs/kit';

export async function POST({ request }) {
    const { username, password } = await request.json();
    
    // Simulate user validation
    const isValid = username === 'user' && password === 'pass';

    if (isValid) {
        return json({ access_token: 'valid-token', user: { username } });
    } else {
        return json({ message: 'Invalid credentials' }, { status: 401 });
    }
}

export async function GET({ request }) {
    const token = request.headers.get('Authorization')?.replace('Bearer ', '');
    
    // Simulate token validation
    const isValid = token === 'valid-token';

    if (isValid) {
        return json({ email: 'user@example.com', message: 'Token is valid!' });
    } else {
        return json({ message: 'Invalid or expired token.' }, { status: 401 });
    }
}