<script>
    export let ticketsData;
</script>

<div class="ticketsBox text-center">
    {#if ticketsData.length == 0}
    <p class="specialParagraph">No tickets could be found.</p>
    {:else if ticketsData}
            <div class="row row-cols-lg-5 row-cols-sm-2 justify-content-center p-2">
            {#each ticketsData as data}
                <div class="col">
                    <h2>Ticket {data.id}</h2>
                    <p>Price: {data.price}<br>
                    Paid: {data.paid}<br>
                    Row: {data.row}<br>
                    Seat: {data.seat_number}<br>
                    VIP: {data.vip}<br>
                    User: {data.user_id}<br>
                    Event: {data.event_id}<br>
                </div>
            {/each}
            </div>
    {:else}
        <p> No data! </p>
    {/if}
</div>

<style>
  .ticketsBox {
    margin-top: 2em;
    border-radius: 1em;
    border-style: solid;
    border-color: #e1e1e1;
    background-color: #f3f3f3;
    padding: 1em;
    text-align: center;
  }

  .specialParagraph {
    font-size: 1em;
  }

  h2 {
    font-size: 1em;
    font-weight: bold;
  }
</style>