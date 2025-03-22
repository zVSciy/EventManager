<script>
    let email = "";
    let password = "";

    async function token() {
        try {
            const response = await fetch("/api/token", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            const result = await response.json();
            alert(result);
        } catch (error) {
            alert("Failed to verify token: " + error.message);
        }
    }
</script>

<form on:submit|preventDefault={token}>
    <input bind:value={email} placeholder="Email" />
    <input bind:value={password} type="password" placeholder="Password" />
    <button>Register</button>
</form>