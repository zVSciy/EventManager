import { json } from '@sveltejs/kit';

export async function POST({ request }) {
    try {
        const { email, password, first_name, last_name, role } = await request.json();
        
        // Simulate user registration
        const user = { email, first_name, last_name, role };

        // Simulate saving the user to the database
        console.log('User registered:', user);

        return json({ message: 'User registered successfully!', user });
    } catch (error) {
        console.error('Error registering user:', error);
        return json({ message: 'Internal Server Error' }, { status: 500 });
    }
}