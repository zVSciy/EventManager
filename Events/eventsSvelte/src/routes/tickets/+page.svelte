<script>
 import { eventStore } from '$lib/index';

let eventId;

// Abonniere den Store
eventStore.subscribe(id => {
    eventId = id;
    console.log(`Received Event ID: ${eventId}`);
});
</script>

<h2>Buy Tickets for Event {eventId}</h2>
