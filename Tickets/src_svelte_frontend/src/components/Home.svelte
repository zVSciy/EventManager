<script>
    import { base } from '$app/paths';
    import TicketsDisplay from './TicketsDisplay.svelte';
    import TokenInvalid from './TokenInvalid.svelte';
    import { onMount } from 'svelte';
    export let data;
    let ticket;
    let eventID;

    // Get ID from Session Storage
    onMount(() => {
        eventID = sessionStorage.getItem('eventId');
        email = sessionStorage.getItem('email');
        password = sessionStorage.getItem('password');
        token();
        ticketUID = email;
        ticketEID = eventID;
    });
    
    let moreTickets = 0;
    let ticketsData = '';
    let errorMessage = '';
  
    let ticketEID = '';
    let ticketPrice = '';
    let ticketRow = '';
    let ticketSeatNumber = '';
    let ticketVIP = 'false';
    let ticketUID = '';

    let ticketToDelete;

    let email = '';
    let password = '';
    let error = false;

    async function token() {
      try {
        const response = await fetch(`${base}/api/token?email=${encodeURIComponent(email)}&password=${encodeURIComponent(password)}`,
        {
          method: "POST",
        });

        const result = await response.json();
        if (result.status == 200) {
          error = false;
        } else {
          error = true;
        }
      } catch (err) {
        alert("An error occured: " + err.message);
      }
    }

    async function getTickets() {
        const response = await fetch(`${base}/api/user?user_id=${encodeURIComponent(email)}&event_id=${encodeURIComponent(ticketEID)}`,
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
        const response = await fetch(`${base}/api/tickets?price=${encodeURIComponent(ticketPrice)}&`+
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

        const response = await fetch(`${base}/api/events?event_id=${eventID}&delete=${moreTickets}&vip=${encodeURIComponent(ticketVIP)}`,
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

    async function getOneTicket() {
        const response = await fetch(`${base}/api/tickets/single?ticket_id=${encodeURIComponent(ticketToDelete)}`,
        {
          method: "GET"
        });

        if (response.ok) {
          errorMessage = '';
          ticket = await response.json();
          ticketVIP = String(ticket.vip);
          deleteTickets(); 
        } else {
          ticket = '';
          errorMessage = 'Failed to fetch tickets data';
        }
    }

    async function deleteTickets() {
        const response = await fetch(`${base}/api/user?user_id=${encodeURIComponent(email)}&ticket_id=${encodeURIComponent(ticketToDelete)}`,
        {
          method: "DELETE"
        });

        if (response.ok) {
          errorMessage = '';
          ticketsData = await response.json();
          if (ticketsData.status == 200 && moreTickets == 1) {
            updateAvailableTickets(); 
          }
        } else {
          ticketsData = '';
          errorMessage = 'Failed to fetch tickets data';
        }
    }

    function logout() {
        // Clear session storage
        sessionStorage.clear();
        // Redirect to the root path (user management)
        window.location.href = "/";
    }
</script>

{#if error == true}
  <TokenInvalid/>
{:else}
  <nav class="navbar navbar-expand-md bg-dark navbar-dark">
      <div class="container-fluid">
          <span class="navbar-brand">Tickets | Main</span>
          <div class="collapse navbar-collapse">
              <ul class="navbar-nav me-auto">
                  <li class="nav-item">
                      <a class="nav-link" href="/app_event">Events</a>
                  </li>
                  <li class="nav-item">
                      <a class="nav-link active" href="{base}">Main</a>
                  </li>
                  {#if data && data.admin}
                      <li class="nav-item">
                          <a class="nav-link" href="{base}/admin">Admin</a>
                      </li>
                  {/if}
              </ul>

              {#if email}
                  <ul class="navbar-nav ms-auto">
                      <li class="nav-item">
                          <span class="navbar-text text-light">Welcome,
                              {email}</span>
                          <a class="btn btn-sm btn-primary ms-2" href="/" on:click={logout}>Logout</a>
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
            <input type="number" class="form-control" bind:value={ticketEID} placeholder="EventID"/>
            <button type="submit" class="btn btn-primary input-group-append">Get Tickets</button>
          </form>
        </div>

        <div class="col-12 mt-4">
          <h2 class="text-center">Buy Tickets</h2>
          <form on:submit|preventDefault={updateAvailableTickets} class="input-group">
            <div class="input-group-prepend">
              <label class="input-group-text">Price, Row, Seat, EID, VIP</label>
            </div>
            <input type="number" class="form-control" bind:value={ticketPrice} placeholder="Price"/>
            <input type="text" class="form-control" bind:value={ticketRow} placeholder="Row"/>
            <input type="number" class="form-control" bind:value={ticketSeatNumber} placeholder="Seat"/>
            <input type="number" class="form-control" bind:value={eventID} placeholder="EID" disabled readonly/>
            <select class="form-select" bind:value={ticketVIP}>
                <option value="false" selected>False</option>
                <option value="true">True</option>
            </select>
            <button disabled={!ticketPrice || !ticketUID } type="submit" class="btn btn-primary input-group-append"on:click={() => { moreTickets = 0 }}>Add Ticket</button>
          </form>  
        </div>

        <div class="col-12 mt-4">
          <h2 class="text-center">Cancel Tickets</h2>
          <form on:submit|preventDefault={getOneTicket} class="input-group">
            <div class="input-group-prepend">
              <label class="input-group-text">TicketID</label>
            </div>
            <input type="number" class="form-control" bind:value={ticketToDelete} placeholder="TicketID"/>
            <button disabled={!ticketToDelete} type="submit" class="btn btn-primary input-group-append"on:click={() => { moreTickets = 1 }}>Delete Ticket</button>
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
{/if}

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