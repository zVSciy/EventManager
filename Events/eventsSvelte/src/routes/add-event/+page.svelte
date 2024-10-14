<script>
    import { onMount } from 'svelte';
    import { fetchEvents } from '$lib/api';

    let name = '';
    let location = '';
    let organiser = '';
    let startDate = '';
    let availableNormalTickets = '';
    let availableVIPTickets = '';
    let error = null;
    let successMessage = '';

    async function addEvent() {
        try {
            const response = await fetch('http://localhost:8000/event/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name,
                    location,
                    organisator: organiser,
                    startdate: startDate,
                    available_normal_tickets: parseInt(availableNormalTickets),
                    available_vip_tickets: parseInt(availableVIPTickets),
                }),
            });

            if (!response.ok) {
                throw new Error('Failed to create event');
            }

            const result = await response.json();
            successMessage = `Event was created successfully with ID: ${result.eventID}`;
            
            // Redirect to the main page after creation
            setTimeout(() => {
                window.location.href = '/'; // Navigiere zurück zur Hauptseite
            }, 2000); // Optional: Zeit geben, um die Erfolgsmeldung anzuzeigen

            // Optionally, reset the form fields
            name = '';
            location = '';
            organiser = '';
            startDate = '';
            availableNormalTickets = '';
            availableVIPTickets = '';
        } catch (err) {
            error = err.message;
        }
    }
</script>



<h1>Add New Event</h1>
<form on:submit|preventDefault={addEvent}>
    <label>
        Name:
        <input type="text" bind:value={name} required>
    </label>
    <br>

    <label>
        Location:
        <input type="text" bind:value={location} required>
    </label>
    <br>

    <label>
        Organiser:
        <input type="text" bind:value={organiser} required>
    </label>
    <br>

    <label>
        Start Date:
        <input type="datetime-local" bind:value={startDate} required>
    </label>
    <br>

    <label>
        Available Normal Tickets:
        <input type="number" bind:value={availableNormalTickets} required>
    </label>
    <br>

    <label>
        Available VIP Tickets:
        <input type="number" bind:value={availableVIPTickets} required>
    </label>
    <br>

    <button type="submit">Add Event</button>
    {#if error}
    <p style="color: red;">{error}</p>
{/if}
{#if successMessage}
    <p style="color: green;">{successMessage}</p>
{/if}
</form>


<style>
    form {
        display: flex;
        flex-direction: column;
        max-width: 400px;
        margin: 20px auto;
        padding: 20px;
        border: 1px solid #ccc;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    }

    label {
        margin-bottom: 10px;
    }
    h1 {
        display: flex;
        align-items: center;
        justify-content: center;
        display: flex;
    }
    input {
        padding: 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
        margin-top: 5px;
    }

    button {
        padding: 10px;
        background-color: #009879;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        margin-top: 10px;
    }

    button:hover {
        background-color: #007f66;
    }
</style>
