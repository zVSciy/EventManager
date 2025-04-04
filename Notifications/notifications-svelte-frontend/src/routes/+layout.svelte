<script>
	import '../app.css';

	let { children } = $props();
</script>

<nav style="padding: 1rem; background-color: #f8f8f8;">
	<div style="display: flex; justify-content: space-between; align-items: center;">
		<div>
			<a href="/app_ticket" style="margin-right: 1rem;">Tickets</a>
			<a href="/app_event">Events</a>
		</div>
		<div>
			<a href="/register">Register</a>
		</div>
	</div>
</nav>

{@render children()}
