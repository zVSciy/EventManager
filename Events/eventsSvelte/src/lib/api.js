const API_BASE_URL = 'http://localhost:8000'; 


export async function fetchEvents() {
    const response = await fetch(`${API_BASE_URL}/event/`);
    if (!response.ok) {
        throw new Error('Failed to fetch events');
    }
    return response.json();
}
export async function updateEvent(eventId, updatedEventData) {
    const response = await fetch(`${API_BASE_URL}/event/${eventId}/`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(updatedEventData),
    });

    if (!response.ok) {
        throw new Error('Failed to update event');
    }

    return response.json();
}



export async function fetchEventById(eventId) {
    const response = await fetch(`http://127.0.0.1:8000/event/${eventId}`);
    if (!response.ok) {
        throw new Error(`Failed to fetch event with ID: ${eventId}`);
    }
    return await response.json();
}


export async function updateCancelStatus(eventId, canceledStatus) {
    const response = await fetch(`${API_BASE_URL}/event/cancel/${eventId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ canceled: canceledStatus }),
    });

    if (!response.ok) {
        throw new Error('Failed to update event status');
    }

    return response.json();
}
