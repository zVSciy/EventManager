<script>
    let email = "";
    let password = "";
    let first_name = "";
    let last_name = "";
    let role = "user";

    async function register() {
        try {
            const response = await fetch("/api/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password, first_name, last_name, role }),
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            const result = await response.json();
            alert(result.message);
        } catch (error) {
            alert("Failed to register user: " + error.message);
        }
    }
</script>

<form on:submit|preventDefault={register}>
    <input bind:value={email} placeholder="Email" />
    <input bind:value={password} type="password" placeholder="Password" />
    <input bind:value={first_name} placeholder="First Name" />
    <input bind:value={last_name} placeholder="Last Name" />
    <button>Register</button>
</form>