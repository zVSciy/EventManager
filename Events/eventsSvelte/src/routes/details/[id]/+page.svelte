<script>
    import { base } from '$app/paths';
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { writable } from 'svelte/store';
    let name = '';
    let location = '';
    let organisator = '';
    let startDate = '';
    let available_normal_tickets = 1;
    let available_vip_tickets = 1;
    let canceled = '';
    
    let event = writable([]);
    console.log(event)
    async function fetchEvent() {
        const response = await fetch(`${base}/api/event/details`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json; charset=UTF-8",
                "Event": event.ID,
            },
        });
        const data = await response.json();
        event.set(data);
        name=data.name;
        location=data.location;
        organisator=data.organisator;
        startDate=new Date(data.startdate).toLocaleString();
        available_normal_tickets=data.available_normal_tickets;
        available_vip_tickets=data.available_vip_tickets
        canceled = data.canceled ? 'Yes' : 'No'
    }
    function goToIndexPage() {
        goto(`${base}`);
    }
    function goToFeedBack() {
        const eventId = event.ID;
        sessionStorage.setItem('eventId', eventId);
        const targetUrl = `/?eventId=${eventId}`;
        window.location.href = targetUrl;  
    }    
    function goToTickets() {
        const eventId = event.ID;
        sessionStorage.setItem('eventId', eventId);
        const targetUrl = `/app_ticket?eventId=${eventId}`;
        window.location.href = targetUrl; 
    }
    onMount(() => {
        event.ID = window.location.href.split(`/`).pop();
        fetchEvent();
    });
</script>


<!-- HTML-Ausgabe für die Event-Details -->
<div class="container">
    <h2>Event Details</h2>
    
    <div class="event-info">
        <p><strong>Name:</strong> {name}</p>
        <p><strong>Location:</strong> {location}</p>
        <p><strong>Organisator:</strong> {organisator}</p>
        <p><strong>Start Date:</strong> {startDate}</p>
        <p><strong>Available Normal Tickets:</strong> {available_normal_tickets}</p>
        <p><strong>Available VIP Tickets:</strong> {available_vip_tickets}</p>
        <p><strong>Canceled:</strong> {canceled}</p>

    </div>

    <div class="button-group">
        <button on:click={goToIndexPage}>Back to Event List</button>
        <button on:click={goToTickets}>Buy Tickets</button>
        <button on:click={goToFeedBack}>Give Feedback</button>
    </div>
</div>

<style>
    /* Allgemeines Styling */
    .container {
        max-width: 600px;
        margin: 2em auto;
        padding: 20px;
        border: 1px solid #ccc;
        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        font-family: 'Arial', sans-serif;
        background-color: #f9f9f9;
    }

    h2 {
        text-align: center;
        color: #333;
    }

    .event-info p {
        font-size: 16px;
        line-height: 1.5;
        margin: 10px 0;
    }

    .event-info p strong {
        color: #555;
    }

    button {
        display: block;
        width: 100%;
        padding: 10px;
        background-color: #343a40;
        color: #fff;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        font-size: 16px;
        margin-top: 20px;
    }

    button:hover {
        background-color: #555;
    }
    .styled-table {
        width: 100%;
        border-collapse: collapse;
        margin: 25px 0;
        font-size: 18px;
        font-family: Arial, sans-serif;
        box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
        border-radius: 12px;
        overflow: hidden;
    }

    .styled-table thead tr {
        background-color: #009879;
        color: #ffffff;
        text-align: left;
        font-weight: bold;
    }

    .styled-table th,
    .styled-table td {
        padding: 12px 15px;
    }

    .styled-table tbody tr {
        border-bottom: 1px solid #dddddd;
    }

    .styled-table tbody tr:nth-of-type(even) {
        background-color: #f3f3f3;
    }

    .styled-table tbody tr:last-of-type {
        border-bottom: 2px solid #009879;
    }

    .styled-table tbody tr:hover {
        background-color: #f1f1f1;
        cursor: pointer;
    }

    .styled-table td a {
        text-decoration: none;
        color: #009879;
        font-weight: bold;
    }

    .styled-table td a:hover {
        text-decoration: underline;
    }

    button {
        padding: 8px 12px;
        border: none;
        border-radius: 5px;
        background-color: #009879;
        color: white;
        cursor: pointer;
    }

    button:hover {
        background-color: #007f68;
    }

    .edit-form {
        margin-top: 20px;
        align-items: center;
        justify-content: center;
        display: flex;
    }

    .edit-form form {
        display: flex;
        flex-direction: column;
        width: 300px;
    }

    .edit-form label {
        margin: 10px 0 5px 0;
    }

    .edit-form input {
        padding: 8px;
        margin-bottom: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    .edit-form button {
        margin-top: 10px;
        padding: 8px;
        background-color: #009879;
        color: white;
        border: none;
        border-radius: 5px;
        cursor: pointer;
    }

    .edit-form button:hover {
        background-color: #007f68;
    }
    h1 {
        display: flex;
        align-items: center;
        justify-content: center;
        display: flex;
    }

</style>