import { json } from '@sveltejs/kit';

const backend_url = 'http://notification_api:8000';

export async function DELETE({ params }) {
    const { id } = params;

    try {
        const response = await fetch(`${backend_url}/notifications/${id}`, {
            method: 'DELETE'
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return json({ success: true });
    } catch (error) {
        console.error("Error deleting notification:", error);
        return json({ error: error.message }, { status: 500 });
    }
}

export async function PUT({ request, params }) {
    const { id } = params;

    try {
        const updateNotificationData = await request.json();
        const response = await fetch(`${backend_url}/notifications/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(updateNotificationData)
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        return json(result);
    } catch (error) {
        console.error("Error updating notification:", error);
        return json({ error: error.message }, { status: 500 });
    }
}
