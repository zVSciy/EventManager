import { error } from "@sveltejs/kit";
import https from 'https';

// Bypass SSL certificate validation (development only)
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

const API_BASE_URL = 'https://backend:8000';

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
                password: password,
                first_name: first_name,
                last_name: last_name,
                role: role
            }),
            agent: new https.Agent({
                rejectUnauthorized: false
            })
        });

        return jsonResponse(response.status);
    } 
    
    catch (error) {
        console.log(error)
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