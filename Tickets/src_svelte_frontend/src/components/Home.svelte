<script>
    import TicketsDisplay from "./TicketsDisplay.svelte";

    let eventID;

    let ticketsData = '';
    let errorMessage = '';

    let ticketID = ''
    let ticketPrice = '';
    let ticketRow = '';
    let ticketSeatNumber = '';
    let ticketVIP = false;
    let ticketUID = '';
    let ticketEID = '';

    let changedTicketID = ''
    let changedTicketPrice = '';
    let changedTicketRow = '';
    let changedTicketSeatNumber = '';
    let changedTicketVIP = '';
    let changedTicketUID = '';
    let changedTicketEID = '';

    let ticketToDelete = '';

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
        `event_id=${encodeURIComponent(ticketEID)}`,
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

    async function editTickets() {
        const response = await fetch(`/api/tickets?ticket_id=${encodeURIComponent(changedTicketID)}&`+
        `price=${encodeURIComponent(changedTicketPrice)}&row=${encodeURIComponent(changedTicketRow)}&`+
        `seat_number=${encodeURIComponent(changedTicketSeatNumber)}&vip=${encodeURIComponent(changedTicketVIP)}&`+
        `user_id=${encodeURIComponent(changedTicketUID)}&event_id=${encodeURIComponent(changedTicketEID)}`,
        {
            method: "PUT"
        });

        if (response.ok) {
          errorMessage = '';
          ticketsData = await response.json();
        } else {
          ticketsData = '';
          errorMessage = 'Failed to fetch tickets data';
        }
    }

    async function deleteTickets() {
        const response = await fetch(`/api/tickets?ticket_id=${encodeURIComponent(ticketToDelete)}`,
        {
            method: "DELETE"
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
        <h2 class="text-center">Add Ticket</h2>
        <form on:submit|preventDefault={addTickets} class="input-group">
          <div class="input-group-prepend">
            <label class="input-group-text">Price, Row, Seat, UID, EID, VIP</label>
          </div>
          <input type="number" class="form-control" bind:value={ticketPrice} placeholder="Price"/>
          <input type="text" class="form-control" bind:value={ticketRow} placeholder="Row"/>
          <input type="number" class="form-control" bind:value={ticketSeatNumber} placeholder="Seat"/>
          <input type="number" class="form-control" bind:value={ticketUID} placeholder="UID"/>
          <input type="number" class="form-control" bind:value={ticketEID} placeholder="EID"/>
          <select class="form-control" bind:value={ticketVIP}>
              <option value="false">False</option>
              <option value="true">True</option>
          </select>
          <button disabled={!ticketPrice && !ticketUID && !ticketEID} type="submit" class="btn btn-primary input-group-append">Add Ticket</button>
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