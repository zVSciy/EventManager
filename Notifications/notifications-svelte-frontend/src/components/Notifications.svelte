<script>
    import { onMount } from 'svelte';
    import { base } from '$app/paths';

    const backend_url = `${base}/api`;

    let notifications = [];
    let newNotification = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };
    let updateNotificationId = null;
    let updateNotificationData = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };

    // Get ID from Session Storage
    onMount(() => {
        email = sessionStorage.getItem('email');
        password = sessionStorage.getItem('password');
        token();
    });

    let email = '';
    let password = '';
    let error = false;

    onMount(async () => {
        await fetchNotifications();
    });

    async function fetchNotifications() {
        try {
            // Use the email as userId when fetching notifications
            const response = await fetch(`${backend_url}/user/${encodeURIComponent(email)}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json(); // Store the parsed JSON in a variable
            console.log(data);
            notifications = data; // Use the stored data
        } catch (error) {
            console.error("Error fetching notifications:", error);
        }
    }

    async function createNotification() {
        try {
            const response = await fetch(`${backend_url}/notifications`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(newNotification)
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            await fetchNotifications();
            newNotification = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };
        } catch (error) {
            console.error("Error creating notification:", error);
        }
    }

    async function updateNotification() {
        try {
            console.log(password)
            console.log(email)
            console.log(error)
            const response = await fetch(`${backend_url}/notifications/${updateNotificationId}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updateNotificationData)
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            await fetchNotifications();
            updateNotificationId = null;
            updateNotificationData = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };
        } catch (error) {
            console.error("Error updating notification:", error);
        }
    }

    async function deleteNotification(id) {
        try {
            const response = await fetch(`${backend_url}/notifications/${id}`, {
                method: 'DELETE'
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            await fetchNotifications();
        } catch (error) {
            console.error("Error deleting notification:", error);
        }
    }

    async function token() {
        let result = "";
        try {
            const response = await fetch(`${base}/api/token?email=${encodeURIComponent(email)}&password=${encodeURIComponent(password)}`, {
                method: "POST",
            });
            result = await response.json();
            if (result.status == 200) {
                error = false;
            } else {
                error = true;
            }
        } catch (err) {
            console.log(result)
            alert("An error occured: " + err.message);
        }
    }
</script>

<style>
    /* New styling to improve UI appearance */
    .container {
        max-width: 800px;
        margin: auto;
        padding: 20px;
        font-family: Arial, sans-serif;
    }
    .header {
        text-align: center;
        margin-bottom: 20px;
    }
    .form-section {
        display: flex;
        gap: 20px;
        flex-wrap: wrap;
        margin-bottom: 20px;
    }
    .form-card {
        flex: 1;
        min-width: 280px;
        padding: 15px;
        border: 1px solid #ddd;
        border-radius: 8px;
        box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        background: #fff;
    }
    .form-card h2 {
        margin-top: 0;
    }
    input {
        width: 100%;
        padding: 8px;
        margin-bottom: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    button {
        padding: 8px 12px;
        background-color: #007acc;
        border: none;
        color: #fff;
        border-radius: 4px;
        cursor: pointer;
    }
    button:hover {
        background-color: #005fa3;
    }
    .notification-list {
        margin-top: 20px;
    }
    .notification {
        padding: 15px;
        border: 1px solid #e0e0e0;
        border-radius: 8px;
        margin-bottom: 15px;
        box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        background: #fff;
    }
    .error-message {
        color: red;
        text-align: center;
    }
</style>

{#if !error}
<div class="container">
    <div class="header">
        <h1>Notifications</h1>
    </div>
    <div class="form-section">
        <!-- Create Notification Card -->
        <div class="form-card">
            <h2>Create Notification</h2>
            <input type="text" placeholder="Description" bind:value={newNotification.description} />
            <input type="text" placeholder="Status" bind:value={newNotification.status} />
            <input type="number" placeholder="Event ID" bind:value={newNotification.eventId} />
            <input type="number" placeholder="Ticket ID" bind:value={newNotification.ticketId} />
            <input type="number" placeholder="Timestamp" bind:value={newNotification.timestamp} />
            <input type="number" placeholder="Payment ID" bind:value={newNotification.paymentId} />
            <input type="number" placeholder="User ID" bind:value={newNotification.userId} />
            <button on:click={createNotification}>Create</button>
        </div>
        <!-- Update Notification Card -->
        <div class="form-card">
            <h2>Update Notification</h2>
            <input type="number" placeholder="Notification ID" bind:value={updateNotificationId} />
            <input type="text" placeholder="Description" bind:value={updateNotificationData.description} />
            <input type="text" placeholder="Status" bind:value={updateNotificationData.status} />
            <input type="number" placeholder="Event ID" bind:value={updateNotificationData.eventId} />
            <input type="number" placeholder="Ticket ID" bind:value={updateNotificationData.ticketId} />
            <input type="number" placeholder="Timestamp" bind:value={updateNotificationData.timestamp} />
            <input type="number" placeholder="Payment ID" bind:value={updateNotificationData.paymentId} />
            <input type="number" placeholder="User ID" bind:value={updateNotificationData.userId} />
            <button on:click={updateNotification}>Update</button>
        </div>
    </div>
    <div class="notification-list">
        {#if notifications.length > 0}
            {#each notifications as notification}
                <div class="notification">
                    <p><strong>ID:</strong> {notification.id}</p>
                    <p><strong>Description:</strong> {notification.description}</p>
                    <p><strong>Status:</strong> {notification.status}</p>
                    <p><strong>Event ID:</strong> {notification.eventId}</p>
                    <p><strong>Ticket ID:</strong> {notification.ticketId}</p>
                    <p><strong>Timestamp:</strong> {new Date(notification.timestamp * 1000).toLocaleString()}</p>
                    <p><strong>Payment ID:</strong> {notification.paymentId}</p>
                    <p><strong>User ID:</strong> {notification.userId}</p>
                    <button on:click={() => deleteNotification(notification.id)}>Delete</button>
                </div>
            {/each}
        {:else}
            <p>No notifications found.</p>
        {/if}
    </div>
</div>
{:else}
<div class="container">
    <p class="error-message">An error occurred. Please verify your authentication at 
        <a href="/register">/register</a>
    </p>
</div>
{/if}