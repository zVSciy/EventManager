<script>
    export let data;
    import Home_Admin from "../../components/Home_Admin.svelte";
</script>

<svelte:head>
    <title>Tickets | Admin</title>
    <meta name="description" content="Tickets | Admin" />
</svelte:head>

<section>
    <main>
        <Home_Admin {data} />
    </main>
</section>