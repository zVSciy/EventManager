<script>
    import TicketsDisplay from "./TicketsDisplay.svelte";

    let eventID;

    let ticketsData = '';
    let errorMessage = '';

    async function getTickets() {
        const response = await fetch(`/api/tickets?event_id=${encodeURIComponent(eventID)}`,
        {
            method: "GET"
        });

        if (response.ok) {
          errorMessage = '';
          ticketsData = await response.json();
        } else {
          ticketsData = '';
          errorMessage = 'Failed to fetch tickets data';
        }
    }

</script>

<nav class="p-2 navbar navbar-expand-sm bg-dark navbar-dark">
    <h1 class="navbar-brand">Main - Tickets</h1>
</nav>

<div class="Home p-3 mt-3 text-center container"> 
    <div class="row">
      <div class="col-12">
        <form on:submit|preventDefault={getTickets} class="input-group">
          <div class="input-group-prepend">
            <label class="input-group-text">EventID (optional)</label>
          </div>
          <input type="number" class="form-control" bind:value={eventID} placeholder="EventID"/>
          <button type="submit" class="btn btn-primary input-group-append">Get Tickets</button>
        </form>
      </div>

    {#if ticketsData && ticketsData.error}
        <div class="message text-danger col-12">
          <h2>{ticketsData.status}</h2>
          <p>{ticketsData.error}</p>
        </div>
    {:else if errorMessage}
        <div class="message text-warning col-12">
          <h2>Global Error</h2>
          <p>{errorMessage}</p>
        </div>
    {:else}
        <TicketsDisplay {ticketsData}/>
    {/if}
    </div>
</div>

<style>
    .Home {
      width: 90em;
    }
  
    .message {
      margin-top: 2em;
      border-radius: 1em;
      border-style: solid;
      border-color: #e1e1e1;
      background-color: #f3f3f3;
      padding: 1em;
      text-align: center;
    }
  
    .message p {
      font-size: 1em;
    }
</style>