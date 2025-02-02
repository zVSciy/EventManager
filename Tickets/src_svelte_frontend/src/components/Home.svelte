<script>
    import TicketsDisplay from "./TicketsDisplay.svelte";
    import { onMount } from 'svelte';
    export let data;

    let eventID;

    // ID aus dem Session Storage abrufen
    onMount(() => {
        const params = new URLSearchParams(window.location.search);
        eventID = params.get('eventId');
        console.log(eventID)
    });
    
    let moreTickets = 0;
    let ticketsData = '';
    let errorMessage = '';

    let ticketPrice = '';
    let ticketRow = '';
    let ticketSeatNumber = '';
    let ticketVIP = 'false';
    let ticketUID = '';

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

    async function addTickets() {
        const response = await fetch(`/api/tickets?price=${encodeURIComponent(ticketPrice)}&`+
        `row=${encodeURIComponent(ticketRow)}&seat_number=${encodeURIComponent(ticketSeatNumber)}&`+
        `vip=${encodeURIComponent(ticketVIP)}&user_id=${encodeURIComponent(ticketUID)}&`+
        `event_id=${encodeURIComponent(eventID)}`,
        {
          method: "POST"
        });

        if (response.ok) {
          errorMessage = '';
          ticketsData = await response.json();
        } else {
          ticketsData = '';
          errorMessage = 'Failed to fetch tickets data';
        }
    }

    async function updateAvailableTickets() {
        const response = await fetch(`/api/events?event_id=${eventID}&delete=${moreTickets}&vip=${encodeURIComponent(ticketVIP)}`,
        {
          method: "PUT"
        });
        
        if (response.ok) {
          errorMessage = '';
          let eventData = await response.json();

          if (eventData[0].status == 200 && moreTickets == 0){
            addTickets();
          } 

        } else {
          eventData = '';
        }
    }
</script>

<nav class="navbar navbar-expand-md bg-dark navbar-dark">
    <div class="container-fluid">
        <span class="navbar-brand">Tickets | Main</span>
        <div class="collapse navbar-collapse">
            <ul class="navbar-nav me-auto">
                <li class="nav-item">
                    <a class="nav-link active" href="/">Main</a>
                </li>
                {#if data && data.admin}
                    <li class="nav-item">
                        <a class="nav-link" href="/admin">Admin</a>
                    </li>
                {/if}
            </ul>

            {#if data && data.username}
                <ul class="navbar-nav ms-auto">
                    <li class="nav-item">
                        <span class="navbar-text text-light">Welcome,
                            {data.username}</span>
                        <a class="btn btn-sm btn-primary ms-2" href="/">Logout</a>
                    </li>
                </ul>
            {/if}
        </div>
    </div>
</nav>

<div class="Home p-3 mt-3 text-center container"> 
    <div class="row">
      <div class="col-12">
        <h2 class="text-center">Get Tickets</h2>
        <form on:submit|preventDefault={getTickets} class="input-group">
          <div class="input-group-prepend">
            <label class="input-group-text">EventID (optional)</label>
          </div>
          <input type="number" class="form-control" bind:value={eventID} placeholder="EventID"/>
          <button type="submit" class="btn btn-primary input-group-append">Get Tickets</button>
        </form>
      </div>

      <div class="col-12 mt-4">
        <h2 class="text-center">Buy Tickets</h2>
        <form on:submit|preventDefault={updateAvailableTickets} class="input-group">
          <div class="input-group-prepend">
            <label class="input-group-text">Price, Row, Seat, UID, VIP</label>
          </div>
          <input type="number" class="form-control" bind:value={ticketPrice} placeholder="Price"/>
          <input type="text" class="form-control" bind:value={ticketRow} placeholder="Row"/>
          <input type="number" class="form-control" bind:value={ticketSeatNumber} placeholder="Seat"/>
          <input type="number" class="form-control" bind:value={ticketUID} placeholder="UID"/>
          <select class="form-select" bind:value={ticketVIP}>
              <option value="false" selected>False</option>
              <option value="true">True</option>
          </select>
          <button disabled={!ticketPrice || !ticketUID } type="submit" class="btn btn-primary input-group-append"on:click={() => { moreTickets = 0 }}>Add Ticket</button>
        </form>  
      </div>

    <div class="col-12 mt-4">
      <h2 class="text-center">Results</h2>
    {#if ticketsData && ticketsData.error}
        <div class="message text-danger">
          <h2>{ticketsData.status}</h2>
          <p>{ticketsData.error}</p>
        </div>
    {:else if errorMessage}
        <div class="message text-warning">
          <h2>Global Error</h2>
          <p>{errorMessage}</p>
        </div>
    {:else}
        <TicketsDisplay {ticketsData}/>
    {/if}
    </div>
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