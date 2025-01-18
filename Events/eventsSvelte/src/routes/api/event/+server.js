import { error } from "@sveltejs/kit";

const API_BASE_URL = 'http://events_api:8000';

export async function GET() {
    try {
     
        const response = await fetch(`${API_BASE_URL}/event/`);
        if (!response.ok) {
            throw new Error('Failed to fetch events');
        }
        return jsonResponse(await response.json(), response.status);
    } catch (err) {
        return jsonResponse({ message: "An error occurred", error: err }, 500);
    }
}
export async function POST({ request }) {
    const { name, location, organisator, startdate, available_normal_tickets, available_vip_tickets } = await request.json();

    try {
        const response = await fetch(`${API_BASE_URL}/event`, {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                location: location,
                organisator: organisator,
                startdate: startdate,
                available_normal_tickets: available_normal_tickets,
                available_vip_tickets: available_vip_tickets
            })
        });

        return jsonResponse(response, response.status);
    } catch (error) {
        return jsonResponse({ message: "An error occurred", error }, 500);
    }
}
export async function PUT({ request, url }) {
    const Id = url.searchParams.get('id'); // Extract the ID from the URL query params
    const updatedEventData = await request.json();

    if (!Id) {
        return jsonResponse({ message: "Event ID is required" }, 400);
    }

    try {
        const response = await fetch(`${API_BASE_URL}/event/${Id}/`, { // Pass ID in the URL
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedEventData),
        });

        if (!response.ok) {
            throw new Error('Failed to update event');
        }

        return jsonResponse(await response.json(), response.status);
    } catch (err) {
        return jsonResponse({ message: "An error occurred", error: err }, 500);
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
