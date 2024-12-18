<script>
    import { onMount } from 'svelte';
    let eventId;

    // ID aus dem Session Storage abrufen
    onMount(() => {
        eventId = sessionStorage.getItem('eventId');
        console.log(`Received Event ID: ${eventId}`);
    });
</script>

<h2>Buy Tickets for Event {eventId}</h2>