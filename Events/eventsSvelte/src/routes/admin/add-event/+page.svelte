<script>
    import { base } from '$app/paths';
    import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

    let name = '';
    let location = '';
    let organiser = '';
    let startDate = '';
    let availableNormalTickets = '';
    let availableVIPTickets = '';
    let error = null;
    let successMessage = '';
    let email = '';
    let password = '';

      onMount(() => {
    
        email = sessionStorage.getItem('email');
        password = sessionStorage.getItem('password')
        console.log(email)
      
    });
     async function token() {
        try {
            const response = await fetch("/api/token", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            const result = await response.json();
            if (result == 200) {
                addEvent()
            }

        } catch (error) {
            alert("Failed to verify token: " + error.message);
        }
    }
    
    async function addEvent() {
      try {
        const response = await fetch(`${base}/api/event`, {
          method: "POST",
          body: JSON.stringify({
            name: name,
            location: location,
            organisator: organiser,
            startdate: startDate,
            available_normal_tickets: availableNormalTickets,
            available_vip_tickets: availableVIPTickets
          }),
          headers: {
            "Content-Type": "application/json; charset=UTF-8",
          }
        });
  
        console.log(response);
  
        if (response.status === 200) {
          alert('Event wurde erfolgreich erstellt!');
          goto(`${base}/admin`)
        } else {
          alert('Fehler beim Erstellen des Event!');
        }
      } catch (error) {
        console.error("Fehler beim Fetchen:", error);
      } 
    } 

</script>



<h1>Add New Event</h1>
<form on:submit|preventDefault={token}>
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
