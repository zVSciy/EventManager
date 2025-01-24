<script>
    import { onMount } from 'svelte';

    let notifications = [];
    let newNotification = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };
    let updateNotificationId = null;
    let updateNotificationData = { description: '', status: '', eventId: '', ticketId: '', timestamp: '', paymentId: '', userId: '' };

    onMount(async () => {
        await fetchNotifications();
    });

    async function fetchNotifications() {
        try {
            const response = await fetch('http://192.168.75.81:8082/notifications');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            notifications = await response.json();
        } catch (error) {
            console.error("Error fetching notifications:", error);
        }
    }

    async function createNotification() {
        try {
            const response = await fetch('http://192.168.75.81:8082/notifications', {
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
            const response = await fetch(`http://192.168.75.81:8082/notifications/${updateNotificationId}`, {
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
            const response = await fetch(`http://192.168.75.81:8082/notifications/${id}`, {
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
</script>

<style>
    .notification {
        border: 1px solid #ccc;
        padding: 10px;
        margin-bottom: 10px;
    }
</style>

<div>
    <h1>Notifications</h1>
    <div>
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
    <div>
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