<script>
    import { apiFetch } from "../../lib/api";

    let userData = {};

    async function fetchUserData() {
        try {
            userData = await apiFetch("/verify", {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            });
        } catch (error) {
            alert("Failed to fetch user data." + error);
        }
    }

    fetchUserData();
</script>

<h1>Welcome, {userData.email}!</h1>