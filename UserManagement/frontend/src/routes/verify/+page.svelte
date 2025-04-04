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
            alert(result + " - Token verified successfully!");
        } catch (error) {
            alert("Failed to verify token: " + error.message);
        }
    }
</script>

<nav>
    <a href="/">Home</a>
    <a href="/register">Register</a>
    <a href="/verify">Verify</a>
</nav>

<h1>Verify Token</h1>
<form on:submit|preventDefault={token}>
    <input bind:value={email} placeholder="Email" required />
    <input bind:value={password} type="password" placeholder="Password" required />
    <button>Verify</button>
</form>