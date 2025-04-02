import { json } from '@sveltejs/kit';

const backend_url = 'http://notification_api:8000';

export async function GET() {
    try {
        const response = await fetch(`${backend_url}/notifications`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const notifications = await response.json();
        return json(notifications);
    } catch (error) {
        console.error("Error fetching notifications:", error);
        return json({ error: error.message }, { status: 500 });
    }
}

export async function POST({ request }) {
    try {
        const newNotification = await request.json();
        const response = await fetch(`${backend_url}/notifications`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(newNotification)
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        return json(result);
    } catch (error) {
        console.error("Error creating notification:", error);
        return json({ error: error.message }, { status: 500 });
    }
}
