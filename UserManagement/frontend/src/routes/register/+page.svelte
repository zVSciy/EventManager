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
            alert(result);
            sessionStorage.setItem('email', email);
            sessionStorage.setItem('password', password);
            const targetUrl = `/app_event`;
            window.location.href = targetUrl;
        } catch (error) {
            alert("Failed to register user: " + error.message);
        }
    }
</script>

<nav>
    <a href="/">Home</a>
    <a href="/register">Register</a>
    <a href="/verify">Verify</a>
</nav>

<h1>Register</h1>
<form on:submit|preventDefault={register}>
    <input bind:value={email} placeholder="Email" required />
    <input bind:value={password} type="password" placeholder="Password" required />
    <input bind:value={first_name} placeholder="First Name" required />
    <input bind:value={last_name} placeholder="Last Name" required />
    <button>Register</button>
</form>