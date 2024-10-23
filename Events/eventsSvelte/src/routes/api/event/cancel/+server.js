import { json } from '@sveltejs/kit';

const API_BASE_URL = 'http://api:8000'; 

export async function PUT({ request, url }) {
    const id = url.searchParams.get('id'); // Fetch event ID from query parameters

    if (!id) {
        return json({ message: "Event ID is required" }, { status: 400 });
    }

    const { canceledStatus } = await request.json(); // Destructure canceled status from request body

    try {
        const response = await fetch(`${API_BASE_URL}/event/cancel/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ canceled: canceledStatus }), // Update cancel status in request body
        });

        if (!response.ok) {
            throw new Error(`Failed to update event status for event ID ${id}`);
        }

        const data = await response.json();
        return json(data, { status: response.status }); // Return the updated event data
    } catch (err) {
        return json({ message: "An error occurred", error: err.message }, { status: 500 });
    }
}
