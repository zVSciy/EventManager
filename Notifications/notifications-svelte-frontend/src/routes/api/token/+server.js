import https from 'https';

// Bypass SSL certificate validation (development only)
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

const API_BASE_URL = 'https://auth_api:8000';

function jsonResponse(data, status) {
    return new Response(JSON.stringify(data), {
        status: status,
        headers: {
            "Content-Type": "application/json",
        }
    });
}

export async function POST({ url }) {
    const email = url.searchParams.get('email');
    const password = url.searchParams.get('password');

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

        return jsonResponse({status: response.status});
    } 
    
    catch (error) {
        return jsonResponse({status: 500, error: error});
    }
}