import { error } from "@sveltejs/kit";

const API_BASE_URL = 'http://backend:8000';

export async function POST({ request }) {
    const { email, password, first_name, last_name, role } = await request.json();

    try {
        const response = await fetch(`${API_BASE_URL}/register`, {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: email,
                hashed_password: password,
                first_name: first_name,
                last_name: last_name,
                role: role
            })
        });

        return jsonResponse(response, response.status);
    } 
    
    catch (error) {
        return jsonResponse({ message: "An error occurred", error }, 500);
    }
}

function jsonResponse(data, status) {
    return new Response(JSON.stringify(data), {
        status: status,
        headers: {
            "Content-Type": "application/json",
        }
    });
}