<script>
    export let ticketsData;
</script>

<div class="ticketsBox text-center">
    {#if ticketsData.length == 0}
        <p class="specialParagraph">No tickets could be found.</p>
        {:else if ticketsData.msg}
        <p class="specialParagraph">{ticketsData.msg}</p>
        {:else if ticketsData}
        <table class="table table-hover table-striped">
            <thead>
                <tr>
                    <th scope="col">ID</th>
                    <th scope="col">Price</th>
                    <th scope="col">Paid</th>
                    <th scope="col">Row</th>
                    <th scope="col">Seat</th>
                    <th scope="col">VIP</th>
                    <th scope="col">User</th>
                    <th scope="col">Event</th>
                </tr>
            </thead>
            <tbody>
                {#each ticketsData as data}
                    <tr>
                        <td>{data.id}</td>
                        <td>{data.price}</td>
                        <td>{data.paid}</td>
                        <td>{data.row}</td>
                        <td>{data.seat_number}</td>
                        <td>{data.vip}</td>
                        <td>{data.user_id}</td>
                        <td>{data.event_id}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {:else}
        <p>
            No data!
        </p>
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