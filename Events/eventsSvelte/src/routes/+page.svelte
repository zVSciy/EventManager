<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    let events = [];
    let error = null;

    async function loadEvents() {
        const response = await fetch('/api/event');
        events = await response.json();

    }

    onMount(loadEvents);

    function viewDetails(event) {
    if (event && event.ID) {
        goto(`/details/${event.ID}`); // Navigiere zur Details-Seite basierend auf der Event-ID
    }
}

</script>

{#if error}
    <p>{error}</p>
{:else}
    <div>
        <h1>Event List</h1>
        <table class="styled-table">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Location</th>
                    <th>Start Date</th>
                    <th>Canceled</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each events as event}
                    <tr>
                        <td>{event.ID}</td> <!-- Sicherstellen, dass ID existiert -->
                        <td>{event.name}</td>
                        <td>{event.location}</td>
                        <td>{new Date(event.startdate).toLocaleString()}</td>
                        <td>{event.canceled ? 'Yes' : 'No'}</td>
                        <td>
                            <button on:click={() => viewDetails(event)}>Details</button>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
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
</style>
