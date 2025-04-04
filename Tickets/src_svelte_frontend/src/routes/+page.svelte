<script>
    import Home from "../components/Home.svelte";
	export let data;
</script>

<svelte:head>
	<title>Tickets | Main</title>
	<meta name="description" content="Tickets | Main" />
</svelte:head>

<section>
	<main>
		<Home {data}/>
	</main>
</section>