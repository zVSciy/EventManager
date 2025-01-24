<script>
    import { login } from "../lib/auth";

    let email = "";
    let password = "";

    async function submitLogin() {
        try {
            await login(email, password);
            alert("Login successful!");
            window.location.href = "/dashboard";
        } catch (error) {
            alert("Login failed!" + error);
        }
    }
</script>

<form on:submit|preventDefault={submitLogin}>
    <input type="email" bind:value={email} placeholder="Email" required />
    <input type="password" bind:value={password} placeholder="Password" required />
    <button type="submit">Login</button>
</form>