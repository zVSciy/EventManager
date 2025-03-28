<script>
    import Notifications from '../components/Notifications.svelte';

    export let notifications = [];
</script>

<main>
    <Notifications {notifications} />
</main>

<style>
    main {
        font-family: Arial, sans-serif;
        max-width: 600px;
        margin: 0 auto;
        padding: 20px;
    }
</style>