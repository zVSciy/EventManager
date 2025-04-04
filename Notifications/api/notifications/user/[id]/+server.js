export async function GET({ params }) {
    const userId = params.id;
    const backend_url = 'http://notification_api:8000';
    const backendURL = `${backend_url}/notifications/user/${userId}`;
    try {
        const res = await fetch(backendURL);
        console.log(res.json());
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
// ...existing code...
