// @ts-nocheck
import { error } from "@sveltejs/kit";


export async function GET({ request }) {
    if (!request.headers.get("referer")) {
        return error(404, "Not found")
    }
    const id = request.headers.get("Event");
    try {
        const response = await fetch(`http://api:8000/event/${id}`);

        const jsonRes = await response.json()

        console.log(id)
        return jsonResponse(jsonRes, response.status)
    }
    catch (error) {
        return jsonResponse({ message: "An error occurred", error }, 500)
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
