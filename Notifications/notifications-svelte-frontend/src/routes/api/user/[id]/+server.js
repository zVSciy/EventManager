export async function GET({ params }) {
    const userId = params.id;
    // Adjust the backend URL if necessary
    const backendURL = `http://notification_api:8000/notifications/user/${userId}`;
    try {
        const res = await fetch(backendURL);
        if (!res.ok) {
            return new Response(JSON.stringify({ error: "Failed to fetch notifications" }), { status: res.status });
        }
        const data = await res.json();
        return new Response(JSON.stringify(data), {
            status: 200,
            headers: { 'Content-Type': 'application/json' }
        });
    } catch (error) {
        return new Response(JSON.stringify({ error: error.message }), { status: 500 });
    }
}