<script>
    let userData = {};

    async function fetchUserData() {
        try {
            const response = await fetch("/api/token", {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            userData = await response.json();
        } catch (error) {
            alert("Failed to fetch user data: " + error.message);
        }
    }

    fetchUserData();
</script>

<h1>Welcome, {userData.email}!</h1>