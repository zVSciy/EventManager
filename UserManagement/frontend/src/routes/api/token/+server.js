import { error } from "@sveltejs/kit";
import https from 'https';

// Bypass SSL certificate validation (development only)
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

const API_BASE_URL = 'https://auth_api:8000';

export async function POST({ request }) {
    const { email, password } = await request.json();

    try {
        const response = await fetch(`${API_BASE_URL}/token`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: new URLSearchParams({
                username: email, // OAuth2PasswordRequestForm expects "username"
                password: password
            }),
            agent: new https.Agent({
                rejectUnauthorized: false
            })
        });

        return jsonResponse(response.status);
        // const responseData = await response.json();
        // return jsonResponse(responseData, response.status);
    } 
    
    catch (error) {
        console.log(error);
        return jsonResponse({ message: "An error occurred", error: error.message }, 500);
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