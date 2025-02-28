<script>
    import { base } from '$app/paths';
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    let events = [];
    let error = null;

    async function loadEvents() {
        const response = await fetch(`${base}/api/event`);
        events = await response.json();

    }

    onMount(loadEvents);

    function viewDetails(event) {
    if (event && event.ID) {
        goto(`${base}/details/${event.ID}`); // Navigiere zur Details-Seite basierend auf der Event-ID
    }
}
export let data;
</script>


<nav class="navbar navbar-expand-lg navbar-light bg-warning">
    <div class="container-fluid">
        <button on:click={() => window.location.href=`${base}/`} style="margin-bottom: 20px; padding: 10px; background-color: #009879; color: white; border: none; border-radius: 4px; cursor: pointer;">Event Manager</button>    

        <div class="d-flex align-items-center ms-auto">
            <span class="me-3">Hi, {data.username}!</span>

            {#if data.admin}
            <button on:click={() => window.location.href=`${base}/admin`} style="margin-bottom: 20px; padding: 10px; background-color: #009879; color: white; border: none; border-radius: 4px; cursor: pointer;">Admin</button>    
            {/if}
        </div>
    </div>
</nav>


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
    h1 {
        display: flex;
        align-items: center;
        justify-content: center;
        display: flex;
    }
</style>
