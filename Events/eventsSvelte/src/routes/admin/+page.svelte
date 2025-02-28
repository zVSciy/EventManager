<script>
    import { base } from '$app/paths';
    import { onMount } from 'svelte';

    let events = [];
    let error = null;
    let editingEvent = null;

    async function loadEvents() {
        const response = await fetch(`${base}/api/event`);
        events = await response.json();
    }

    onMount(loadEvents);

    async function toggleCancel(event) {
        const newCanceledStatus = !event.canceled; // Toggle current canceled status
        try {
            // Make the API call to update the event cancel status
            const response = await fetch(`${base}/api/event/cancel?id=${event.ID}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    canceledStatus: newCanceledStatus // Send the new canceled status
                }),
            });

            if (!response.ok) {
                throw new Error('Failed to update event cancel status');
            }

            await loadEvents(); // Reload events after status update
        } catch (err) {
            error = 'Failed to update event status';
        }
    }

    function startEditing(event) {
        editingEvent = { ...event }; // Copy the event data for editing
    }

    async function updateEvent(id) {
        try {
            const response = await fetch(`${base}/api/event?id=${id}`, { // Pass ID in the query parameter
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(editingEvent), // Use editingEvent directly
            });

            if (!response.ok) {
                throw new Error('Failed to update event');
            }

            return await response.json();
        } catch (err) {
            console.error('Error updating event:', err);
            throw err;
        }
    }

    async function saveEvent() {
        try {
            await updateEvent(editingEvent.ID); // Pass the ID of the event being edited
            editingEvent = null; // End the editing
            await loadEvents();
        } catch (err) {
            error = 'Failed to update event';
        }
    }

    function cancelEditing() {
        editingEvent = null;
    }
    export let data;
</script>


<nav class="navbar navbar-expand-lg navbar-light bg-warning">
    <div class="container-fluid">
        <button on:click={() => window.location.href=`${base}`} style="margin-bottom: 20px; padding: 10px; background-color: #009879; color: white; border: none; border-radius: 4px; cursor: pointer;">Event Manager</button>    

        <div class="d-flex align-items-center ms-auto">
            <span class="me-3">Hi, {data.username}!</span>

        </div>
    </div>
</nav>



{#if error}
    <p>{error}</p>
{:else}
    <div>
        <h1>Event List</h1>
        {#if data.admin}
        <button on:click={() => window.location.href=`${base}/admin/add-event`} style="margin-bottom: 20px; padding: 10px; background-color: #009879; color: white; border: none; border-radius: 4px; cursor: pointer;">Add New Event</button>    
        {/if}    
        <table class="styled-table">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Location</th>
                    <th>Start Date</th>
                    <th>Available Normal Tickets</th>
                    <th>Available VIP Tickets</th>
                    <th>Canceled</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each events as event}
                    <tr>
                        <td><a href={`${base}/event/${event.ID}`}>{event.ID}</a></td>
                        <td>{event.name}</td>
                        <td>{event.location}</td>
                        <td>{new Date(event.startdate).toLocaleString()}</td>
                        <td>{event.available_normal_tickets}</td>
                        <td>{event.available_vip_tickets}</td>
                        <td>{event.canceled ? 'Yes' : 'No'}</td>
                        <td>
                            {#if data.admin}
                            <button on:click={() => toggleCancel(event)}>
                                {event.canceled ? 'Entcancel Event' : 'Cancel Event'}
                            </button>
                            <button on:click={() => startEditing(event)}>Edit</button>
                            {/if}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>

        <!-- Formular zum Bearbeiten eines Events -->
        {#if editingEvent}
            <div class="edit-form">
                <form on:submit|preventDefault={saveEvent}>
                    <label for="name">Name:</label>
                    <input type="text" id="name" bind:value={editingEvent.name} />

                    <label for="location">Location:</label>
                    <input type="text" id="location" bind:value={editingEvent.location} />

                    <label for="startdate">Start Date:</label>
                    <input type="datetime-local" id="startdate" bind:value={editingEvent.startdate} />

                    <label for="available_normal_tickets">Available Normal Tickets:</label>
                    <input type="number" id="available_normal_tickets" bind:value={editingEvent.available_normal_tickets} />

                    <label for="available_vip_tickets">Available VIP Tickets:</label>
                    <input type="number" id="available_vip_tickets" bind:value={editingEvent.available_vip_tickets} />

                    <button type="submit">Save</button>
                    <button type="button" on:click={cancelEditing}>Cancel</button>
                </form>
            </div>
        {/if}
    </div>
{/if}

<style>
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
    .navbar {
        padding: 1rem;
        border-bottom: 2px solid #ffc107;
    }

    .navbar-brand {
        font-size: 1.5rem;
        font-weight: bold;
        color: #333;
    }

    .navbar-brand:hover {
        text-decoration: none;
        color: #000;
    }

    .ms-auto {
        margin-left: auto; /* Rechts ausgerichtet */
    }

    .me-3 {
        margin-right: 1rem; /* Abstand zwischen Benutzername und Admin-Link */
    }

    .btn-secondary {
        background-color: #6c757d;
        color: white;
        border: none;
        padding: 0.5rem 1rem;
        border-radius: 5px;
        font-size: 1rem;
        text-decoration: none;
        cursor: pointer;
    }

    .btn-secondary:hover {
        background-color: #5a6268;
        text-decoration: none;
    }
</style>
