<script>
    import { apiFetch } from "../../lib/api";

    let email = "";
    let password = "";
    let first_name = "";
    let last_name = "";
    let role = "user";

    async function register() {
        await apiFetch("/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password, first_name, last_name, role }),
        });
        alert("User registered successfully!");
    }
</script>

<form on:submit|preventDefault={register}>
    <input bind:value={email} placeholder="Email" />
    <input bind:value={password} type="password" placeholder="Password" />
    <button>Register</button>
</form>
